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
	// get all previous statuses from the database
	// init report struct

	report := models.MixStatusReport{
		PubKey: status.PubKey,
	}
	// check whether this one is IPv4 or IPv6
	// calculate 1 most recent
	// calculate last hour uptime
	// calculate last day uptime
	// calculate last month uptime
	// return report
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
	return int(float32(up) / float32(numStatuses) * 100)
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
