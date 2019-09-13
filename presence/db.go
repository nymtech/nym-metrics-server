package presence

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
)

// Db holds presence information
type Db interface {
	Add(models.MixNodePresence)
	List() map[string]models.MixNodePresence
}

type db struct {
	mixNodes map[string]models.MixNodePresence
}

func newPresenceDb() *db {
	return &db{
		mixNodes: map[string]models.MixNodePresence{},
	}
}

func (db *db) Add(presence models.MixNodePresence) {
	db.killOldsters()
	db.mixNodes[presence.PubKey] = presence
}

func (db *db) List() map[string]models.MixNodePresence {
	db.killOldsters()
	return db.mixNodes
}

func (db *db) get(key string) models.MixNodePresence {
	return db.mixNodes[key]
}

// kill any stale presence info (older than 5 seconds)
// TODO: kill magic number 5 by pulling this out into a config
func (db *db) killOldsters() {
	for key := range db.mixNodes {
		presence := db.mixNodes[key]
		if presence.LastSeen <= window(5) {
			delete(db.mixNodes, key)
		}
	}
}

// defines staleness
func window(seconds int64) int64 {
	d := time.Duration(-1 * seconds)
	return timemock.Now().Add(time.Duration(d * time.Second)).Unix()
}
