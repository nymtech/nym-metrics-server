package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nymtech/directory-server/healthcheck"
	"github.com/nymtech/directory-server/metrics"
	"github.com/nymtech/directory-server/presence"

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

	// Register all the controller routes
	healthcheck.New().RegisterRoutes(router)
	metrics.New().RegisterRoutes(router)
	presence.New().RegisterRoutes(router)

	return router
}
