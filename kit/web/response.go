package web

import (
	"encoding/json"
	"net/http"
)

// Respond converts a Go value to JSON and sends it to the client.
func Respond(w http.ResponseWriter, data any, statusCode int) {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
	}

	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		println(err)
	}

	jsonData = append(jsonData, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		println(err)
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}
