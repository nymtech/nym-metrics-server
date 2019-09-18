package pki

import (
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
	// CreateNode(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

// New returns a new pki.Controller
func New(config *Config) Controller {
	return &controller{newService(config)}
}

func (controller *controller) RegisterRoutes(router *gin.Engine) {
	// router.POST("/api/nodes", controller.CreateNode)
}

// CreateNode adds a node to the PKI
// @Summary Create a node in the PKI
// @Description Nodes should post their public key info to this method when they start.
// @ID createObject
// @Accept  json
// @Produce  json
// @Tags pki
// Param   object      body   models.ObjectRequest     true  "object"
// Success 200 {object} models.ObjectIDResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/nodes [post]
// func (controller *controller) CreateNode(c *gin.Context) {
// 	log.Println("CreateNode not yet implemented")
// }
