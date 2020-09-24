package mixmining

import (
	"github.com/jinzhu/gorm"
	"github.com/nymtech/nym-directory/models"

	// needed for Gorm to get its sqlite dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is the Gorm orm for mixmining
var DB *gorm.DB

// IDb holds status information
type IDb interface {
	Add(models.PersistedMixStatus)
	List(pubkey string, limit int) []models.PersistedMixStatus
	SaveMixStatusReport(models.MixStatusReport)
	ListDateRange(pubkey string, start int, end int) []models.PersistedMixStatus
}

// Db is a hashtable that holds mixnode uptime mixmining
type Db struct {
	orm *gorm.DB
}

// NewDb constructor
func NewDb() *Db {
	database, err := gorm.Open("sqlite3", "nym-mixmining.db")

	if err != nil {
		panic("Failed to connect to orm!")
	}

	database.AutoMigrate(&models.PersistedMixStatus{})
	d := Db{
		database,
	}
	return &d
}

// Add saves a PersistedMixStatus
func (db *Db) Add(status models.PersistedMixStatus) {
	db.orm.Create(status)
}

// List returns all models.PersistedMixStatus in the orm
func (db *Db) List(pubkey string, limit int) []models.PersistedMixStatus {
	var statuses []models.PersistedMixStatus
	if err := db.orm.Order("timestamp desc").Limit(limit).Where("pub_key = ?", pubkey).Find(&statuses).Error; err != nil {
		return make([]models.PersistedMixStatus, 0)
	}
	return statuses
}

// ListDateRange lists all persisted mix statuses for a node within the specified date range
func (db *Db) ListDateRange(pubkey string, start int, end int) []models.PersistedMixStatus {
	var statuses []models.PersistedMixStatus
	if err := db.orm.Order("timestamp desc").Where("pub_key = ?", pubkey).Where("timestamp >= ?", start).Where("timestamp <= ?", end).Find(&statuses).Error; err != nil {
		return make([]models.PersistedMixStatus, 0)
	}
	return statuses
}

// SaveMixStatusReport creates or updates a status summary report for a given mixnode in the database
func (db *Db) SaveMixStatusReport(report models.MixStatusReport) {

}
