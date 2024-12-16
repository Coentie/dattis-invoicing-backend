package database

import "dattis/models"

func Seed() {
	var count int64

	DB.Model(&models.Invoice{}).Count(&count)

	if count == 0 {
		DB.Create([]*models.Invoice{
			{Name: "Test invoice 1", Paid: false},
			{Name: "Test invoice 2", Paid: false},
			{Name: "Test invoice 3", Paid: false},
			{Name: "Test invoice 4", Paid: false},
		})
	}

	DB.Model(&models.Customer{}).Count(&count)

	if count == 0 {
		DB.Create([]*models.Customer{
			{Name: "Hanze-ICT", Email: "frank.bokkers@euromaster.com"},
			{Name: "Test klant", Email: "test@email.nl"},
			{Name: "Test klant 2", Email: "test2@email.nl"},
			{Name: "Test klant 3", Email: "test3@email.nl"},
		})
	}

}
