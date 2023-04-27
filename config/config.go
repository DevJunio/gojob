package config

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLite
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing SQLite: %v", err)
	}

	return errors.New("fake Error")
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = newLogger()
	return logger
}
