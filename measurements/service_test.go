package measurements

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/nym-directory/measurements/mocks"
	"github.com/nymtech/nym-directory/models"
	. "github.com/onsi/ginkgo"
	"gotest.tools/assert"
)

var _ = Describe("metrics.Service", func() {
	var mockDb mocks.IDb
	var m1 models.MixStatus
	var m2 models.MixStatus
	var p1 models.PersistedMixStatus
	var p2 models.PersistedMixStatus

	var serv Service
	var now = time.Now()
	timemock.Freeze(now)
	var frozenNow = timemock.Now().UnixNano()

	// set up fixtures
	m1 = models.MixStatus{
		PubKey:    "key1",
		IPVersion: "4",
		Up:        true,
	}

	p1 = models.PersistedMixStatus{
		MixStatus: m1,
		Timestamp: frozenNow,
	}

	m2 = models.MixStatus{
		PubKey:    "key2",
		IPVersion: "6",
		Up:        false,
	}

	p2 = models.PersistedMixStatus{
		MixStatus: m2,
		Timestamp: frozenNow,
	}

	Describe("Adding a mix status", func() {
		It("should add a PersistedMixStatus to the db", func() {
			mockDb = *new(mocks.IDb)
			serv = *NewService(&mockDb)
			mockDb.On("Add", p1)

			serv.CreateMixStatus(m1)

			mockDb.AssertCalled(GinkgoT(), "Add", p1)
		})
	})
	Describe("Listing mix statuses", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				mockDb = *new(mocks.IDb)

				list := []models.PersistedMixStatus{p1, p2}

				serv = *NewService(&mockDb)
				mockDb.On("List").Return(list)

				result := serv.List()

				mockDb.AssertCalled(GinkgoT(), "List")
				assert.Equal(GinkgoT(), list[0].MixStatus.PubKey, result[0].MixStatus.PubKey)
				assert.Equal(GinkgoT(), list[1].MixStatus.PubKey, result[1].MixStatus.PubKey)
			})
		})
	})
})
