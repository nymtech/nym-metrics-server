package presence

import (
	"time"

	"github.com/nymtech/directory-server/models"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var (
	presence1 models.MixNodePresence
	presence2 models.MixNodePresence
)

var _ = Describe("Presence Db", func() {
	var db *db
	BeforeEach(func() {
		DbFixtures()
		db = newPresenceDb()
	})
	Describe("constructor", func() {
		It("initializes a db with an empty mixnodes presence map", func() {
			assert.Len(GinkgoT(), db.List(), 0)
		})
	})

	Describe("listing mixnodes", func() {
		Context("when none have been added", func() {
			It("returns an empty map", func() {
				assert.Len(GinkgoT(), db.List(), 0)
			})
		})
		Context("after adding a presence", func() {
			It("returns the map correctly", func() {
				db.Add(presence1)
				assert.Len(GinkgoT(), db.List(), 1)
			})
			It("gets the presence by its public key", func() {
				db.Add(presence1)
				assert.Equal(GinkgoT(), presence1, db.Get(presence1.PubKey))
			})
		})
		Context("after adding two presences", func() {
			It("returns the map correctly", func() {
				db.Add(presence1)
				assert.Len(GinkgoT(), db.List(), 1)
			})
			It("gets the presence by its public key", func() {
				db.Add(presence1)
				assert.Equal(GinkgoT(), presence1, db.Get(presence1.PubKey))
				assert.Equal(GinkgoT(), presence2, db.Get(presence2.PubKey))
			})
		})
	})

})

func DbFixtures() {
	var mix1 = models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "foo.com:8000",
			PubKey: "pubkey1",
		},
		Layer: 1,
	}

	presence1 = models.MixNodePresence{
		MixHostInfo: mix1,
		LastSeen:    time.Now().Unix(),
	}

	var mix2 = models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "bar.com:8000",
			PubKey: "pubkey2",
		},
		Layer: 2,
	}

	presence2 = models.MixNodePresence{
		MixHostInfo: mix2,
		LastSeen:    time.Now().Unix(),
	}
}
