package presence

import "github.com/nymtech/directory-server/models"

// Db holds presence information
type Db interface {
	Add()
	Get()
	List()
}

type db struct {
	mixNodes map[string]models.MixNodePresence
}

func newPresenceDb() *db {
	return &db{
		mixNodes: map[string]models.MixNodePresence{},
	}
}

func (db db) Add(presence models.MixNodePresence) {
	db.mixNodes[presence.PubKey] = presence
}

func (db db) Get(key string) models.MixNodePresence {
	return db.mixNodes[key]
}

func (db db) List() map[string]models.MixNodePresence {
	return db.mixNodes
}
