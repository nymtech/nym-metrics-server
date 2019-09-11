package presence

import (
	"time"

	"github.com/nymtech/directory-server/models"
)

type service struct {
	mixNodes []models.Presence
}

// Service defines the REST service interface for presence.
type Service interface {
	NotifyMixNodePresence(up models.HostInfo) error
	Up() error
}

func newService(cfg *Config) *service {
	return &service{}
}

func (service *service) NotifyMixNodePresence(info models.HostInfo) error {
	presence := models.Presence{
		HostInfo: info,
		LastSeen: time.Now().UnixNano(),
	}
	service.mixNodes = append(service.mixNodes, presence)
	return nil
}

func (service *service) Up() ([]models.Presence, error) {
	return service.mixNodes, nil
}
