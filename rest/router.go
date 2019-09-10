package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type controller interface {
	RegisterRoutes(*gin.Engine)
}

// @title Nym API
// @version 1.0
// @description Nym REST API endpoints

// @license.name Apache
// @license.url https://github.com/nymtech/directory-server/license

func (s *Server) makeRouter(controllers ...controller) *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Add cors
	router.Use(cors.Default())

	// Serve Swagger frontend static files using gin-swagger middleware
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	for _, v := range controllers {
		v.RegisterRoutes(router)
	}

	return router
}
