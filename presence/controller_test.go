package presence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/presence/fixtures"
	"github.com/nymtech/nym-directory/presence/mocks"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Presence Controller", func() {
	Describe("creating a mix node presence", func() {
		Context("containing xss", func() {
			It("should strip the xss attack", func() {
				router, mockService, mockSanitizer := SetupRouter()
				mockSanitizer.On("Sanitize", fixtures.XssMixHost()).Return(fixtures.GoodHost())
				mockService.On("AddMixNodePresence", fixtures.GoodHost())
				j, _ := json.Marshal(fixtures.XssMixHost())

				resp := performRequest(router, "POST", "/api/presence/mixnodes", j)
				var response map[string]string
				json.Unmarshal([]byte(resp.Body.String()), &response)
				fmt.Printf("RESPONSE: %v", response)

				assert.Equal(GinkgoT(), 201, resp.Code)
				mockSanitizer.AssertCalled(GinkgoT(), "Sanitize", fixtures.XssMixHost())
				mockService.AssertCalled(GinkgoT(), "AddMixNodePresence", fixtures.GoodHost())
			})
		})
	})
})

func SetupRouter() (*gin.Engine, *mocks.IService, *mocks.Sanitizer) {
	mockSanitizer := new(mocks.Sanitizer)
	mockService := new(mocks.IService)

	cfg := Config{
		Sanitizer: mockSanitizer,
		Service:   mockService,
	}

	router := gin.Default()

	controller := New(cfg)
	controller.RegisterRoutes(router)
	return router, mockService, mockSanitizer
}

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	buf := bytes.NewBuffer(body)
	req, _ := http.NewRequest(method, path, buf)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
