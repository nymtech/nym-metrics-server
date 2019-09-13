package presence

import (
	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
	"github.com/nymtech/directory-server/presence/mocks"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var (
	mix1   models.MixHostInfo
	mix2   models.MixHostInfo
	pres1  models.MixNodePresence
	mockDb mocks.Db

	serv     service
	initTime int64
)

var _ = Describe("presence.Service", func() {

	BeforeEach(func() {
		ServiceFixtures()
		mockDb = *new(mocks.Db)
		serv = *newService(&mockDb)
	})

	Describe("Adding presence info", func() {
		Context("when receiving a mixnode info", func() {
			It("should add a presence to the db", func() {
				mockDb.On("Add", pres1)
				serv.AddMixNodePresence(mix1)
				mockDb.AssertCalled(GinkgoT(), "Add", pres1)
			})
		})
	})
	Describe("Listing presence info", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				list := map[string]models.MixNodePresence{
					pres1.PubKey: pres1,
				}
				mockDb.On("List").Return(list)
				result := serv.Topology()
				mockDb.AssertCalled(GinkgoT(), "List")
				assert.Equal(GinkgoT(), list, result)

			})
		})
	})
})

func ServiceFixtures() {
	mix1 = models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "foo.com:8000",
			PubKey: "pubkey1",
		},
		Layer: 1,
	}

	pres1 = models.MixNodePresence{
		MixHostInfo: mix1,
		LastSeen:    timemock.Now().Unix(),
	}
}
