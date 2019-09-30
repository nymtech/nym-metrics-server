package presence

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/metrics/mocks"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Controller", func() {

})

func SetupRouter() (*gin.Engine, *mocks.IService, *mocks.Sanitizer) {
	mockSanitizer := new(mocks.Sanitizer)
	mockService := new(mocks.IService)

	// cfg := Config{
	// 	Sanitizer: mockSanitizer,
	// 	Service:   mockService,
	// }

	router := gin.Default()

	// controller := New(cfg)
	// controller.RegisterRoutes(router)
	return router, mockService, mockSanitizer
}

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	buf := bytes.NewBuffer(body)
	req, _ := http.NewRequest(method, path, buf)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
