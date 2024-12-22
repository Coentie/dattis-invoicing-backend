package tax

import (
	"dattis/database"
	"dattis/models"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var taxRates []models.Tax

	database.DB.Find(&taxRates)

	data, err := json.Marshal(taxRates)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
