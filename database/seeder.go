package database

import "dattis/models"

func Seed() {
	var invoiceCount int64
	DB.Model(&models.Invoice{}).Count(&invoiceCount)

	if invoiceCount == 0 {
		DB.Create([]*models.Invoice{
			{Name: "Test invoice 1", Paid: false},
			{Name: "Test invoice 2", Paid: false},
			{Name: "Test invoice 3", Paid: false},
			{Name: "Test invoice 4", Paid: false},
		})
	}
}
