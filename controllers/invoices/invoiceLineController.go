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

func CreateLine(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	invoiceId := chi.URLParam(r, "InvoiceId")

	res := database.DB.First(&invoice, invoiceId)

	if res.Error != nil {
		responses.Error(w, res.Error)
		return
	}

	var params requests.InvoiceLineCreateRequest
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		responses.Error(w, err)
		return
	}

	line := models.InvoiceLine{
		InvoiceId: invoice.Id,
		Amount:    params.Amount * 100,
		Name:      params.Name,
		UnitPrice: params.UnitPrice * 100,
	}

	database.DB.Create(&line)

	responses.Json(w, line)
}
