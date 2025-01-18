package responses

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)

	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)
	if err != nil {
		println(err.Error())
		return
	}
}

func Success(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Error(w http.ResponseWriter, err error) {
	println(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}
