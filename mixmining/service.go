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

// IService defines the REST service interface for metrics.
type IService interface {
	CreateMixStatus(metric models.MixStatus)
	List(pubkey string) []models.PersistedMixStatus
}

// NewService constructor
func NewService(db IDb) *Service {
	return &Service{
		db: db,
	}
}

// CreateMixStatus adds a new PersistedMixStatus in the orm.
func (service *Service) CreateMixStatus(mixStatus models.MixStatus) {
	persistedMixStatus := models.PersistedMixStatus{
		MixStatus: mixStatus,
		Timestamp: timemock.Now().UnixNano(),
	}
	service.db.Add(persistedMixStatus)
}

// List lists the given number mix metrics
func (service *Service) List(pubkey string) []models.PersistedMixStatus {
	return service.db.List(pubkey, 1000)
}

// SaveStatusReport builds and saves a status report for a mixnode. The report can be updated once
// whenever we receive a new status, and the saved result can then be queried. This keeps us from
// having to build the report dynamically on every request at runtime.
func (service *Service) SaveStatusReport(status models.PersistedMixStatus) models.MixStatusReport {
	uptimeReport := models.UptimeReport{
		IPVersion:    status.IPVersion,
		MostRecent:   status.Up,
		Last5Minutes: service.CalculateUptime(status.PubKey, status.IPVersion, minutesAgo(5)),
		LastHour:     service.CalculateUptime(status.PubKey, status.IPVersion, minutesAgo(60)),
		LastDay:      service.CalculateUptime(status.PubKey, status.IPVersion, daysAgo(1)),
		LastWeek:     service.CalculateUptime(status.PubKey, status.IPVersion, daysAgo(30)),
		LastMonth:    0,
	}
	var report models.MixStatusReport
	report, err := service.db.LoadReport(status.PubKey)
	if err != nil {
		report = models.MixStatusReport{}
	}

	if status.IPVersion == "4" {
		report.IPV4Status = uptimeReport
	} else if status.IPVersion == "6" {
		report.IPV6Status = uptimeReport
	}
	service.db.SaveMixStatusReport(report)
	return report
}

// CalculateUptime calculates percentage uptime for a given node, protocol since a specific time
func (service *Service) CalculateUptime(pubkey string, ipVersion string, since int64) int {
	statuses := service.db.ListDateRange(pubkey, ipVersion, now(), since)
	numStatuses := len(statuses)
	if numStatuses == 0 {
		return 0
	}
	up := 0
	for _, status := range statuses {
		if status.Up {
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
