package router

import (
	"github.com/devjunio/gojob/config"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	logger   *config.Logger
	path     string
	port     string
	Fullpath string
)

func SetupRouter() *gin.Engine {
	logger = config.SetLogger("router")
	route := gin.Default()

	path = os.Getenv("PATH")
	if path == "" {
		path = "0.0.0.0"
	}

	port = ":" + os.Getenv("PORT")
	if port == ":" {
		port += "8080"
	}

	Fullpath = path + port

	setRouteConfig(route)

	return route
}

func setDefaultRoute(path []string, route string) string {
	if len(path) > 0 {
		route = path[0]
	}

	return route
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
