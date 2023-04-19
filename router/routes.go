package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoute(router *gin.Engine, yay...string) {
	var p = ""
	if len(yay) > 0 { p = yay[0] } else { p = "/api/v0" }

	v0 := router.Group(p)
	{
		path := "/openings"
		v0.GET(path, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H {
				"body": "getando",
			})
		} )

		v0.POST(path, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H {
				"body": "postando",
			})
		} )

		v0.DELETE(path, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H {
				"body": "deletando",
			})
		} )
	}
}
