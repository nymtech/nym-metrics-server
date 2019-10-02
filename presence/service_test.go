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
		mix1      models.MixHostInfo
		presence1 models.MixNodePresence
		coco1     models.CocoHostInfo
		presence2 models.CocoPresence
		provider1 models.MixProviderHostInfo
		presence3 models.MixProviderPresence
		mockDb    mocks.IDb

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
				Host:   "foo.com:8000",
				PubKey: "pubkey1",
			},
			Layer: 1,
		}

		presence1 = models.MixNodePresence{
			MixHostInfo: mix1,
			LastSeen:    timemock.Now().UnixNano(),
		}

		coco1 = models.CocoHostInfo{
			HostInfo: models.HostInfo{
				Host:   "bar.com:8000",
				PubKey: "pubkey2",
			},
			Type: "foo",
		}
		presence2 = models.CocoPresence{
			CocoHostInfo: coco1,
			LastSeen:     timemock.Now().UnixNano(),
		}

		provider1 = models.MixProviderHostInfo{
			HostInfo: models.HostInfo{
				Host:   "baz.com:8000",
				PubKey: "pubkey2",
			},
			RegisteredClients: []models.RegisteredClient{},
		}

		presence3 = models.MixProviderPresence{
			MixProviderHostInfo: provider1,
			LastSeen:            timemock.Now().UnixNano(),
		}
	})

	Describe("Adding presence info", func() {
		Context("for a mixnode", func() {
			It("should add a presence to the db", func() {
				mockDb.On("AddMix", presence1)
				serv.AddMixNodePresence(mix1, "foo.com")
				mockDb.AssertCalled(GinkgoT(), "AddMix", presence1)
			})
			Context("with a different self-reported IP vs server-reported IP", func() {
				It("should add a presence to the db using the server-reported IP + the self-reported port", func() {
					var presenceToSend = presence1
					presenceToSend.Host = "slammich.com:8000"
					mockDb.On("AddMix", presenceToSend)
					serv.AddMixNodePresence(mix1, "slammich.com")
					mockDb.AssertCalled(GinkgoT(), "AddMix", presenceToSend)
				})
			})
		})
		Context("for a coconode", func() {
			It("should add a presence to the db", func() {
				mockDb.On("AddCoco", presence2)
				serv.AddCocoNodePresence(coco1, "bar.com")
				mockDb.AssertCalled(GinkgoT(), "AddCoco", presence2)
			})
			Context("with a different self-reported IP vs server-reported IP", func() {
				It("should add a presence to the db using the server-reported IP + the self-reported port", func() {
					var presenceToSend = presence2
					presenceToSend.Host = "foo.com:8000"
					mockDb.On("AddCoco", presence2)
					serv.AddCocoNodePresence(coco1, "bar.com")
					mockDb.AssertCalled(GinkgoT(), "AddCoco", presence2)
				})
			})
		})
		Context("for a provider node", func() {
			It("should add a presence to the db", func() {
				mockDb.On("AddMixProvider", presence3)
				serv.AddMixProviderPresence(provider1, "baz.com")
				mockDb.AssertCalled(GinkgoT(), "AddMixProvider", presence3)
			})
			Context("with a different self-reported IP vs server-reported IP", func() {
				It("should add a presence to the db using the server-reported IP + the self-reported port", func() {
					var presenceToSend = presence3
					presenceToSend.Host = "foo.com:8000"
					mockDb.On("AddMixProvider", presenceToSend)
					serv.AddMixProviderPresence(provider1, "foo.com")
					mockDb.AssertCalled(GinkgoT(), "AddMixProvider", presenceToSend)
				})
			})
		})
	})
	Describe("Listing presence info", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				list := []models.MixNodePresence{
					presence1,
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

	Describe("Building the IP of a metrics report", func() {
		Context("from a localhost request", func() {
			It("uses server-reported host with the self-reported port", func() {
				ipa := ipAssigner{}
				result, _ := ipa.AssignIP("localhost", ":8080")
				assert.Equal(GinkgoT(), "localhost:8080", result)
			})
		})
		Context("from a 127.0.0.1 request", func() {
			It("uses server-reported host with the self-reported port", func() {
				ipa := ipAssigner{}
				result, _ := ipa.AssignIP("127.0.0.1", ":8080")
				assert.Equal(GinkgoT(), "127.0.0.1:8080", result)
			})
		})
		Context("from a remote request", func() {
			It("returns server-reported host with the self-reported port", func() {
				ipa := ipAssigner{}
				result, _ := ipa.AssignIP("foo.com", ":8080")
				assert.Equal(GinkgoT(), "foo.com:8080", result)
			})
			Context("with a self-reported host body differing from the server-reported host", func() {
				It("uses server-reported host with the self-reported port", func() {
					ipa := ipAssigner{}
					result, _ := ipa.AssignIP("bar.com", "foo.com:8080")
					assert.Equal(GinkgoT(), "bar.com:8080", result)
				})
			})
		})
	})
})
