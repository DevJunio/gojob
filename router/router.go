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

func setDefaultRoute(path []string, route string) string {
	var p = ""

	if len(path) > 0 {
		p = path[0]
	} else {
		p = route
	}

	return p
}
