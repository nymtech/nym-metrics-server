package metrics

import (
	"log"

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
	router.POST("/api/metrics/mixes", controller.CreateMixMetric)
	router.GET("/api/metrics/mixes", controller.CreateMixMetric)
}

// CreateMixMetric adds a node to the PKI
// @Summary Create a metric detailing how many messages a given mixnode sent
// @Description You'd never want to run this in production, but for demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.
// @ID createMixMetric
// @Accept  json
// @Produce  json
// @Tags metrics
// @Param   object      body   models.ObjectRequest     true  "object"
// @Success 200 {object} models.ObjectIDResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/metrics/mixes [post]
func (controller *controller) CreateMixMetric(c *gin.Context) {
	log.Println("CreateMixMetric not yet implemented")
}

// ListMixMetrics lists mixnode activity
// @Summary Lists mixnode activity in the past 1 second
// @Description You'd never want to run this in production, but for demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.
// @ID listMixMetrics
// @Accept  json
// @Produce  json
// @Tags metrics
// @Param   object      body   models.ObjectRequest     true  "object"
// @Success 200 {object} models.ObjectIDResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/metrics/mixes [get]
func (controller *controller) ListMixMetrics(c *gin.Context) {
	log.Println("ListMixMetrics not yet implemented")
}
