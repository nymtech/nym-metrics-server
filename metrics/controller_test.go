package metrics

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/metrics/mocks"
	"github.com/nymtech/nym-directory/models"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("MetricsController", func() {
	Describe("creating a metric", func() {
		Context("containing xss", func() {
			It("should strip the xss attack", func() {
				router, mockService, mockSanitizer := SetupRouter()
				mockSanitizer.On("Sanitize", xssMetric()).Return(goodMetric())
				mockService.On("CreateMixMetric", goodMetric())
				json, _ := json.Marshal(xssMetric())

				resp := performRequest(router, "POST", "/api/metrics/mixes", json)

				assert.Equal(GinkgoT(), 201, resp.Code)
				mockSanitizer.AssertCalled(GinkgoT(), "Sanitize", xssMetric())
				mockService.AssertCalled(GinkgoT(), "CreateMixMetric", goodMetric())
			})
		})
	})
	Describe("listing metrics", func() {
		Context("when no metrics exist", func() {
			It("should return an empty list of metrics", func() {
				router, mockService, mockSanitizer := SetupRouter()
				_ = mockSanitizer // nothing to sanitize here
				mockService.On("List").Return([]models.PersistedMixMetric{})

				resp := performRequest(router, "GET", "/api/metrics/mixes", nil)

				assert.Equal(GinkgoT(), 200, resp.Code)
				mockService.AssertExpectations(GinkgoT())
			})
		})
		Context("when metrics exist", func() {
			It("should return them", func() {
				router, mockService, mockSanitizer := SetupRouter()
				_ = mockSanitizer // nothing to sanitize here

				mockService.On("List").Return(mixMetricsList())

				resp := performRequest(router, "GET", "/api/metrics/mixes", nil)
				var response []models.PersistedMixMetric
				json.Unmarshal([]byte(resp.Body.String()), &response)

				assert.Equal(GinkgoT(), 200, resp.Code)
				assert.Equal(GinkgoT(), mixMetricsList(), response)
				mockService.AssertExpectations(GinkgoT())
			})
		})
	})
})

func SetupRouter() (*gin.Engine, *mocks.IService, *mocks.Sanitizer) {
	mockSanitizer := new(mocks.Sanitizer)
	mockService := new(mocks.IService)

	metricsConfig := Config{
		Sanitizer: mockSanitizer,
		Service:   mockService,
	}

	router := gin.Default()

	controller := New(metricsConfig)
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

// fixtures

func mixMetricsList() []models.PersistedMixMetric {
	r1 := uint(1)
	m1 := models.PersistedMixMetric{
		MixMetric: models.MixMetric{
			PubKey:   "pubkey1",
			Received: &r1,
		},
	}

	r2 := uint(2)
	m2 := models.PersistedMixMetric{
		MixMetric: models.MixMetric{
			PubKey:   "pubkey2",
			Received: &r2,
		},
	}

	metrics := []models.PersistedMixMetric{m1, m2}
	return metrics
}
