package router

import (
	"fmt"
	"gojob/handler"
	opening "gojob/handler/opening"

	"github.com/gin-gonic/gin"
)

func initializeRoute(router *gin.Engine, path ...string) {
	handler.InitializeHandler()
	basePath := setDefaultRoute(path, "/api/v0")
	v0 := router.Group(basePath)
	{
		path := "/openings"
		v0.GET(path, opening.ShowOpeningHandler)
		v0.POST(path, opening.CreateOpeningHandler)
		v0.DELETE(path, opening.DeleteOpeningHandler)
		v0.PATCH(path, opening.UpdateOpeningHandler)
		v0.GET(fmt.Sprintf("%s%s", path, "/list"), opening.ListOpeningHandler)
	}
}
