package metrics

import (
	"sync"
	"time"

	"github.com/nymtech/directory-server/models"
)

type db struct {
	sync.Mutex
	mixMetrics []models.MixMetric
	ticker     *time.Ticker
}

// Db holds metrics information
type Db interface {
	Add(models.MixMetric)
	List() []models.MixMetric
}

func newMetricsDb() *db {
	ticker := time.NewTicker(10 * time.Second)

	d := db{
		mixMetrics: []models.MixMetric{},
	}
	d.ticker = ticker
	go dbCleaner(ticker, &d)

	return &d
}

// Add adds a models.MixMetric to the database
func (db *db) Add(metric models.MixMetric) {
	db.Lock()
	defer db.Unlock()
	db.mixMetrics = append(db.mixMetrics, metric)
}

// List returns all models.MixMetric in the database
func (db *db) List() []models.MixMetric {
	db.Lock()
	defer db.Unlock()
	return db.mixMetrics
}

// dbCleaner periodically clears the database
func dbCleaner(ticker *time.Ticker, database *db) {
	for {
		select {
		case <-ticker.C:
			database.clear()
		}
	}
}

// clear kills any stale presence info
func (db *db) clear() {
	db.Lock()
	defer db.Unlock()
	db.mixMetrics = db.mixMetrics[:0]
}
