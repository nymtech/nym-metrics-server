package mixmining

import (
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

// CreateMixStatus adds a new PersistedMixStatus in the database.
func (service *Service) CreateMixStatus(status models.MixStatus) {
	persist := models.PersistedMixStatus{
		MixStatus: status,
		Timestamp: timemock.Now().UnixNano(),
	}
	service.db.Add(persist)

}

// List lists all mix metrics in the database
func (service *Service) List(pubkey string) []models.PersistedMixStatus {
	return service.db.List(pubkey)
}
