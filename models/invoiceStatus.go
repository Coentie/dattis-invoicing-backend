package models

import "time"

type InvoiceStatus struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Invoices []Invoice `json:"invoices"`
}
