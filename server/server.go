package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/directory-server/healthcheck"
	"github.com/nymtech/directory-server/metrics"
	"github.com/nymtech/directory-server/presence"
	"github.com/nymtech/directory-server/server/html"
	"github.com/nymtech/directory-server/server/websocket"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New returns a new REST API server
func New() *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Add cors middleware
	router.Use(cors.Default())

	// Serve Swagger frontend static files using gin-swagger middleware
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Add HTML templates to the router
	t, err := html.LoadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/server/html/index.html", nil)
	})

	// Set up websocket handlers
	hub := websocket.NewHub()
	go hub.Run()

	router.GET("/ws", func(c *gin.Context) {
		websocket.Serve(hub, c.Writer, c.Request)
	})

	cfg := metrics.Config{
		Hub: hub,
	}

	// Register all HTTP controller routes
	healthcheck.New().RegisterRoutes(router)
	metrics.New(cfg).RegisterRoutes(router)
	presence.New().RegisterRoutes(router)

	return router
}
