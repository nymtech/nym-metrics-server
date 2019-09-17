package metrics

import "github.com/nymtech/directory-server/models"

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
	service.db.Add(metric)
}

func (service *service) List() []models.MixMetric {
	return service.db.List()
}
