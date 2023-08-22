package main

import (
	"github.com/devjunio/gojob/config"
	"github.com/devjunio/gojob/repository"
	"github.com/devjunio/gojob/router"
	"github.com/joho/godotenv"
)

var (
	log *config.Logger
)

func main() {
	log = config.SetLogger("main")

	if err := godotenv.Load(".env"); err != nil {
		log.Warn("No .env found")
	}

	if err := repository.InitDatabase(); err != nil {
		log.Errorf("database initialization error: %v", err)
		return
	}

	route := router.SetupRouter()

	err := route.Run(router.Fullpath)
	if err != nil {
		log.Errorf("router initialization error: %v", err)
		return
	}

	repository.InitLogger()
}
