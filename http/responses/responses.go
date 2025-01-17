package responses

import "net/http"

func Success(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Error(w http.ResponseWriter, err error) {
	println(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}
