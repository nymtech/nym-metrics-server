package metrics

import (
	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
)

type service struct {
	db Db
}

// Service defines the REST service interface for metrics.
type Service interface {
	CreateMixMetric() error
	List() []models.MixMetric
}

func newService(db Db) *service {
	return &service{db: db}
}

func (service *service) CreateMixMetric(metric models.MixMetric) {
	persist := models.PersistedMixMetric{
		MixMetric: metric,
		Timestamp: timemock.Now().UnixNano(),
	}
	service.db.Add(persist)
}

func (service *service) List() []models.PersistedMixMetric {
	return service.db.List()
}
