package presence

import (
	"time"

	"github.com/nymtech/directory-server/models"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var (
	mix1 models.MixHostInfo
	mix2 models.MixHostInfo

	serv     service
	initTime int64
)

func CreateFixtures() {
	mix1 = models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "foo.com:8000",
			PubKey: "pubkey1",
		},
		Layer: 1,
	}

	mix2 = models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "bar.com:8000",
			PubKey: "pubkey2",
		},
		Layer: 2,
	}
}

var _ = Describe("presence.Service", func() {

	BeforeEach(func() {
		CreateFixtures()
		initTime = time.Now().Unix()
		serv = *newService()

	})

	Describe("Network topology", func() {
		Context("At service construction", func() {
			It("should be empty", func() {
				assert.Empty(GinkgoT(), serv.Topology())
			})
		})
		Context("When no nodes have been added yet", func() {
			It("should add the mixnode to the mixnodes list", func() {
				serv.AddMixNodePresence(mix1)
				assert.Len(GinkgoT(), serv.Topology(), 1)
				assert.Equal(GinkgoT(), mix1.HostInfo, serv.mixNodes[0].HostInfo)
			})
			It("should include a unix timestamp greater than when we started", func() {
				serv.AddMixNodePresence(mix1)
				assert.True(GinkgoT(), serv.mixNodes[0].LastSeen >= initTime)
			})
		})
		Context("When 2 nodes are added", func() {
			It("should add the mixnodes to the mixnodes list", func() {
				serv.AddMixNodePresence(mix1)
				serv.AddMixNodePresence(mix2)
				assert.Len(GinkgoT(), serv.Topology(), 2)
			})
		})
	})

	Describe("Getting mixnet topology", func() {
		Context("when there are old topology reports in the list (older than 5 seconds)", func() {
			It("should return the list stripped of old presence reports", func() {
				assert.Len(GinkgoT(), serv.Topology(), 0)
			})
		})
	})
})
