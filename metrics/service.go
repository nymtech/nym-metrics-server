package metrics

type service struct{}

// Service defines the REST service interface for Nodes.
type Service interface {
	CreateMetricsReport() error
}

func newService(cfg *Config) *service {
	return &service{}
}

func (service *service) CreateMetricsReport() error {
	return nil
}
