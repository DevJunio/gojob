package router

import "github.com/gin-gonic/gin"

func Initialize() {
	route := gin.Default()
	initializeRoute(route)
	err := route.Run(":8080")
	if err != nil {
		panic("error: initializing route on router.go")
		return
	}
}
