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

func (service *service) CreateMixMetric(metric models.MixMetric) error {
	service.db.Add(metric)
	return nil
}

func (service *service) List() []models.MixMetric {
	return service.db.List()
}
