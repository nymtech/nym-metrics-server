package mixmining

import (
	"fmt"

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

	Describe("adding and retrieving one measurement", func() {
		Context("a new db", func() {
			It("should add one measurement to the db, with a timestamp", func() {
				db := NewDb()
				status := fixtures.GoodPersistedMixStatus()
				db.Add(status)
				measurements := db.List(status.PubKey)
				fmt.Printf("mixmining: %+v", measurements)
				assert.Len(GinkgoT(), measurements, 1)
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
