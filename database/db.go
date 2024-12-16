package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func New() error {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
			os.Getenv("APP_DB_USERNAME"),
			os.Getenv("APP_DB_PASSWORD"),
			os.Getenv("APP_DB_HOST"),
			os.Getenv("APP_DB_PORT"),
			os.Getenv("APP_DB_NAME"))), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db
	return nil
}
