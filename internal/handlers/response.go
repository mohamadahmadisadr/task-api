package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeError(w http.ResponseWriter, text string, code int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	res := ErrorResponse{
		Error: text,
	}

	json.NewEncoder(w).Encode(res)
}
