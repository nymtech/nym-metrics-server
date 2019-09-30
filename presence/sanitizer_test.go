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
			It("sanitizes input", func() {
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
	Context("for MixProviderHostInfo", func() {
		Context("when XSS is present", func() {
			It("sanitizes input", func() {
				policy := bluemonday.UGCPolicy()
				sanitizer := NewMixproviderSanitizer(policy)

				result := sanitizer.Sanitize(xssMixProviderHost())
				assert.Equal(GinkgoT(), goodMixProviderHost(), result)
			})
		})
		Context("when XSS is not present", func() {
			It("doesn't change input", func() {
				policy := bluemonday.UGCPolicy()
				sanitizer := NewMixproviderSanitizer(policy)
				result := sanitizer.Sanitize(goodMixProviderHost())
				assert.Equal(GinkgoT(), goodMixProviderHost(), result)
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

func goodMixProviderHost() models.MixProviderHostInfo {
	client1 := models.RegisteredClient{PubKey: "client1"}
	client2 := models.RegisteredClient{PubKey: "client2"}
	clients := []models.RegisteredClient{client1, client2}
	good := models.MixProviderHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host",
			PubKey: "pubkey",
		},
		RegisteredClients: clients,
	}
	return good
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

func xssMixHost() models.MixHostInfo {
	xss := models.MixHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host<script>alert('gotcha')</script>",
			PubKey: "pubkey<script>alert('gotcha')</script>",
		},
	}
	return xss
}

func xssMixProviderHost() models.MixProviderHostInfo {
	client1 := models.RegisteredClient{PubKey: "client1<script>alert('gotcha')</script>"}
	client2 := models.RegisteredClient{PubKey: "client2<script>alert('gotcha')</script>"}
	clients := []models.RegisteredClient{client1, client2}
	xss := models.MixProviderHostInfo{
		HostInfo: models.HostInfo{
			Host:   "host<script>alert('gotcha')</script>",
			PubKey: "pubkey<script>alert('gotcha')</script>",
		},
		RegisteredClients: clients,
	}
	return xss
}
