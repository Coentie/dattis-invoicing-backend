package invoices

import (
	"dattis/database"
	"dattis/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Indexes the invoices.
func Index(w http.ResponseWriter, r *http.Request) {
	var invoices []models.Invoice

	database.DB.Find(&invoices)

	data, err := json.Marshal(invoices)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// Update updates the invoice with a status of paid.
func Update(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	id := chi.URLParam(r, "id")
	res := database.DB.First(&invoice, id)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	invoice.Paid = true
	database.DB.Save(&invoice)

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(invoice)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
