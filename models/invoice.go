package models

import (
	"time"
)

type Invoice struct {
	Id              int64     `json:"id" gorm:"primaryKey"`
	InvoiceStatusId int64     `json:"invoiceStatusId"`
	CustomerId      int64     `json:"customerId"`
	Name            string    `json:"name"`
	Paid            bool      `json:"paid"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	InvoiceStatus InvoiceStatus `json:"invoiceStatus"`
	Customer      Customer      `json:"customer"`
	
	InvoiceLines []InvoiceLine `json:"invoiceLines"`
}
