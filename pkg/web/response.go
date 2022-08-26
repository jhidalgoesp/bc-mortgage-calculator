package web

import (
	"encoding/json"
	"log"
	"net/http"
)

// Respond converts a Go value to JSON and sends it to the client.
func Respond(w http.ResponseWriter, data any, statusCode int) {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
	}

	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("unable to encode response: %w", err)
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		log.Println("unable to write response: %w", err)
	}
}
