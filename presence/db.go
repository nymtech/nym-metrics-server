package presence

import (
	"sync"
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/nym-directory/models"
)

// IDb holds presence information
type IDb interface {
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

// Topology returns the full network Topology
//
// This implementation is a little bit involved, and you might wonder why we
// don't simply make the db fields into slices (instead of maps) and get rid of
// all this map-to-slice conversion code. The answer is that the maps nicely
// overwrite the keyed value for a host even if multiple updates for a single
// host are received within the timeWindow. So we get a nice bag of presences,
// without duplicates, and don't have to worry much about timing. The tradeoff
// is some extra code here:
func (db *db) Topology() models.Topology {
	db.killOldsters()

	cocoNodes := []models.CocoPresence{}
	mixNodes := []models.MixNodePresence{}
	mixProviderNodes := []models.MixProviderPresence{}

	for _, value := range db.cocoNodes {
		cocoNodes = append(cocoNodes, value)
	}

	for _, value := range db.mixNodes {
		mixNodes = append(mixNodes, value)
	}

	for _, value := range db.mixProviderNodes {
		mixProviderNodes = append(mixProviderNodes, value)
	}

	t := models.Topology{
		CocoNodes:        cocoNodes,
		MixNodes:         mixNodes,
		MixProviderNodes: mixProviderNodes,
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
