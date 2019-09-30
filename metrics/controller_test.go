package metrics

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/metrics/mocks"
	. "github.com/onsi/ginkgo"
	"gotest.tools/assert"
)

var _ = Describe("MetricsController", func() {
	Describe("creating a metric", func() {
		Context("containing xss", func() {
			It("should strip the xss attack", func() {
				router, controller := SetupRouter()
				_ = controller

				json, _ := json.Marshal(xssMetric())
				resp := performRequest(router, "POST", "/api/metrics/mixes", json)
				assert.Equal(GinkgoT(), 201, resp.Code)
			})
		})
	})
	Describe("listing metrics", func() {
		Context("when no metrics exist", func() {
			It("should return an empty list", func() {

			})
		})
		Context("when metrics exist", func() {
			It("should return them", func() {

			})
		})
	})
})

func SetupRouter() (*gin.Engine, Controller) {
	mockSanitizer := new(mocks.Sanitizer)
	mockService := new(mocks.IService)

	mockSanitizer.On("Sanitize", xssMetric()).Return(goodMetric())
	mockService.On("CreateMixMetric", goodMetric())

	metricsConfig := Config{
		Sanitizer: mockSanitizer,
		Service:   mockService,
	}

	router := gin.Default()

	controller := New(metricsConfig)
	controller.RegisterRoutes(router)
	return router, controller
}

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	buf := bytes.NewBuffer(body)
	req, _ := http.NewRequest(method, path, buf)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
