package database

import "dattis/models"

func Seed() {
	var count int64

	DB.Model(&models.InvoiceStatus{}).Count(&count)

	if count == 0 {
		DB.Create([]*models.InvoiceStatus{
			{Id: 1, Name: "Draft"},
			{Id: 2, Name: "Send"},
			{Id: 3, Name: "Paid"},
			{Id: 4, Name: "Taxes paid"},
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

	DB.Model(&models.Invoice{}).Count(&count)

	if count == 0 {
		DB.Create([]*models.Invoice{
			{Name: "Test invoice 1", Paid: false, InvoiceStatusId: 1, CustomerId: 1},
			{Name: "Test invoice 2", Paid: false, InvoiceStatusId: 1, CustomerId: 1},
			{Name: "Test invoice 3", Paid: false, InvoiceStatusId: 1, CustomerId: 1},
			{Name: "Test invoice 4", Paid: false, InvoiceStatusId: 1, CustomerId: 1},
		})
	}

	DB.Model(&models.Tax{}).Count(&count)

	if count == 0 {
		DB.Create([]*models.Tax{
			{Name: "0%", Value: 0},
			{Name: "6%", Value: 6},
			{Name: "9%", Value: 9},
			{Name: "21%", Value: 21},
		})
	}
}
