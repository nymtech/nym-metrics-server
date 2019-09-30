package presence

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/nymtech/nym-directory/models"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Sanitizer", func() {
	Describe("sanitizing inputs", func() {
		Context("for CocoHostInfo", func() {
			Context("when XSS is present", func() {
				It("sanitizes input", func() {
					policy := bluemonday.UGCPolicy()
					sanitizer := NewCoconodeSanitizer(policy)

					result := sanitizer.Sanitize(xssCocoHost())
					assert.Equal(GinkgoT(), goodCocoHost(), result)
				})
			})
			Context("when XSS is not present", func() {
				It("doesn't change input", func() {
					policy := bluemonday.UGCPolicy()
					sanitizer := NewCoconodeSanitizer(policy)
					result := sanitizer.Sanitize(goodCocoHost())
					assert.Equal(GinkgoT(), goodCocoHost(), result)
				})
			})
		})
	})
	Context("for MixHostInfo", func() {
		Context("when XSS is present", func() {
			FIt("sanitizes input", func() {
				policy := bluemonday.UGCPolicy()
				sanitizer := NewMixnodeSanitizer(policy)

				result := sanitizer.Sanitize(xssMixHost())
				assert.Equal(GinkgoT(), goodHost(), result)
			})
		})
		Context("when XSS is not present", func() {
			It("doesn't change input", func() {
				policy := bluemonday.UGCPolicy()
				sanitizer := NewMixnodeSanitizer(policy)
				result := sanitizer.Sanitize(goodHost())
				assert.Equal(GinkgoT(), goodHost(), result)
			})
		})
	})
})

func goodCocoHost() models.CocoHostInfo {
	good := models.CocoHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host",
			PubKey: "pubkey",
		},
		Type: "type",
	}
	return good
}

func goodHost() models.MixHostInfo {
	good := models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host",
			PubKey: "pubkey",
		},
	}
	return good
}

func xssMixHost() models.MixHostInfo {
	xss := models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host<script>alert('gotcha')",
			PubKey: "pubkey<script>alert('gotcha')",
		},
	}
	return xss
}

func xssCocoHost() models.CocoHostInfo {
	xss := models.CocoHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host<script>alert('gotcha')",
			PubKey: "pubkey<script>alert('gotcha')",
		},
		Type: "type<script>alert('gotcha')",
	}
	return xss
}
