package mixmining

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/nymtech/nym-directory/models"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/mixmining/fixtures"
	"github.com/nymtech/nym-directory/mixmining/mocks"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Controller", func() {
	Describe("creating a mix status", func() {
		Context("containing xss", func() {
			It("should strip the xss attack and proceed normally", func() {
				router, mockService, mockSanitizer := SetupRouter()

				mockSanitizer.On("Sanitize", fixtures.XSSMixStatus()).Return(fixtures.GoodMixStatus())
				mockService.On("CreateMixStatus", fixtures.GoodMixStatus())
				j, _ := json.Marshal(fixtures.XSSMixStatus())

				resp := performRequest(router, "POST", "/api/mixmining", j)
				var response map[string]string
				json.Unmarshal([]byte(resp.Body.String()), &response)

				assert.Equal(GinkgoT(), 201, resp.Code)
				mockSanitizer.AssertCalled(GinkgoT(), "Sanitize", fixtures.XSSMixStatus())
				mockService.AssertCalled(GinkgoT(), "CreateMixStatus", fixtures.GoodMixStatus())
			})
		})
	})

	Describe("listing statuses for a node", func() {
		Context("when no statuses have yet been saved", func() {
			It("returns an empty list", func() {
				router, mockService, _ := SetupRouter()
				mockService.On("List", "foo").Return([]models.PersistedMixStatus{})
				resp := performRequest(router, "GET", "/api/mixmining/foo", nil)

				assert.Equal(GinkgoT(), 200, resp.Code)
			})

		})
		Context("when some statuses exist", func() {
			It("should return the list of statuses as json", func() {
				router, mockService, _ := SetupRouter()
				mockService.On("List", "pubkey1").Return(fixtures.MixStatusesList())
				url := "/api/mixmining/pubkey1"
				resp := performRequest(router, "GET", url, nil)
				var response []models.PersistedMixStatus
				json.Unmarshal([]byte(resp.Body.String()), &response)

				assert.Equal(GinkgoT(), 200, resp.Code)
				assert.Equal(GinkgoT(), fixtures.MixStatusesList(), response)
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
