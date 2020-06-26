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
	AddGateway(models.GatewayPresence)
	Allow(string)
	Disallow(string)
	ListDisallowed() []string
	Topology() models.Topology
}

type db struct {
	// TODO: it's slightly inefficient to have a single mutex for all database, because right now
	// if a mix node was being added, we wouldn't be able to touch cocoNodes
	sync.Mutex
	cocoNodes        map[string]models.CocoPresence
	disallowed       []string
	mixNodes         map[string]models.MixNodePresence
	mixProviderNodes map[string]models.MixProviderPresence
	gateways         map[string]models.GatewayPresence
}

// NewDb constructor...
func NewDb() *db {
	return &db{
		cocoNodes:        map[string]models.CocoPresence{},
		disallowed:       []string{},
		mixNodes:         map[string]models.MixNodePresence{},
		mixProviderNodes: map[string]models.MixProviderPresence{},
		gateways:         map[string]models.GatewayPresence{},
	}
}

func (db *db) AddCoco(presence models.CocoPresence) {
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

func (db *db) AddGateway(presence models.GatewayPresence) {
	db.Lock()
	defer db.Unlock()
	db.killOldsters()
	db.gateways[presence.IdentityKey] = presence
}

func (db *db) Allow(pubkey string) {
	db.Lock()
	defer db.Unlock()
	db.disallowed = remove(db.disallowed, pubkey)
}

func (db *db) Disallow(pubkey string) {
	db.Lock()
	defer db.Unlock()
	for _, disallowedNode := range db.disallowed {
		if disallowedNode == pubkey {
			return
		}
	}
	db.disallowed = append(db.disallowed, pubkey)
}

func (db *db) ListDisallowed() []string {
	return db.disallowed
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
	gatewayNodes := []models.GatewayPresence{}

	for _, value := range db.cocoNodes {
		cocoNodes = append(cocoNodes, value)
	}

	for _, value := range db.mixNodes {
		mixNodes = append(mixNodes, value)
	}

	for _, value := range db.mixProviderNodes {
		mixProviderNodes = append(mixProviderNodes, value)
	}

	for _, value := range db.gateways {
		gatewayNodes = append(gatewayNodes, value)
	}

	t := models.Topology{
		CocoNodes:        cocoNodes,
		MixNodes:         mixNodes,
		MixProviderNodes: mixProviderNodes,
		Gateways:         gatewayNodes,
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
	for key, presence := range db.gateways {
		if presence.LastSeen <= timeWindow() {
			delete(db.gateways, key)
		}
	}
}

// timeWindow defines staleness
// TODO: kill magic number by pulling this out into a config
func timeWindow() int64 {
	d := time.Duration(-10)
	return timemock.Now().Add(time.Duration(d * time.Second)).UnixNano()
}

func remove(list []string, item string) []string {
	for i, other := range list {
		if other == item {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}
