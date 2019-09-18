package metrics

import (
	"time"

	"github.com/nymtech/directory-server/metrics/mocks"
	"github.com/nymtech/directory-server/models"
	. "github.com/onsi/ginkgo"
	"gotest.tools/assert"
)

var _ = Describe("metrics.Service", func() {
	var mockDb mocks.Db
	var m1 models.MixMetric
	var m2 models.MixMetric
	var p1 models.PersistedMixMetric
	var p2 models.PersistedMixMetric

	var serv service
	var received uint = 99
	var now int64 = time.Now().UnixNano()

	// set up fixtures
	m1 = models.MixMetric{
		PubKey:   "key1",
		Received: &received,
		Sent:     map[string]uint{"mixnode3": 99, "mixnode4": 101},
	}

	p1 = models.PersistedMixMetric{
		MixMetric: m1,
		Timestamp: now,
	}

	m2 = models.MixMetric{
		PubKey:   "key2",
		Received: &received,
		Sent:     map[string]uint{"mixnode3": 102, "mixnode4": 103},
	}

	p2 = models.PersistedMixMetric{
		MixMetric: m2,
		Timestamp: now,
	}

	Describe("Adding a mixmetric", func() {
		It("should add a PersistedMixMetric to the db", func() {
			mockDb = *new(mocks.Db)
			serv = *newService(&mockDb)
			mockDb.On("Add", p1)

			serv.CreateMixMetric(m1)
			mockDb.AssertCalled(GinkgoT(), "Add", p1)
		})
	})
	Describe("Listing mixmetrics", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				mockDb = *new(mocks.Db)
				list := []models.PersistedMixMetric{p1, p2}

				serv = *newService(&mockDb)
				mockDb.On("List").Return(list)
				result := serv.List()
				mockDb.AssertCalled(GinkgoT(), "List")
				assert.Equal(GinkgoT(), list[0].MixMetric.PubKey, result[0].MixMetric.PubKey)
				assert.Equal(GinkgoT(), list[1].MixMetric.PubKey, result[1].MixMetric.PubKey)
			})
		})
	})
})
