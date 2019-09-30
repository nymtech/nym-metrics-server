package metrics

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/nymtech/nym-directory/metrics/mocks"
	"github.com/nymtech/nym-directory/models"
	. "github.com/onsi/ginkgo"
	"gotest.tools/assert"
)

var _ = Describe("MetricsController", func() {
	Describe("creating a metric", func() {
		Context("containing xss", func() {
			It("should strip the xss attack", func() {
				sent := make(map[string]uint)
				sent["foo<script>alert('gotcha')</script>"] = 1
				received := uint(1)
				badMetric := models.MixMetric{
					PubKey:   "pubkey<script>alert('gotcha')</script>",
					Sent:     sent,
					Received: &received,
				}

				router, controller := SetupRouter()
				// controller.
				_ = controller

				json, _ := json.Marshal(badMetric)
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
	sanitizer := bluemonday.UGCPolicy()
	mockService := new(mocks.IService)

	mockService.On("CreateMixMetric", badMetric()).Return("pubkey")

	metricsConfig := Config{
		Sanitizer: *sanitizer,
		Service:   mockService,
	}

	router := gin.Default()

	controller := New(metricsConfig)
	controller.RegisterRoutes(router)
	return router, controller
}

func badMetric() models.MixMetric {
	sent := make(map[string]uint)
	sent["foo<script>alert('gotcha')</script>"] = 1
	received := uint(1)
	m := models.MixMetric{
		PubKey:   "bar<script>alert('gotcha')</script>",
		Sent:     sent,
		Received: &received,
	}
	return m
}

func sanitizedMetric() models.MixMetric {
	sent := make(map[string]uint)
	sent["foo"] = 1
	received := uint(1)
	m := models.MixMetric{
		PubKey:   "bar",
		Sent:     sent,
		Received: &received,
	}
	return m
}

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	buf := bytes.NewBuffer(body)
	req, _ := http.NewRequest(method, path, buf)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
