package database

import (
	"dattis/models"
)

func Migrate() error {
	err := DB.AutoMigrate(&models.Invoice{})

	return err
}
