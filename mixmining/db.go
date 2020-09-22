package mixmining

import "github.com/nymtech/nym-directory/models"

// IDb holds status information
type IDb interface {
	Add(models.PersistedMixStatus)
	List(pubkey string) []models.PersistedMixStatus
}

// Db is a hashtable that holds mixnode uptime mixmining
type Db struct {
	mixStatuses map[string][]models.PersistedMixStatus
}

// NewDb constructor
func NewDb() *Db {
	d := Db{
		mixStatuses: make(map[string][]models.PersistedMixStatus),
	}
	return &d
}

// List returns all models.PersistedMixStatus in the database
func (db *Db) List(pubkey string) []models.PersistedMixStatus {
	if val, ok := db.mixStatuses[pubkey]; ok {
		return val
	}
	return make([]models.PersistedMixStatus, 0)
}

// Add saves a PersistedMixStatus
func (db *Db) Add(metric models.PersistedMixStatus) {
	list := db.mixStatuses[metric.PubKey]
	db.mixStatuses[metric.PubKey] = append(list, metric)
}
