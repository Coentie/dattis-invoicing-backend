package invoices

import (
	"dattis/database"
	"dattis/http/requests"
	"dattis/http/responses"
	"dattis/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Index the invoices.
func Index(w http.ResponseWriter, r *http.Request) {
	var invoices []models.Invoice

	database.DB.Find(&invoices)

	responses.Json(w, invoices)
}

// Show one invoice
func Show(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	invoiceId := chi.URLParam(r, "invoiceId")
	res := database.DB.Preload("Customer").Preload("InvoiceStatus").Preload("InvoiceLines").First(&invoice, invoiceId)

	if res.Error != nil {
		responses.Error(w, res.Error)
		return
	}

	responses.Json(w, invoice)
}

// Create endpoint for invoices
func Create(w http.ResponseWriter, r *http.Request) {
	var params requests.InvoiceCreateRequest
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		responses.Error(w, err)
		return
	}

	invoice := models.Invoice{
		Name:            params.Name,
		CustomerId:      params.Customer,
		InvoiceStatusId: 1, // Default to draft
	}

	res := database.DB.Create(&invoice)

	if res.Error != nil {
		responses.Error(w, res.Error)
		return
	}

	responses.Json(w, invoice)
}

// Update updates the invoice with a status of paid.
func Update(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	id := chi.URLParam(r, "invoiceId")
	res := database.DB.First(&invoice, id)

	if res.Error != nil {
		responses.Error(w, res.Error)
		return
	}

	invoice.Paid = true
	database.DB.Save(&invoice)
	responses.Json(w, invoice)
}
