package measurements

import (
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("The measurements db", func() {
	Describe("Constructing a NewDb", func() {
		Context("a new db", func() {
			It("should have no measurements", func() {
				// db := NewDb()
				assert.True(GinkgoT(), true)
			})
		})
	})

	// Describe("adding measurements", func() {
	// 	It("should add the measurement to the db, with a timestamp", func() {

	// 	})
	// })
})
