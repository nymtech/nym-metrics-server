package presence

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Config ...
type Config struct {
	// Db badger.
}

// controller is the presence controller
type controller struct {
	service *service
}

// Controller is the presence controller
type Controller interface {
	NotifyPresence(c *gin.Context)
	Up(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

// New returns a new pki.Controller
func New(config *Config) Controller {
	return &controller{newService(config)}
}

func (controller *controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/presence", controller.NotifyPresence)
	router.GET("/api/presence/up", controller.Up)
}

// NotifyPresence lets a node tell the directory server it's alive
// @Summary Lets a node tell the directory server it's alive
// @Description Nym coconodes and mixnodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.
// @ID notifyPresence
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   rest.ObjectRequest     true  "object"
// @Success 200 {object} rest.ObjectIDResponse
// @Failure 400 {object} rest.Error
// @Failure 404 {object} rest.Error
// @Failure 500 {object} rest.Error
// @Router /api/presence [post]
func (controller *controller) NotifyPresence(c *gin.Context) {
	fmt.Println("foomp")
}

// Up lists which Nym nodes are currently known
// @Summary Lists which Nym nodes are alive
// @Description Nym coconodes and mixnodes periodically ping the directory server to register that they're alive. This method provides a list of nodes which have been most recently seen.
// @ID up
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   rest.ObjectRequest     true  "object"
// @Success 200 {object} rest.ObjectIDResponse
// @Failure 400 {object} rest.Error
// @Failure 404 {object} rest.Error
// @Failure 500 {object} rest.Error
// @Router /api/presence/up [get]
func (controller *controller) Up(c *gin.Context) {
	fmt.Println("up")
}
