package presence

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/constants"
	"github.com/nymtech/nym-directory/models"
)

// Config for this controller
type Config struct {
	CocoHostSanitizer        CocoHostSanitizer
	MixHostSanitizer         MixHostSanitizer
	MixProviderHostSanitizer MixProviderHostSanitizer
	Service                  IService
}

// controller is the presence controller
type controller struct {
	service           IService
	cocoHostSanitizer CocoHostSanitizer
	mixHostSanitizer  MixHostSanitizer
}

// Controller is the presence controller interface
type Controller interface {
	AddCocoNodePresence(c *gin.Context)
	AddMixNodePresence(c *gin.Context)
	Topology(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

// New constructor
func New(cfg Config) Controller {
	return &controller{
		cfg.Service,
		cfg.CocoHostSanitizer,
		cfg.MixHostSanitizer,
	}
}

// RegisterRoutes registers controller routes in Gin.
func (controller *controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/presence/coconodes", controller.AddCocoNodePresence)
	router.POST("/api/presence/mixnodes", controller.AddMixNodePresence)
	router.POST("/api/presence/mixproviders", controller.AddMixProviderPresence)
	router.GET("/api/presence/topology", controller.Topology)
}

// AddMixNodePresence ...
// @Summary Lets mixnode a node tell the directory server it's alive
// @Description Nym mixnodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.
// @ID addMixNode
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
	var mixHost models.MixHostInfo
	if err := c.ShouldBindJSON(&mixHost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ip, _, err := net.SplitHostPort(mixHost.Host)
	if (ip == "localhost" || net.ParseIP(ip).IsLoopback()) && err == nil {
		// keep host info we received
	} else {
		mixHost.HostInfo.Host = net.JoinHostPort(c.ClientIP(), constants.DefaultMixPort)
	}
	sanitized := controller.mixHostSanitizer.Sanitize(mixHost)
	controller.service.AddMixNodePresence(sanitized)
	c.JSON(http.StatusCreated, gin.H{"ok": true})
}

// AddCocoNodePresence ...
// @Summary Lets a coconut node tell the directory server it's alive
// @Description Nym Coconut nodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.
// @ID addCocoNode
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   models.CocoHostInfo     true  "object"
// @Success 201
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/coconodes [post]
func (controller *controller) AddCocoNodePresence(c *gin.Context) {
	var cocoHost models.CocoHostInfo
	if err := c.ShouldBindJSON(&cocoHost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ip, _, err := net.SplitHostPort(cocoHost.Host)
	if (ip == "localhost" || net.ParseIP(ip).IsLoopback()) && err == nil {
		// keep host info we received
	} else {
		cocoHost.Host = net.JoinHostPort(c.ClientIP(), constants.DefaultMixPort)
	}
	sanitized := controller.cocoHostSanitizer.Sanitize(cocoHost)
	controller.service.AddCocoNodePresence(sanitized)
	c.JSON(http.StatusCreated, gin.H{"ok": true})
}

// AddMixNodePresence ...
// @Summary Lets a mixnode tell the directory server it's alive
// @Description Nym mix providers can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.
// @ID addMixProvider
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   models.MixProviderHostInfo     true  "object"
// @Success 201
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/providers [post]
func (controller *controller) AddMixProviderPresence(c *gin.Context) {
	var json models.MixProviderHostInfo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ip, _, err := net.SplitHostPort(json.Host)
	if (ip == "localhost" || net.ParseIP(ip).IsLoopback()) && err == nil {
		// keep host info we received
	} else {
		json.HostInfo.Host = net.JoinHostPort(c.ClientIP(), constants.DefaultMixPort)
	}
	controller.service.AddMixProviderPresence(json)
	c.JSON(http.StatusCreated, gin.H{"ok": true})
}

// Topology ...
// @Summary Lists which Nym mixnodes, providers, and coconodes are alive
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
