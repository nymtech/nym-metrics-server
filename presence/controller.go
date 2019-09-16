package presence

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/directory-server/models"
)

// controller is the presence controller
type controller struct {
	service *service
}

// Controller is the presence controller interface
type Controller interface {
	AddCocoNodePresence(c *gin.Context)
	AddMixNodePresence(c *gin.Context)
	Topology(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

// New constructor
func New() Controller {
	db := newPresenceDb()
	return &controller{newService(db)}
}

// RegisterRoutes registers controller routes in Gin.
func (controller *controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/presence/coconodes", controller.AddCocoNodePresence)
	router.POST("/api/presence/mixnodes", controller.AddMixNodePresence)
	router.POST("/api/presence/mixproviders", controller.AddMixProviderPresence)
	router.GET("/api/presence/topology", controller.Topology)
}

// AddMixNodePresence lets a mixnode tell the directory server it's alive
// @Summary Lets a node tell the directory server it's alive
// @Description Nym mixnodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.
// @ID notifyMixNode
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   models.MixHostInfo     true  "object"
// @Success 201
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/mixnodes [post]
func (controller *controller) AddMixNodePresence(c *gin.Context) {
	var json models.MixHostInfo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.service.AddMixNodePresence(json)
	c.JSON(http.StatusCreated, gin.H{"ok": true})
}

// AddCocoNodePresence lets a coconut node tell the directory server it's alive
// @Summary Lets a node tell the directory server it's alive
// @Description Nym mixnodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.
// @ID notifyMixNode
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   models.HostInfo     true  "object"
// @Success 201
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/coconodes [post]
func (controller *controller) AddCocoNodePresence(c *gin.Context) {
	var hostInfo models.HostInfo
	if err := c.ShouldBindJSON(&hostInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.service.AddCocoNodePresence(hostInfo)
	c.JSON(http.StatusCreated, gin.H{"ok": true})
}

// AddMixNodePresence lets a mix provider tell the directory server it's alive
// @Summary Lets a node tell the directory server it's alive
// @Description Nym mix provider can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.
// @ID notifyMixNode
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   models.MixProviderHostInfo     true  "object"
// @Success 201
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/mixnodes [post]
func (controller *controller) AddMixProviderPresence(c *gin.Context) {
	var json models.MixProviderHostInfo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.service.AddMixProviderPresence(json)
	c.JSON(http.StatusCreated, gin.H{"ok": true})
}

// Topology lists which Nym nodes are currently known
// @Summary Lists which Nym mixnodes and coconodes are alive
// @Description Nym nodes periodically ping the directory server to register that they're alive. This method provides a list of nodes which have been most recently seen.
// @ID topology
// @Accept  json
// @Produce  json
// @Tags presence
// @Success 200 {object} models.Topology
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/topology [get]
func (controller *controller) Topology(c *gin.Context) {
	topology := controller.service.Topology()
	c.JSON(http.StatusOK, topology)
}
