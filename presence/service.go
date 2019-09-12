package presence

import (
	"time"

	"github.com/nymtech/directory-server/models"
)

type service struct {
	mixNodes  []models.MixNodePresence
	cocoNodes []models.Presence
}

// Service defines the REST service interface for presence.
type Service interface {
	NotifyMixNodePresence(up models.HostInfo) error
	NotifyCocoNodePresence(up models.HostInfo) error
	Up() error
}

func newService() *service {
	return &service{}
}

func (service *service) AddMixNodePresence(info models.MixHostInfo) error {
	presence := models.MixNodePresence{
		MixHostInfo: info,
		LastSeen:    time.Now().UnixNano(),
	}
	service.mixNodes = append(service.mixNodes, presence)
	return nil
}

func (service *service) AddCocoNodePresence(info models.HostInfo) error {
	presence := models.Presence{
		HostInfo: info,
		LastSeen: time.Now().UnixNano(),
	}
	service.cocoNodes = append(service.cocoNodes, presence)
	return nil
}

func (service *service) Up() ([]models.MixNodePresence, error) {
	return service.mixNodes, nil
}
