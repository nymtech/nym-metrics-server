package mixmining

import (
	"github.com/nymtech/nym-directory/mixmining/fixtures"
	"github.com/nymtech/nym-directory/models"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("The mixmining db", func() {
	Describe("Constructing a NewDb", func() {
		Context("a new db", func() {
			It("should have no mixmining", func() {
				db := NewDb()
				db.orm.Exec("DELETE FROM persisted_mix_statuses")
				assert.Len(GinkgoT(), db.List("foo", 5), 0)
			})
		})
	})

	Describe("adding and retrieving measurements", func() {
		Context("a new db", func() {
			It("should add measurements to the db, with a timestamp, and be able to retrieve them afterwards", func() {
				db := NewDb()
				db.orm.Exec("DELETE FROM persisted_mix_statuses")
				status := fixtures.GoodPersistedMixStatus()

				// add one
				db.Add(status)
				measurements := db.List(status.PubKey, 5)
				assert.Len(GinkgoT(), measurements, 1)
				assert.Equal(GinkgoT(), status, measurements[0])

				// add another
				db.Add(status)
				measurements = db.List(status.PubKey, 5)
				assert.Len(GinkgoT(), measurements, 2)
				assert.Equal(GinkgoT(), status, measurements[0])
				assert.Equal(GinkgoT(), status, measurements[1])

			})
		})
	})

	Describe("listing mix statuses within a date range", func() {
		Context("for an empty db", func() {
			It("should return an empty slice", func() {
				db := NewDb()
				db.orm.Exec("DELETE FROM persisted_mix_statuses")
				assert.Len(GinkgoT(), db.ListDateRange("foo", 1, 1), 0)
			})
		})
		Context("when one status exists in the range and one outside", func() {
			It("should return only the status within the range", func() {
				db := NewDb()
				db.orm.Exec("DELETE FROM persisted_mix_statuses")
				data := fixtures.GoodMixStatus()
				statusInRange := models.PersistedMixStatus{
					MixStatus: data,
					Timestamp: 500,
				}
				statusOutOfRange := models.PersistedMixStatus{
					MixStatus: data,
					Timestamp: 1000,
				}
				db.Add(statusInRange)
				db.Add(statusOutOfRange)

				result := db.ListDateRange(data.PubKey, 0, 500)
				assert.Len(GinkgoT(), result, 1)
				assert.Equal(GinkgoT(), statusInRange, result[0])
			})
		})
	})

	Describe("listing mix statuses with a limit", func() {
		Context("for an empty db", func() {
			It("should return an empty slice", func() {
				db := NewDb()
				db.orm.Exec("DELETE FROM persisted_mix_statuses")
				assert.Len(GinkgoT(), db.List("foo", 5), 0)
			})
		})
	})

})
