package presence

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/presence/fixtures"
	"github.com/nymtech/nym-directory/presence/mocks"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Presence Controller", func() {
	Describe("creating a coconode presence", func() {
		Context("containing xss", func() {
			It("should strip the xss attack", func() {
				cocoSan := new(mocks.CocoHostSanitizer)
				mockService := new(mocks.IService)

				cfg := Config{
					CocoHostSanitizer: cocoSan,
					Service:           mockService,
				}

				router := gin.Default()

				controller := New(cfg)
				controller.RegisterRoutes(router)
				cocoSan.On("Sanitize", fixtures.XssCocoHost()).Return(fixtures.GoodCocoHost())
				mockService.On("AddCocoNodePresence", fixtures.GoodCocoHost(), "")
				j, _ := json.Marshal(fixtures.XssCocoHost())

				resp := performRequest(router, "POST", "/api/presence/coconodes", j)
				var response map[string]string
				json.Unmarshal([]byte(resp.Body.String()), &response)

				assert.Equal(GinkgoT(), 201, resp.Code)
				cocoSan.AssertCalled(GinkgoT(), "Sanitize", fixtures.XssCocoHost())
				mockService.AssertCalled(GinkgoT(), "AddCocoNodePresence", fixtures.GoodCocoHost(), "")
			})
		})
	})

	Describe("creating a mix node presence", func() {
		Context("containing xss", func() {
			It("should strip the xss attack", func() {
				mockSanitizer := new(mocks.MixHostSanitizer)
				mockService := new(mocks.IService)

				cfg := Config{
					MixHostSanitizer: mockSanitizer,
					Service:          mockService,
				}

				router := gin.Default()

				controller := New(cfg)
				controller.RegisterRoutes(router)

				mockSanitizer.On("Sanitize", fixtures.XssMixHost()).Return(fixtures.GoodMixHost())
				mockService.On("AddMixNodePresence", fixtures.GoodMixHost())
				j, _ := json.Marshal(fixtures.XssMixHost())

				resp := performRequest(router, "POST", "/api/presence/mixnodes", j)
				var response map[string]string
				json.Unmarshal([]byte(resp.Body.String()), &response)

				assert.Equal(GinkgoT(), 201, resp.Code)
				mockSanitizer.AssertCalled(GinkgoT(), "Sanitize", fixtures.XssMixHost())
				mockService.AssertCalled(GinkgoT(), "AddMixNodePresence", fixtures.GoodMixHost())
			})
		})
	})

	Describe("creating a mix provider node presence", func() {
		Context("containing xss", func() {
			It("should strip the xss attack", func() {
				mockSanitizer := new(mocks.MixProviderHostSanitizer)
				mockService := new(mocks.IService)

				cfg := Config{
					MixProviderHostSanitizer: mockSanitizer,
					Service:                  mockService,
				}

				router := gin.Default()

				controller := New(cfg)
				controller.RegisterRoutes(router)

				mockSanitizer.On("Sanitize", fixtures.XssMixProviderHost()).Return(fixtures.GoodMixProviderHost())
				mockService.On("AddMixProviderPresence", fixtures.GoodMixProviderHost())
				j, _ := json.Marshal(fixtures.XssMixProviderHost())

				resp := performRequest(router, "POST", "/api/presence/mixproviders", j)
				var response map[string]string
				json.Unmarshal([]byte(resp.Body.String()), &response)

				assert.Equal(GinkgoT(), 201, resp.Code)
				mockSanitizer.AssertCalled(GinkgoT(), "Sanitize", fixtures.XssMixProviderHost())
				mockService.AssertCalled(GinkgoT(), "AddMixProviderPresence", fixtures.GoodMixProviderHost())
			})
		})
	})

	Describe("disallowing a node", func() {
		Context("with a properly formatted node key", func() {
			It("should tell the service to disallow the node", func() {
				mockSanitizer := new(mocks.MixProviderHostSanitizer)
				mockService := new(mocks.IService)

				cfg := Config{
					MixProviderHostSanitizer: mockSanitizer,
					Service:                  mockService,
				}

				router := gin.Default()

				controller := New(cfg)
				controller.RegisterRoutes(router)

				hostKey, _ := json.Marshal(fixtures.Disallow())
				mockService.On("Disallow", fixtures.Disallow())

				resp := performRequest(router, "POST", "/api/presence/disallow", hostKey)
				var response map[string]string
				json.Unmarshal([]byte(resp.Body.String()), &response)

				assert.Equal(GinkgoT(), 201, resp.Code)
			})
		})
	})
})

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	buf := bytes.NewBuffer(body)
	req, _ := http.NewRequest(method, path, buf)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
