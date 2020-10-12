package mixmining

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/nymtech/nym-directory/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the Gorm orm for mixmining
var DB *gorm.DB

// IDb holds status information
type IDb interface {
	Add(models.PersistedMixStatus)
	BatchAdd(status []models.PersistedMixStatus)
	List(pubkey string, limit int) []models.PersistedMixStatus
	ListDateRange(pubkey string, ipVersion string, start int64, end int64) []models.PersistedMixStatus
	LoadReport(pubkey string) models.MixStatusReport
	LoadNonStaleReports() models.BatchMixStatusReport
	BatchLoadReports(pubkeys []string) models.BatchMixStatusReport
	SaveMixStatusReport(models.MixStatusReport)
	SaveBatchMixStatusReport(models.BatchMixStatusReport)
}

// Db is a hashtable that holds mixnode uptime mixmining
type Db struct {
	orm *gorm.DB
}

// NewDb constructor
func NewDb() *Db {
	database, err := gorm.Open(sqlite.Open(dbPath()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to orm!")
	}

	database.AutoMigrate(&models.PersistedMixStatus{})
	database.AutoMigrate(&models.MixStatusReport{})

	d := Db{
		database,
	}
	return &d
}

func dbPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := usr.HomeDir + "/.nym/"
	os.MkdirAll(dbPath, os.ModePerm)
	db := dbPath + "mixmining.db"
	fmt.Printf("db is: %s\n", db)
	return db
}

// Add saves a PersistedMixStatus
func (db *Db) Add(status models.PersistedMixStatus) {
	db.orm.Create(status)
}

// BatchAdd saves multiple PersistedMixStatus
func (db *Db) BatchAdd(status []models.PersistedMixStatus) {
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

// ListDateRange lists all persisted mix statuses for a node for either IPv4 or IPv6 within the specified date range
func (db *Db) ListDateRange(pubkey string, ipVersion string, start int64, end int64) []models.PersistedMixStatus {
	var statuses []models.PersistedMixStatus
	if err := db.orm.Order("timestamp desc").Where("pub_key = ?", pubkey).Where("ip_version = ?", ipVersion).Where("timestamp >= ?", start).Where("timestamp <= ?", end).Find(&statuses).Error; err != nil {
		return make([]models.PersistedMixStatus, 0)
	}
	return statuses
}

// SaveMixStatusReport creates or updates a status summary report for a given mixnode in the database
func (db *Db) SaveMixStatusReport(report models.MixStatusReport) {
	create := db.orm.Save(report)
	if create.Error != nil {
		fmt.Printf("Mix status report creation error: %+v", create.Error)
	}
}

// SaveBatchMixStatusReport creates or updates a status summary report for multiple mixnodex in the database
func (db *Db) SaveBatchMixStatusReport(report models.BatchMixStatusReport) {
	if result := db.orm.Save(report.Report); result.Error != nil {
		fmt.Printf("Batch Mix status report save error: %+v", result.Error)
	}
}

// LoadReport retrieves a models.MixStatusReport.
// If a report isn't found, it crudely generates a new instance and returns that instead.
func (db *Db) LoadReport(pubkey string) models.MixStatusReport {
	var report models.MixStatusReport

	if retrieve := db.orm.First(&report, "pub_key  = ?", pubkey); retrieve.Error != nil {
		fmt.Printf("ERROR while retrieving mix status report %+v", retrieve.Error)
		return models.MixStatusReport{}
	}
	return report
}

// LoadNonStaleReports retrieves a models.BatchMixStatusReport, such that each mixnode
// in the retrieved report must have been online for over 50% of time in the last day.
// If a report isn't found, it crudely generates a new instance and returns that instead.
func (db *Db) LoadNonStaleReports() models.BatchMixStatusReport {
	var reports []models.MixStatusReport

	if retrieve := db.orm.Where("last_day_ip_v4 >= 50").Or("last_day_ip_v6 >= 50").Find(&reports); retrieve.Error != nil {
		fmt.Printf("ERROR while retrieving multiple mix status report %+v", retrieve.Error)
		return models.BatchMixStatusReport{Report: make([]models.MixStatusReport, 0)}
	}
	return models.BatchMixStatusReport{Report: reports}
}

// BatchLoadReports retrieves a models.BatchMixStatusReport based on provided set of public keys.
// If a report isn't found, it crudely generates a new instance and returns that instead.
func (db *Db) BatchLoadReports(pubkeys []string) models.BatchMixStatusReport {
	var reports []models.MixStatusReport

	if retrieve := db.orm.Where("pub_key IN ?", pubkeys).Find(&reports); retrieve.Error != nil {
		fmt.Printf("ERROR while retrieving multiple mix status report %+v", retrieve.Error)
		return models.BatchMixStatusReport{Report: make([]models.MixStatusReport, 0)}
	}
	return models.BatchMixStatusReport{Report: reports}
}
