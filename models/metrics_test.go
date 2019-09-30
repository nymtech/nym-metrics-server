package models

import (
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Metrics", func() {
	Describe("sanitizing XSS", func() {
		Context("when no XSS is present", func() {
			It("should not touch input", func() {
				sent := make(map[string]uint)
				sent["foo"] = 1
				received := uint(1)
				metric := MixMetric{
					PubKey:   "pubkey",
					Sent:     sent,
					Received: &received,
				}
				metric.Sanitize()
				assert.Equal(GinkgoT(), metric.PubKey, "pubkey")
				assert.NotNil(GinkgoT(), metric.Sent["foo"])
			})
		})
		Context("when XSS is present", func() {
			It("should strip XSS", func() {
				sent := make(map[string]uint)
				sent["foo<script>alert('gotcha')</script>"] = 1
				received := uint(1)
				metric := MixMetric{
					PubKey:   "pubkey<script>alert('gotcha')</script>",
					Sent:     sent,
					Received: &received,
				}
				metric.Sanitize()
				// assert.Equal(GinkgoT(), metric.Sent["foo"], uint(0))
				assert.Equal(GinkgoT(), "pubkey", metric.PubKey)
			})
		})
	})
})
