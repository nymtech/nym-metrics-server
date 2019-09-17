package metrics

import (
	"sync"

	"github.com/nymtech/directory-server/models"
)

type db struct {
	sync.Mutex
	mixMetrics []models.MixMetric
}

// Db holds presence information
type Db interface {
	Add(models.MixMetric)
	List() []models.MixMetric
}

func newMetricsDb() *db {
	return &db{
		mixMetrics: []models.MixMetric{},
	}
}

func (db *db) Add(metric models.MixMetric) {
	db.Lock()
	defer db.Unlock()
	db.mixMetrics = append(db.mixMetrics, metric)
}

func (db *db) List() []models.MixMetric {
	db.Lock()
	defer db.Unlock()
	return db.mixMetrics
}
