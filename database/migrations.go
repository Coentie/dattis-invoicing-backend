package database

import (
	"dattis/models"
)

func Migrate() error {
	err := DB.AutoMigrate(&models.Invoice{})
	err = DB.AutoMigrate(&models.Customer{})
	err = DB.AutoMigrate(&models.Tax{})

	return err
}
