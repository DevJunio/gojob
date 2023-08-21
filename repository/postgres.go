package repository

import (
	"errors"
	"fmt"
	"github.com/devjunio/gojob/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGDatabase struct {
	user string
	pass string
	name string
	host string
	port string
}

func InitDatabase() error {
	InitLogger()
	database := &PGDatabase{
		user: os.Getenv("POSTGRES_USER"),
		pass: os.Getenv("POSTGRES_PASSWORD"),
		name: os.Getenv("POSTGRES_DB"),
		host: os.Getenv("DB_HOST"),
		port: os.Getenv("DB_PORT"),
	}

	if db != nil {
		Logger.Error("error on connecting to database: database connection already exist")
		return errors.New("database connection already exist")
	}

	var err error
	db, err = gorm.Open(postgres.Open(url(database)), &gorm.Config{})
	if err != nil {
		Logger.Debugf("database url: %v", url(database))
		Logger.Errorf("could not open a database connection: %v", err.Error())
		return fmt.Errorf("could not open a database connection: %v", err)
	}
	Logger.Info("database connection successfully established")

	err = db.AutoMigrate(&model.Opening{})
	if err != nil {
		Logger.Errorf("error on auto-migration: %v", err.Error())
		return fmt.Errorf("could not automigrate: %v", err)
	}
	Logger.Info("database automigration successfully executed")

	return nil
}

func url(d *PGDatabase) string {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		d.user, d.pass, d.name, d.host, d.port)

	return connStr
}
