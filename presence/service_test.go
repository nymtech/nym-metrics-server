package presence

import (
	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
	"github.com/nymtech/directory-server/presence/mocks"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("presence.Service", func() {
	var (
		mix1      models.MixHostInfo
		presence1 models.MixNodePresence
		mockDb    mocks.Db

		serv service
	)
	BeforeEach(func() {
		mockDb = *new(mocks.Db)
		serv = *newService(&mockDb)

		// Set up fixtures
		mix1 = models.MixHostInfo{
			HostInfo: models.HostInfo{
				Host:   "foo.com:8000",
				PubKey: "pubkey1",
			},
			Layer: 1,
		}

		presence1 = models.MixNodePresence{
			MixHostInfo: mix1,
			LastSeen:    timemock.Now().Unix(),
		}
	})

	Describe("Adding presence info", func() {
		Context("when receiving a mixnode info", func() {
			It("should add a presence to the db", func() {
				mockDb.On("Add", presence1)
				serv.AddMixNodePresence(mix1)
				mockDb.AssertCalled(GinkgoT(), "Add", presence1)
			})
		})
	})
	Describe("Listing presence info", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				list := map[string]models.MixNodePresence{
					presence1.PubKey: presence1,
				}
				mockDb.On("List").Return(list)
				result := serv.Topology()
				mockDb.AssertCalled(GinkgoT(), "List")
				assert.Equal(GinkgoT(), list, result)

			})
		})
	})
})
