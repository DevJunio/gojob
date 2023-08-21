package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getID receives the id from the request and returns it if it is not empty,
// otherwise it returns an empty string
func getID(ctx *gin.Context) string {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return ""
	}

	return id
}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, data any) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}
