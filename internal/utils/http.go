package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes a JSON response with the specified status code
func WriteJSON(w http.ResponseWriter, response interface{}, statusCode int) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Marshal the response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	// Write the response
	if _, err := w.Write(jsonResponse); err != nil {
		// Can't write to the response at this point, but log would be useful
		// In a real app, we'd have access to a logger here
		return
	}
}

// ReadJSON reads a JSON request body into the provided struct
func ReadJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
