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

	data, err := json.Marshal(invoices)

	if err != nil {
		responses.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Show(w http.ResponseWriter, r *http.Request) {

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

	data, err := json.Marshal(invoice)

	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Success(w, data)
}

// Update updates the invoice with a status of paid.
func Update(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	id := chi.URLParam(r, "id")
	res := database.DB.First(&invoice, id)

	if res.Error != nil {
		responses.Error(w, res.Error)
		return
	}

	invoice.Paid = true
	database.DB.Save(&invoice)

	data, err := json.Marshal(invoice)

	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Success(w, data)
}
