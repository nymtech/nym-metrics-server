package presence

import (
	"github.com/nymtech/directory-server/models"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("presence.Service", func() {

	var (
		hostInfo1 models.MixHostInfo
		hostInfo2 models.MixHostInfo
		serv      service
	)

	BeforeEach(func() {
		serv = *newService()
		hi := models.HostInfo{
			Host:   "foo.com",
			PubKey: "pubkey",
		}
		hostInfo1 = models.MixHostInfo{
			HostInfo: hi,
			Layer:    1,
		}
		hostInfo2 := hostInfo1
		hostInfo2.Host = "bar.com"
		hostInfo2.Layer = 2
	})

	Describe("Notifying mixnode presence", func() {
		Context("At service construction", func() {
			It("should have an empty mixNodes list", func() {
				assert.Equal(GinkgoT(), 0, len(serv.mixNodes))
			})
		})
		Context("When no nodes have been added yet", func() {
			It("should add the mixnode to the mixnodes list", func() {
				serv.AddMixNodePresence(hostInfo1)
				assert.Equal(GinkgoT(), 1, len(serv.mixNodes))
			})
		})
		Context("When 2 nodes are added", func() {
			It("should add the mixnode to the mixnodes list", func() {
				serv.AddMixNodePresence(hostInfo1)
				serv.AddMixNodePresence(hostInfo2)
				assert.Equal(GinkgoT(), 2, len(serv.mixNodes))
			})
		})
	})

	Describe("Getting mixnet topology", func() {

	})
})
