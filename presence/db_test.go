package presence

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Presence Db", func() {
	var (
		presence1 models.MixNodePresence
		presence2 models.MixNodePresence
	)
	var db *db
	BeforeEach(func() {
		db = newPresenceDb()

		// Set up fixtures
		var mix1 = models.MixHostInfo{
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

		var mix2 = models.MixHostInfo{
			HostInfo: models.HostInfo{
				Host:   "bar.com:8000",
				PubKey: "pubkey2",
			},
			Layer: 2,
		}
		presence2 = models.MixNodePresence{
			MixHostInfo: mix2,
			LastSeen:    timemock.Now().Unix(),
		}
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
				assert.Equal(GinkgoT(), presence1, db.get(presence1.PubKey))
			})
		})
		Context("after adding two presences", func() {
			It("returns the map correctly", func() {
				db.Add(presence1)
				assert.Len(GinkgoT(), db.List(), 1)
			})
			It("contains the correct presences", func() {
				db.Add(presence1)
				db.Add(presence2)
				assert.Equal(GinkgoT(), presence1, db.get(presence1.PubKey))
				assert.Equal(GinkgoT(), presence2, db.get(presence2.PubKey))
			})
		})
		Describe("Presences", func() {
			Context("more than 5 seconds old", func() {
				It("are not returned by List()", func() {
					oldtime := time.Now().Add(time.Duration(-5 * time.Second)).Unix()
					presence1.LastSeen = oldtime
					db.Add(presence1)
					db.Add(presence2)
					assert.Len(GinkgoT(), db.List(), 1)
					assert.Equal(GinkgoT(), presence2, db.get(presence2.PubKey))
				})
			})
		})
	})
})
