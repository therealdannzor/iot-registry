package api

import (
	"encoding/json"
	"net/http"
)

func response(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func errorResponse(w http.ResponseWriter, code int, err error) {
	customError := map[string]string{"error": err.Error()}
	body, err := json.Marshal(customError)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}
	response(w, code, []byte(body))
}
