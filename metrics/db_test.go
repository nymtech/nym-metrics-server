package metrics

import (
	"github.com/nymtech/directory-server/models"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Metrics Db", func() {
	var db *db
	var metric1 models.MixMetric
	var metric2 models.MixMetric

	// set up fixtures
	metric1 = models.MixMetric{
		PubKey:   "key1",
		Received: 99,
		Sent:     map[string]uint{"mixnode3": 99, "mixnode4": 101},
	}

	metric2 = models.MixMetric{
		PubKey:   "key2",
		Received: 99,
		Sent:     map[string]uint{"mixnode3": 102, "mixnode4": 103},
	}

	Describe("retrieving mixnet metrics", func() {
		Context("when no metrics have been added", func() {
			It("should return an empty metrics list", func() {
				db = newMetricsDb()
				assert.Len(GinkgoT(), db.List(), 0)
			})
		})
	})
	Describe("adding mixnet metrics", func() {
		Context("adding 1", func() {
			It("should contain 1 metric", func() {
				db = newMetricsDb()
				db.Add(metric1)
				assert.Len(GinkgoT(), db.List(), 1)
			})
		})
		Context("adding 2", func() {
			It("should contain 2 metrics", func() {
				db = newMetricsDb()
				db.Add(metric1)
				db.Add(metric2)
				assert.Len(GinkgoT(), db.List(), 2)
			})
		})
	})
})
