package measurements

import "github.com/nymtech/nym-directory/models"

// IDb holds status information
type IDb interface {
	Add(models.PersistedMixStatus)
	List() []models.PersistedMixStatus
}

// Db is a hashtable that holds mixnode uptime measurements
type Db struct {
}

// NewDb constructor
func NewDb() *Db {
	d := Db{}
	return &d
}

// List returns all models.PersistedMixStatus in the database
func (db *Db) List() []models.PersistedMixStatus {
	return make([]models.PersistedMixStatus, 0)
}

// Add saves a PersistedMixStatus
func (db *Db) Add(metric models.PersistedMixStatus) {

}
