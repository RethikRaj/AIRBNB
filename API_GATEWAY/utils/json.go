package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func writeJsonResponse(w http.ResponseWriter, response *response) error {
	w.Header().Set("Content-Type", "application/json") // Set the header content type to JSON

	w.WriteHeader(response.Status) // set the http status code
	// Encode and write to response
	return json.NewEncoder(w).Encode(response)
}

func WriteSuccessJsonResponse(w http.ResponseWriter, statusCode int, message string, data any) error {
	successResp := response{
		Message: message,
		Data:    data,
		Status:  statusCode,
		Success: true,
	}
	return writeJsonResponse(w, &successResp)
}

func WriteErrorJsonResponse(w http.ResponseWriter, statusCode int, message string, err error) error {
	errorResp := response{
		Message: message,
		Status:  statusCode,
		Success: false,
		Error:   err.Error(),
	}

	return writeJsonResponse(w, &errorResp)
}
