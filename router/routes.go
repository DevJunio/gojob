package router

import (
	"fmt"
	"gojob/handler/opening"

	"github.com/gin-gonic/gin"
)

func initializeRoute(router *gin.Engine, yay ...string) {
	var p = ""
	if len(yay) > 0 {
		p = yay[0]
	} else {
		p = "/api/v0"
	}

	v0 := router.Group(p)
	{
		path := "/openings"
		v0.GET(path, handler.ShowOpeningHandler)
		v0.POST(path, handler.CreateOpeningHandler)
		v0.DELETE(path, handler.DeleteOpeningHandler)
		v0.PATCH(path, handler.UpdateOpeningHandler)
		v0.GET(fmt.Sprintf("%s%s", path, "/list"), handler.ListOpeningHandler)
	}
}
