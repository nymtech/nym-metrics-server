package presence

type service struct{}

// Service defines the REST service interface for Nodes.
type Service interface {
	NotifyPresence() error
	Up() error
}

func newService(cfg *Config) *service {
	return &service{}
}

func (service *service) NotifyPresence() error {
	return nil
}

func (service *service) Up() error {
	return nil
}
