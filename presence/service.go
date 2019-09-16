package presence

import (
	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
)

type service struct {
	db Db
}

// Service defines the REST service interface for presence.
type Service interface {
	NotifyMixNodePresence(up models.HostInfo) error
	NotifyCocoNodePresence(up models.HostInfo) error
	Up() error
}

func newService(db Db) *service {
	return &service{db: db}
}

func (service *service) AddMixNodePresence(info models.MixHostInfo) {
	presence := models.MixNodePresence{
		MixHostInfo: info,
		LastSeen:    timemock.Now().Unix(),
	}
	service.db.AddMix(presence)
}

// func (service *service) AddCocoNodePresence(info models.HostInfo) {
// 	presence := models.Presence{
// 		HostInfo: info,
// 		LastSeen: timemock.Now().Unix(),
// 	}
// 	service.cocoNodes = append(service.cocoNodes, presence)
// }

func (service *service) List() models.Topology {
	return service.db.Topology()
}
