package measurements

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/measurements/fixtures"
	"github.com/nymtech/nym-directory/measurements/mocks"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Controller", func() {
	Describe("creating a mix status", func() {
		Context("containing xss", func() {
			It("should strip the xss attack and proceed normally", func() {
				mockSanitizer := new(mocks.Sanitizer)
				mockService := new(mocks.IService)

				cfg := Config{
					Sanitizer: mockSanitizer,
					Service:   mockService,
				}

				router := gin.Default()

				controller := New(cfg)
				controller.RegisterRoutes(router)
				mockSanitizer.On("Sanitize", fixtures.XSSMixStatus()).Return(fixtures.GoodMixStatus())
				mockService.On("CreateMixStatus", fixtures.GoodMixStatus())
				j, _ := json.Marshal(fixtures.XSSMixStatus())

				resp := performRequest(router, "POST", "/api/measurements", j)
				var response map[string]string
				json.Unmarshal([]byte(resp.Body.String()), &response)

				assert.Equal(GinkgoT(), 201, resp.Code)
				mockSanitizer.AssertCalled(GinkgoT(), "Sanitize", fixtures.XSSMixStatus())
				mockService.AssertCalled(GinkgoT(), "CreateMixStatus", fixtures.GoodMixStatus())
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
