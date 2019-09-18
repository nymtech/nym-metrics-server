package presence

import (
	"sync"
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
)

// Db holds presence information
type Db interface {
	AddCoco(models.Presence)
	AddMix(models.MixNodePresence)
	AddMixProvider(models.MixProviderPresence)
	Topology() models.Topology
}

type db struct {
	// TODO: it's slightly inefficient to have a single mutex for all database, because right now
	// if a mix node was being added, we wouldn't be able to touch cocoNodes
	sync.Mutex
	cocoNodes        map[string]models.Presence
	mixNodes         map[string]models.MixNodePresence
	mixProviderNodes map[string]models.MixProviderPresence
}

func newPresenceDb() *db {
	return &db{
		cocoNodes:        map[string]models.Presence{},
		mixNodes:         map[string]models.MixNodePresence{},
		mixProviderNodes: map[string]models.MixProviderPresence{},
	}
}

func (db *db) AddCoco(presence models.Presence) {
	db.Lock()
	defer db.Unlock()
	db.killOldsters()
	db.cocoNodes[presence.PubKey] = presence
}

func (db *db) AddMix(presence models.MixNodePresence) {
	db.Lock()
	defer db.Unlock()
	db.killOldsters()
	db.mixNodes[presence.PubKey] = presence
}

func (db *db) AddMixProvider(presence models.MixProviderPresence) {
	db.Lock()
	defer db.Unlock()
	db.killOldsters()
	db.mixProviderNodes[presence.PubKey] = presence
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
	for key, presence := range db.mixNodes {
		if presence.LastSeen <= timeWindow() {
			delete(db.mixNodes, key)
		}
	}
	for key, presence := range db.cocoNodes {
		if presence.LastSeen <= timeWindow() {
			delete(db.cocoNodes, key)
		}
	}
	for key, presence := range db.mixProviderNodes {
		if presence.LastSeen <= timeWindow() {
			delete(db.mixProviderNodes, key)
		}
	}
}

// timeWindow defines staleness
// TODO: kill magic number by pulling this out into a config
func timeWindow() int64 {
	d := time.Duration(-5)
	return timemock.Now().Add(time.Duration(d * time.Second)).UnixNano()
}
