package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type controller interface {
	RegisterRoutes(*gin.Engine)
}

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
