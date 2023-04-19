package router

import "github.com/gin-gonic/gin"

func Initialize() {
	route := gin.Default()
	initializeRoute(route)
	route.Run(":8080")
}

