package mixmining

import (
	"github.com/nymtech/nym-directory/mixmining/fixtures"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("The mixmining db", func() {
	Describe("Constructing a NewDb", func() {
		Context("a new db", func() {
			It("should have no mixmining", func() {
				db := NewDb()
				assert.Len(GinkgoT(), db.List("foo"), 0)
			})
		})
	})

	Describe("adding and retrieving measurements", func() {
		Context("a new db", func() {
			It("should add measurements to the db, with a timestamp, and be able to retrieve them afterwards", func() {
				db := NewDb()
				status := fixtures.GoodPersistedMixStatus()

				// add one
				db.Add(status)
				measurements := db.List(status.PubKey)
				assert.Len(GinkgoT(), measurements, 1)
				assert.Equal(GinkgoT(), status, measurements[0])

				// add another
				db.Add(status)
				measurements = db.List(status.PubKey)
				assert.Len(GinkgoT(), measurements, 2)
				assert.Equal(GinkgoT(), status, measurements[0])
				assert.Equal(GinkgoT(), status, measurements[1])

			})
		})
	})

	Describe("listing mixmining", func() {
		Context("for an empty db", func() {
			It("should return an empty slice", func() {
				db := NewDb()
				assert.Len(GinkgoT(), db.List("foo"), 0)
			})
		})
	})

})
