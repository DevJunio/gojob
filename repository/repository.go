package repository

import (
	"github.com/devjunio/gojob/config"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	db     *gorm.DB
)

func InitLogger() {
	Logger = config.SetLogger("repository")
}
