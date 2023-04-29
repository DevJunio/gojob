package router

import (
	"gojob/handler"
	docs "gojob/docs"
	opening "gojob/handler/opening"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoute(router *gin.Engine, path ...string) {
	handler.InitializeHandler()
	basePath := setDefaultRoute(path, "/api/v0")
	docs.SwaggerInfo.BasePath = basePath
	v0 := router.Group(basePath)
	{
		path := "/openings"
		v0.GET(path, opening.ListOpeningHandler)
		v0.GET("/openings/:id", opening.ShowOpeningHandler)
		v0.POST(path, opening.CreateOpeningHandler)
		v0.DELETE(path, opening.DeleteOpeningHandler)
		v0.PATCH(path, opening.UpdateOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
