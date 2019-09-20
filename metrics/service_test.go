package metrics

import (
	"encoding/json"
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/metrics/mocks"
	"github.com/nymtech/directory-server/models"
	. "github.com/onsi/ginkgo"
	"gotest.tools/assert"

	wsMocks "github.com/nymtech/directory-server/server/websocket/mocks"
)

var _ = Describe("metrics.Service", func() {
	var mockDb mocks.Db
	var m1 models.MixMetric
	var m2 models.MixMetric
	var p1 models.PersistedMixMetric
	var p2 models.PersistedMixMetric

	var serv service
	var received uint = 99
	var now = time.Now()
	timemock.Freeze(now)
	var frozenNow = timemock.Now().UnixNano()

	// set up fixtures
	m1 = models.MixMetric{
		PubKey:   "key1",
		Received: &received,
		Sent:     map[string]uint{"mixnode3": 99, "mixnode4": 101},
	}

	p1 = models.PersistedMixMetric{
		MixMetric: m1,
		Timestamp: frozenNow,
	}

	m2 = models.MixMetric{
		PubKey:   "key2",
		Received: &received,
		Sent:     map[string]uint{"mixnode3": 102, "mixnode4": 103},
	}

	p2 = models.PersistedMixMetric{
		MixMetric: m2,
		Timestamp: frozenNow,
	}

	Describe("Adding a mixmetric", func() {
		It("should add a PersistedMixMetric to the db and notify the Hub", func() {
			mockDb = *new(mocks.Db)
			mockHub := *new(wsMocks.Broadcaster)
			serv = *newService(&mockDb, &mockHub)
			mockDb.On("Add", p1)
			j, _ := json.Marshal(p1)
			mockHub.On("Notify", j)

			serv.CreateMixMetric(m1)

			mockDb.AssertCalled(GinkgoT(), "Add", p1)
			mockHub.AssertCalled(GinkgoT(), "Notify", j)
		})
	})
	Describe("Listing mixmetrics", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				mockDb = *new(mocks.Db)
				mockHub := *new(wsMocks.Broadcaster)

				list := []models.PersistedMixMetric{p1, p2}

				serv = *newService(&mockDb, &mockHub)
				mockDb.On("List").Return(list)

				result := serv.List()

				mockDb.AssertCalled(GinkgoT(), "List")
				assert.Equal(GinkgoT(), list[0].MixMetric.PubKey, result[0].MixMetric.PubKey)
				assert.Equal(GinkgoT(), list[1].MixMetric.PubKey, result[1].MixMetric.PubKey)
			})
		})
	})
})
