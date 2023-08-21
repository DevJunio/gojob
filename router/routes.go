package router

import (
	"github.com/devjunio/gojob/docs"
	"github.com/devjunio/gojob/service"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   Gojob
// @contact.url    https://junio.dev/contact
// @contact.email  contact@junio.dev

// @license.name  GLP 3.0
func setRouteConfig(router *gin.Engine, path ...string) {
	basePath := setDefaultRoute(path, "/api/v1")
	docs.SwaggerInfo.BasePath = basePath
	v0 := router.Group(basePath)
	opService := service.InitService()
	baseRoute := "/openings"
	logger.Info("Initializing routes")
	if err := router.SetTrustedProxies([]string{
		os.Getenv("PROXY_URL"),
		"localhost",
		"0.0.0.0",
	}); err != nil {
		return
	}

	{
		v0.GET(baseRoute, opService.List)       // Get
		v0.GET(baseRoute+"/:id", opService.Get) // List
		v0.POST(baseRoute, opService.Create)
		v0.DELETE(baseRoute+"/:id", opService.Delete)
		v0.PATCH(baseRoute+"/:id", opService.Update)
	}

	logger.Info("Routes initialized")
	logger.Infof("Listening and serving HTTP on %s", "http://"+fullpath+basePath)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	logger.Info("Swagger initialized")
	logger.Infof("Swagger HTTP on %s", "http://"+fullpath+"/swagger")
}
