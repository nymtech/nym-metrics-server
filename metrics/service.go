package metrics

import (
	"encoding/json"
	"fmt"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
	"github.com/nymtech/directory-server/server/websocket"
)

type service struct {
	db  Db
	hub websocket.Broadcaster
}

// Service defines the REST service interface for metrics.
type Service interface {
	CreateMixMetric() error
	List() []models.MixMetric
}

func newService(db Db, hub websocket.Broadcaster) *service {
	return &service{
		db:  db,
		hub: hub,
	}
}

func (service *service) CreateMixMetric(metric models.MixMetric) {
	persist := models.PersistedMixMetric{
		MixMetric: metric,
		Timestamp: timemock.Now().UnixNano(),
	}
	service.db.Add(persist)

	b, err := json.Marshal(persist)
	if err != nil {
		fmt.Println(err)
		return
	}
	service.hub.Notify(string(b))
}

func (service *service) List() []models.PersistedMixMetric {
	return service.db.List()
}
