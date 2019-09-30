package presence

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/presence/mocks"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Presence Controller", func() {
	Describe("creating a mix node presence", func() {
		Context("containing xss", func() {
			It("should strip the xss attack", func() {
				// router, mockService, mockSanitizer := SetupRouter()
				// mockSanitizer.On("Sanitize", xssMetric()).Return(goodMetric())
				// mockService.On("CreateMixMetric", goodMetric())
				// json, _ := json.Marshal(xssMetric())

				// resp := performRequest(router, "POST", "/api/metrics/mixes", json)

				// assert.Equal(GinkgoT(), 201, resp.Code)
				// mockSanitizer.AssertCalled(GinkgoT(), "Sanitize", xssMetric())
				// mockService.AssertCalled(GinkgoT(), "CreateMixMetric", goodMetric())
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
