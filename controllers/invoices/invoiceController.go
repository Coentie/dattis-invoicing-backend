package invoices

import (
	"dattis/database"
	"dattis/http/requests"
	"dattis/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Index the invoices.
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

// Create endpoint for invoices
func Create(w http.ResponseWriter, r *http.Request) {
	var params requests.InvoiceCreateRequest
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(params)
	w.WriteHeader(http.StatusOK)

	invoice := models.Invoice{
		Name:       params.Name,
		CustomerId: params.Customer,
	}

	database.DB.Create(&invoice)

	data, err := json.Marshal(invoice)

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

	data, err := json.Marshal(invoice)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
