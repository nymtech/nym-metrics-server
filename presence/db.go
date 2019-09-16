package presence

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
)

// Db holds presence information
type Db interface {
	AddMix(models.MixNodePresence)
	Topology() models.Topology
}

type db struct {
	cocoNodes map[string]models.Presence
	mixNodes  map[string]models.MixNodePresence
}

func newPresenceDb() *db {
	return &db{
		cocoNodes: map[string]models.Presence{},
		mixNodes:  map[string]models.MixNodePresence{},
	}
}

func (db *db) AddCocoNode(presence models.Presence) {
	db.cocoNodes[presence.PubKey] = presence
}

func (db *db) AddMix(presence models.MixNodePresence) {
	db.killOldsters()
	db.mixNodes[presence.PubKey] = presence
}

func (db *db) Topology() models.Topology {
	db.killOldsters()
	t := models.Topology{
		CocoNodes: db.cocoNodes,
		MixNodes:  db.mixNodes,
	}
	return t
}

// kill any stale presence info (older than 5 seconds)
func (db *db) killOldsters() {
	for key := range db.mixNodes {
		presence := db.mixNodes[key]
		if presence.LastSeen <= timeWindow() {
			delete(db.mixNodes, key)
		}
	}
	// for key := range db.cocoNodes {
	// 	presence := db.cocoNodes[key]
	// 	if presence.LastSeen <= timeWindow() {
	// 		delete(db.cocoNodes, key)
	// 	}
	// }
}

// defines staleness
// TODO: kill magic number by pulling this out into a config
func timeWindow() int64 {
	d := time.Duration(-5)
	return timemock.Now().Add(time.Duration(d * time.Second)).Unix()
}
