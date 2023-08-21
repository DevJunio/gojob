package router

import (
	"github.com/devjunio/gojob/config"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	logger   *config.Logger
	path     string
	port     string
	fullpath string
)

func Initialize() error {
	logger = config.SetLogger("router")
	route := gin.Default()

	path = os.Getenv("APP_PATH")
	if path == "" {
		path = "0.0.0.0"
	}

	port = ":" + os.Getenv("APP_PORT")
	if port == ":" {
		port += "8080"
	}

	fullpath = path + port

	setRouteConfig(route)
	err := route.Run(fullpath)
	if err != nil {
		logger.Error("error: initializing route on router.go")
		return err
	}

	return nil
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
