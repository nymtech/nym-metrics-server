package metrics

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Config ...
type Config struct {
	// Db badger.
}

// controller is the PKI controller
type controller struct {
	service *service
}

// Controller is the Key-Value controller
type Controller interface {
	CreateMixMetric(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

// New returns a new pki.Controller
func New(config *Config) Controller {
	return &controller{newService(config)}
}

func (controller *controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/metrics/mixmetric", controller.CreateMixMetric)
}

// CreateMetric adds a node to the PKI
// @Summary Create a metric detailing how many messages a given mixnode sent
// @Description You'd never want to run this in production, but for demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.
// @ID createMetric
// @Accept  json
// @Produce  json
// @Tags metrics
// @Param   object      body   rest.ObjectRequest     true  "object"
// @Success 200 {object} rest.ObjectIDResponse
// @Failure 400 {object} rest.Error
// @Failure 404 {object} rest.Error
// @Failure 500 {object} rest.Error
// @Router /api/metrics [post]
func (controller *controller) CreateMixMetric(c *gin.Context) {
	fmt.Println("foomp")
}
