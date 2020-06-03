package presence

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/nym-directory/models"
	"github.com/nymtech/nym-directory/presence/mocks"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("presence.Service", func() {
	var (
		mix1              models.MixHostInfo
		mixpresence1      models.MixNodePresence
		coco1             models.CocoHostInfo
		cocopresence2     models.CocoPresence
		provider1         models.MixProviderHostInfo
		providerpresence3 models.MixProviderPresence
		mockDb            mocks.IDb

		serv service
	)
	BeforeEach(func() {
		mockDb = *new(mocks.IDb)
		serv = *NewService(&mockDb)
		var now = time.Now()
		timemock.Freeze(now)

		// Set up fixtures
		mix1 = models.MixHostInfo{
			HostInfo: models.HostInfo{
				Host:     "foo.com:8000",
				PubKey:   "pubkey1",
				Location: defaultLocation,
			},
			Layer: 1,
		}

		mixpresence1 = models.MixNodePresence{
			MixHostInfo: mix1,
			LastSeen:    timemock.Now().UnixNano(),
		}

		coco1 = models.CocoHostInfo{
			HostInfo: models.HostInfo{
				Host:     "bar.com:8000",
				PubKey:   "pubkey2",
				Location: defaultLocation,
			},
			Type: "foo",
		}
		cocopresence2 = models.CocoPresence{
			CocoHostInfo: coco1,
			LastSeen:     timemock.Now().UnixNano(),
		}

		provider1 = models.MixProviderHostInfo{
			MixnetListener:    "foo.com:8000",
			ClientListener:    "foo.com:8001",
			Location:          defaultLocation,
			PubKey:            "pubkey2",
			RegisteredClients: []models.RegisteredClient{},
		}

		providerpresence3 = models.MixProviderPresence{
			MixProviderHostInfo: provider1,
			LastSeen:            timemock.Now().UnixNano(),
		}
	})

	Describe("Adding presence info", func() {
		Context("for a mixnode", func() {
			It("should add a presence to the db", func() {
				mockDb.On("AddMix", mixpresence1)
				serv.AddMixNodePresence(mix1)
				mockDb.AssertCalled(GinkgoT(), "AddMix", mixpresence1)
			})
		})
		Context("for a coconode", func() {
			It("should add a presence to the db", func() {
				mockDb.On("AddCoco", cocopresence2)
				serv.AddCocoNodePresence(coco1, "bar.com")
				mockDb.AssertCalled(GinkgoT(), "AddCoco", cocopresence2)
			})
		})
		Context("for a provider node", func() {
			It("should add a presence to the db", func() {
				mockDb.On("AddMixProvider", providerpresence3)
				serv.AddMixProviderPresence(provider1)
				mockDb.AssertCalled(GinkgoT(), "AddMixProvider", providerpresence3)
			})
		})
	})
	Describe("Getting the Topology", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				list := []models.MixNodePresence{
					mixpresence1,
				}
				topology := models.Topology{
					MixNodes: list,
				}
				mockDb.On("Topology").Return(topology)
				result := serv.Topology()
				mockDb.AssertCalled(GinkgoT(), "Topology")
				assert.Equal(GinkgoT(), topology, result)
			})
		})

				}
				topology := models.Topology{
					MixNodes: list,
				}
				mockDb.On("Topology").Return(topology)
				result := serv.Topology()
				mockDb.AssertCalled(GinkgoT(), "Topology")
				assert.Equal(GinkgoT(), topology, result)
			})
		})
	})
})
