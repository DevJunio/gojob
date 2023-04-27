package main

import (
	"gojob/config"
	"gojob/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize the config
	err := config.Init()
	if err == nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
