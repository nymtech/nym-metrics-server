package metrics

import (
	"github.com/nymtech/directory-server/metrics/mocks"
	"github.com/nymtech/directory-server/models"
	. "github.com/onsi/ginkgo"
	"gotest.tools/assert"
)

var _ = Describe("metrics.Service", func() {
	var mockDb mocks.Db
	var metric1 models.MixMetric
	var metric2 models.MixMetric

	var serv service

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

	Describe("Adding a mixmetric", func() {
		It("should add to the db", func() {
			mockDb = *new(mocks.Db)
			serv = *newService(&mockDb)
			mockDb.On("Add", metric1)

			serv.CreateMixMetric(metric1)
			mockDb.AssertCalled(GinkgoT(), "Add", metric1)
		})
	})
	Describe("Listing mixmetrics", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				mockDb = *new(mocks.Db)
				list := []models.MixMetric{metric1, metric2}

				serv = *newService(&mockDb)
				mockDb.On("List").Return(list)
				result := serv.List()
				mockDb.AssertCalled(GinkgoT(), "List")
				assert.Equal(GinkgoT(), list[0].PubKey, result[0].PubKey)
				assert.Equal(GinkgoT(), list[1].PubKey, result[1].PubKey)
			})
		})
	})
})
