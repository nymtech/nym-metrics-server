package mixmining

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/nym-directory/models"
)

// Service struct
type Service struct {
	db IDb
}

// IService defines the REST service interface for mixmining.
type IService interface {
	CreateMixStatus(mixStatus models.MixStatus) models.PersistedMixStatus
	List(pubkey string) []models.PersistedMixStatus
	SaveStatusReport(status models.PersistedMixStatus) models.MixStatusReport
	GetStatusReport(pubkey string) models.MixStatusReport

	SaveBatchStatusReport(status []models.PersistedMixStatus) models.BatchMixStatusReport
	BatchCreateMixStatus(batchMixStatus models.BatchMixStatus) []models.PersistedMixStatus
	BatchGetMixStatusReport() models.BatchMixStatusReport
}

// NewService constructor
func NewService(db IDb) *Service {
	return &Service{
		db: db,
	}
}

// CreateMixStatus adds a new PersistedMixStatus in the orm.
func (service *Service) CreateMixStatus(mixStatus models.MixStatus) models.PersistedMixStatus {
	persistedMixStatus := models.PersistedMixStatus{
		MixStatus: mixStatus,
		Timestamp: timemock.Now().UnixNano(),
	}
	service.db.Add(persistedMixStatus)
	return persistedMixStatus
}

// List lists the given number mix metrics
func (service *Service) List(pubkey string) []models.PersistedMixStatus {
	return service.db.List(pubkey, 1000)
}

// GetStatusReport gets a single MixStatusReport by node public key
func (service *Service) GetStatusReport(pubkey string) models.MixStatusReport {
	return service.db.LoadReport(pubkey)
}

func (service *Service) BatchCreateMixStatus(batchMixStatus models.BatchMixStatus) []models.PersistedMixStatus {
	statusList := make([]models.PersistedMixStatus, len(batchMixStatus.Status))
	for i, mixStatus := range batchMixStatus.Status {
		persistedMixStatus := models.PersistedMixStatus{
			MixStatus: mixStatus,
			Timestamp: timemock.Now().UnixNano(),
		}
		statusList[i] = persistedMixStatus
	}
	service.db.BatchAdd(statusList)

	return statusList
}

func (service *Service) BatchGetMixStatusReport() models.BatchMixStatusReport {
// TODO
	return models.BatchMixStatusReport{}
}

func (service *Service) SaveBatchStatusReport(status []models.PersistedMixStatus) models.BatchMixStatusReport {
	// TODO: COMBINE REPORTS IF THEY USE THE SAME KEY (V4 and V6)
	
	pubkeys := make([]string, len(status))
	for i := range status {
		pubkeys[i] = status[i].PubKey
	}
	batchReport := service.db.BatchLoadReports(pubkeys)

	// that's super crude but I don't think db results are guaranteed to come in order, plus some entries might
	// not exist
	reportMap := make(map[string]int)
	for i, report := range batchReport.Report {
		reportMap[report.PubKey] = i
	}

	for _, mixStatus := range status {
		if reportIdx, ok := reportMap[mixStatus.PubKey]; ok {
			service.dealWithStatusReport(&batchReport.Report[reportIdx], &mixStatus)
		} else {
			var freshReport models.MixStatusReport
			service.dealWithStatusReport(&freshReport, &mixStatus)
			batchReport.Report = append(batchReport.Report, freshReport)
			reportMap[freshReport.PubKey] = len(batchReport.Report) - 1
		}
	}

	service.db.SaveBatchMixStatusReport(batchReport)
	return batchReport
}

func (service *Service) dealWithStatusReport(report *models.MixStatusReport, status *models.PersistedMixStatus) {
	report.PubKey = status.PubKey // crude, we do this in case it's a fresh struct returned from the db

	if status.IPVersion == "4" {
		report.MostRecentIPV4 = *status.Up
		report.Last5MinutesIPV4 = service.CalculateUptime(status.PubKey, "4", minutesAgo(5))
		report.LastHourIPV4 = service.CalculateUptime(status.PubKey, "4", minutesAgo(60))
		report.LastDayIPV4 = service.CalculateUptime(status.PubKey, "4", daysAgo(1))
		report.LastWeekIPV4 = service.CalculateUptime(status.PubKey, "4", daysAgo(7))
		report.LastMonthIPV4 = service.CalculateUptime(status.PubKey, "4", daysAgo(30))
	} else if status.IPVersion == "6" {
		report.MostRecentIPV6 = *status.Up
		report.Last5MinutesIPV6 = service.CalculateUptime(status.PubKey, "6", minutesAgo(5))
		report.LastHourIPV6 = service.CalculateUptime(status.PubKey, "6", minutesAgo(60))
		report.LastDayIPV6 = service.CalculateUptime(status.PubKey, "6", daysAgo(1))
		report.LastWeekIPV6 = service.CalculateUptime(status.PubKey, "6", daysAgo(7))
		report.LastMonthIPV6 = service.CalculateUptime(status.PubKey, "6", daysAgo(30))
	}
}

// SaveStatusReport builds and saves a status report for a mixnode. The report can be updated once
// whenever we receive a new status, and the saved result can then be queried. This keeps us from
// having to build the report dynamically on every request at runtime.
func (service *Service) SaveStatusReport(status models.PersistedMixStatus) models.MixStatusReport {
	report := service.db.LoadReport(status.PubKey)

	service.dealWithStatusReport(&report, &status)
	service.db.SaveMixStatusReport(report)
	return report
}

// CalculateUptime calculates percentage uptime for a given node, protocol since a specific time
func (service *Service) CalculateUptime(pubkey string, ipVersion string, since int64) int {
	statuses := service.db.ListDateRange(pubkey, ipVersion, since, now())
	numStatuses := len(statuses)
	if numStatuses == 0 {
		return 0
	}
	up := 0
	for _, status := range statuses {
		if *status.Up {
			up = up + 1
		}
	}
	return service.calculatePercent(up, numStatuses)
}

func (service *Service) calculatePercent(num int, outOf int) int {
	return int(float32(num) / float32(outOf) * 100)
}

func now() int64 {
	return timemock.Now().UnixNano()
}

func daysAgo(days int) int64 {
	now := timemock.Now()
	return now.Add(time.Duration(-days) * time.Hour * 24).UnixNano()
}

func minutesAgo(minutes int) int64 {
	now := timemock.Now()
	return now.Add(time.Duration(-minutes) * time.Minute).UnixNano()
}

func secondsAgo(seconds int) int64 {
	now := timemock.Now()
	return now.Add(time.Duration(-seconds) * time.Second).UnixNano()
}
