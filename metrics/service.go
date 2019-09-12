package metrics

import "github.com/nymtech/directory-server/models"

type service struct{}

// Service defines the REST service interface for metrics.
type Service interface {
	CreateMixMetric() error
}

func newService(cfg *Config) *service {
	return &service{}
}

func (service *service) CreateMixMetric(metric models.MixMetric) error {
	return nil
}
