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
	report := service.buildMixStatusReport(persistedMixStatus)
	service.db.SaveMixStatusReport(report)
}

// List lists the given number mix metrics
func (service *Service) List(pubkey string) []models.PersistedMixStatus {
	return service.db.List(pubkey, 1000)
}

func (service *Service) buildMixStatusReport(status models.PersistedMixStatus) models.MixStatusReport {
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
	return report
}

// CalculateUptime calculates percentage uptime for a given node, protocol during a specific time period
func (service *Service) CalculateUptime(pubkey string, ipVersion string, timePeriod int64) int {

	return 0
}

func thirtyDaysAgo() int64 {
	now := timemock.Now()
	return now.Add(time.Duration(-30) * time.Hour * 24).UnixNano()
}
