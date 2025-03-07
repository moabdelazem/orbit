package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSON writes a JSON response with the specified status code
func WriteJSON(w http.ResponseWriter, response interface{}, statusCode int, logger *log.Logger) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Marshal the response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		if logger != nil {
			logger.Printf("Error marshalling JSON response: %v", err)
		}
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	// Write the response
	if _, err := w.Write(jsonResponse); err != nil {
		if logger != nil {
			logger.Printf("Error writing JSON response: %v", err)
		}
		logger.Fatal("Failed to write response")
		return
	}
}

// ReadJSON reads a JSON request body into the provided struct
func ReadJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
