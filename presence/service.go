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
	NotifyMixNodePresence(up models.MixHostInfo)
	NotifyCocoNodePresence(up models.HostInfo)
	Topology() models.Topology
}

func newService(db Db) *service {
	return &service{db: db}
}

func (service *service) AddMixProviderPresence(info models.MixProviderHostInfo) {
	presence := models.MixProviderPresence{
		MixProviderHostInfo: info,
		LastSeen:            timemock.Now().UnixNano(),
	}
	service.db.AddMixProvider(presence)
}

func (service *service) AddMixNodePresence(info models.MixHostInfo) {
	presence := models.MixNodePresence{
		MixHostInfo: info,
		LastSeen:    timemock.Now().UnixNano(),
	}
	service.db.AddMix(presence)
}

func (service *service) AddCocoNodePresence(info models.CocoHostInfo) {
	presence := models.CocoPresence{
		CocoHostInfo: info,
		LastSeen:     timemock.Now().UnixNano(),
	}
	service.db.AddCoco(presence)
}

func (service *service) Topology() models.Topology {
	return service.db.Topology()
}
