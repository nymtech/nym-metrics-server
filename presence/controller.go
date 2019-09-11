package presence

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/directory-server/models"
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
	NotifyMixNodePresence(c *gin.Context)
	Up(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

// New returns a new pki.Controller
func New(config *Config) Controller {
	return &controller{newService(config)}
}

func (controller *controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/presence/mixnodes", controller.NotifyMixNodePresence)
	router.GET("/api/presence/mixnodes", controller.Up)
}

// NotifyMixNodePresence lets a node tell the directory server it's alive
// @Summary Lets a node tell the directory server it's alive
// @Description Nym mixnodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.
// @ID notifyMixNode
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   models.UpMsg     true  "object"
// @Success 201
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/mixnodes [post]
func (controller *controller) NotifyMixNodePresence(c *gin.Context) {
	var json models.UpMsg
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.service.NotifyMixNodePresence(json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ok": true})
}

// Up lists which Nym nodes are currently known
// @Summary Lists which Nym mixnodes are alive
// @Description Nym mixnodes periodically ping the directory server to register that they're alive. This method provides a list of mixnodes which have been most recently seen.
// @ID mixNodesUp
// @Accept  json
// @Produce  json
// @Tags presence
// @Success 200 {array} models.Presence
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/mixnodes [get]
func (controller *controller) Up(c *gin.Context) {
	presence, err := controller.service.Up()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{err.Error()})
		return
	}
	c.JSON(http.StatusOK, presence)
}
