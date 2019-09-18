package server

import (
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/directory-server/healthcheck"
	"github.com/nymtech/directory-server/metrics"
	"github.com/nymtech/directory-server/presence"
	"github.com/nymtech/directory-server/server/websocket"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New returns a new REST API server
func New() *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)

	// Add cors middleware
	router.Use(cors.Default())

	// Serve Swagger frontend static files using gin-swagger middleware
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register all HTTP controller routes
	healthcheck.New().RegisterRoutes(router)
	metrics.New().RegisterRoutes(router)
	presence.New().RegisterRoutes(router)

	// Set up websocket handlers
	hub := websocket.NewHub()
	go hub.Run()

	router.GET("/ws", func(c *gin.Context) {
		websocket.Serve(hub, c.Writer, c.Request)
	})

	return router
}

// loadTemplate loads templates embedded by go-assets-builder
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
