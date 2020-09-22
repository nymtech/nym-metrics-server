package measurements

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/models"
)

// Config for this controller
type Config struct {
	Sanitizer Sanitizer
	Service   IService
}

// controller is the metrics controller
type controller struct {
	service   IService
	sanitizer Sanitizer
}

// Controller ...
type Controller interface {
	CreateMixStatus(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

// New returns a new measurements.Controller
func New(cfg Config) Controller {
	return &controller{cfg.Service, cfg.Sanitizer}
}

func (controller *controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/measurements", controller.CreateMixStatus)
	router.GET("/api/measurements", controller.ListMeasurements)
}

// ListMeasurements lists mixnode statuses
// @Summary Lists mixnode activity
// @Description Lists all mixnode statuses
// @ID listMixMeasurements
// @Accept  json
// @Produce  json
// @Tags metrics
// Param   object      body   models.ObjectRequest     true  "object"
// @Success 200 {array} models.MixMetric
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/measurements [get]
func (controller *controller) ListMeasurements(c *gin.Context) {
	measurements := controller.service.List()
	c.JSON(http.StatusOK, measurements)
}

// CreateMixStatus ...
// @Summary Lets the network monitor create a new uptime status for a mix
// @Description Nym network monitor sends packets through the system and checks if they make it. The network monitor then hits this method to report whether the node was up at a given time.
// @ID addMixStatus
// @Accept  json
// @Produce  json
// @Tags measurements
// @Param   object      body   models.MixStatus     true  "object"
// @Success 201
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/measurements [post]
func (controller *controller) CreateMixStatus(c *gin.Context) {
	var status models.MixStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sanitized := controller.sanitizer.Sanitize(status)
	controller.service.CreateMixStatus(sanitized)
	c.JSON(http.StatusCreated, gin.H{"ok": true})
}

func deserialize(body []byte) {

}
