package models

import "time"

type InvoiceLine struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	InvoiceId int64     `json:"invoiceId"`
	Name      string    `json:"name"`
	Amount    int64     `json:"amount"`
	UnitPrice int64     `json:"unitPrice"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Invoice Invoice `json:"Invoice"`
}
