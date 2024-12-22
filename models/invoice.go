package models

import (
	"time"
)

type Invoice struct {
	Id              int64     `json:"id" gorm:"primaryKey"`
	InvoiceStatusId int64     `json:"invoice_status_id"`
	CustomerId      int64     `json:"customer_id"`
	Name            string    `json:"name"`
	Paid            bool      `json:"paid"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	InvoiceStatus InvoiceStatus `json:"invoice_status"`
	Customer      Customer      `json:"customer"`
}
