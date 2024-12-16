package customers

import (
	"dattis/database"
	"dattis/models"
	"encoding/json"
	"net/http"
)

// Indexes the invoices.
func Index(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer

	database.DB.Find(&customers)

	data, err := json.Marshal(customers)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
