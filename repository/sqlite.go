package repository

import (
	"fmt"
	"github.com/devjunio/gojob/config"
	"github.com/devjunio/gojob/model"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
}

//goland:noinspection GoUnusedExportedFunction
func InitializeSQLite() (*gorm.DB, error) {
	logger := config.SetLogger("sqlite")
	dbDir := "./db"
	dbPath := fmt.Sprint(dbDir, "/main.db")

	// Check if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating a new one...")

		// Create the database file and directory
		err = os.MkdirAll(dbDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		err = file.Close()
		if err != nil {
			return nil, err
		}
	}

	// Create DB and connect
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error while opening sqlite: %v", err)
		return nil, err
	}

	// Migrate last schema
	err = db.AutoMigrate(&model.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
