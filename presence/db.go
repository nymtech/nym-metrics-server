package presence

import (
	"sync"
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/nym-directory/models"
)

// Db holds presence information
type Db interface {
	AddCoco(models.CocoPresence)
	AddMix(models.MixNodePresence)
	AddMixProvider(models.MixProviderPresence)
	Topology() models.Topology
}

type db struct {
	// TODO: it's slightly inefficient to have a single mutex for all database, because right now
	// if a mix node was being added, we wouldn't be able to touch cocoNodes
	sync.Mutex
	cocoNodes        []models.CocoPresence
	mixNodes         []models.MixNodePresence
	mixProviderNodes []models.MixProviderPresence
}

func newPresenceDb() *db {
	return &db{
		cocoNodes:        []models.CocoPresence{},
		mixNodes:         []models.MixNodePresence{},
		mixProviderNodes: []models.MixProviderPresence{},
	}
}

func (db *db) AddCoco(presence models.CocoPresence) {
	db.Lock()
	defer db.Unlock()
	db.killOldsters()
	db.cocoNodes = append(db.cocoNodes, presence)
}

func (db *db) AddMix(presence models.MixNodePresence) {
	db.Lock()
	defer db.Unlock()
	db.killOldsters()
	db.mixNodes = append(db.mixNodes, presence)
}

func (db *db) AddMixProvider(presence models.MixProviderPresence) {
	db.Lock()
	defer db.Unlock()
	db.killOldsters()
	db.mixProviderNodes = append(db.mixProviderNodes, presence)
}

func (db *db) Topology() models.Topology {
	db.killOldsters()
	t := models.Topology{
		CocoNodes:        db.cocoNodes,
		MixNodes:         db.mixNodes,
		MixProviderNodes: db.mixProviderNodes,
	}
	return t
}

// killOldsters kills any stale presence info
func (db *db) killOldsters() {
	for index, presence := range db.mixNodes {
		if presence.LastSeen <= timeWindow() {
			ret := make([]models.MixNodePresence, 0)
			ret = append(ret, db.mixNodes[:index]...)
			db.mixNodes = append(ret, db.mixNodes[index+1:]...)
		}
	}
	for index, presence := range db.cocoNodes {
		if presence.LastSeen <= timeWindow() {
			ret := make([]models.CocoPresence, 0)
			ret = append(ret, db.cocoNodes[:index]...)
			db.cocoNodes = append(ret, db.cocoNodes[index+1:]...)
		}
	}
	for index, presence := range db.mixProviderNodes {
		if presence.LastSeen <= timeWindow() {
			ret := make([]models.MixProviderPresence, 0)
			ret = append(ret, db.mixProviderNodes[:index]...)
			db.mixProviderNodes = append(ret, db.mixProviderNodes[index+1:]...)
		}
	}
}

// timeWindow defines staleness
// TODO: kill magic number by pulling this out into a config
func timeWindow() int64 {
	d := time.Duration(-5)
	return timemock.Now().Add(time.Duration(d * time.Second)).UnixNano()
}
