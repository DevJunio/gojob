package router

import (
	"github.com/devjunio/gojob/docs"
	"github.com/devjunio/gojob/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gojob API
// @description This is a sample server for Gojob.
// @version 1
// @host 0.0.0.0:8080
// @BasePath /api/v1

// @contact.name   Junio Santos
// @contact.url    https://junio.dev/contact
// @contact.email  contact@junio.dev

// @license.name  GPL 3.0
// @license url https://www.gnu.org/licenses/gpl-3.0.en.html
func setRouteConfig(router *gin.Engine, path ...string) {
	basePath := setDefaultRoute(path, "/api/v1")
	docs.SwaggerInfo.BasePath = basePath
	v0 := router.Group(basePath)
	opService := service.InitService()
	baseRoute := "/openings"
	logger.Info("Initializing routes")

	{
		v0.GET(baseRoute, opService.List)             // Get by ID
		v0.GET(baseRoute+"/:id", opService.Get)       // List all
		v0.POST(baseRoute, opService.Create)          // Create
		v0.DELETE(baseRoute+"/:id", opService.Delete) // Delete
		v0.PATCH(baseRoute+"/:id", opService.Update)  // Update
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	logger.Info("Routes initialized")
	logger.Infof("Listening and serving HTTP on %s", "http://"+Fullpath+basePath)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	logger.Info("Swagger initialized")
	logger.Infof("Swagger HTTP on %s", "http://"+Fullpath+"/swagger")
}
