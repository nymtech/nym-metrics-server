package metrics

import "github.com/nymtech/directory-server/models"

type db struct {
	mixMetrics []models.MixMetric
}

func newMetricsDb() *db {
	return &db{
		mixMetrics: []models.MixMetric{},
	}
}

func (db *db) Add(metric models.MixMetric) {
	db.mixMetrics = append(db.mixMetrics, metric)
}

func (db *db) List() []models.MixMetric {
	return db.mixMetrics
}
