package response

import (
	"encoding/json"
	"net/http"

	"github.com/Mahathirrr/task-management-backend/internal/model"
)

// JSON mengirim response JSON
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Error mengirim error response
func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, model.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	})
}

// ValidationError mengirim validation error response
func ValidationError(w http.ResponseWriter, errors []model.ValidationError) {
	JSON(w, http.StatusBadRequest, model.ValidationErrorResponse{
		Error:   "Validation failed",
		Details: errors,
	})
}

// Success mengirim success response
func Success(w http.ResponseWriter, message string) {
	JSON(w, http.StatusOK, model.SuccessResponse{
		Message: message,
	})
}

// Created mengirim created response
func Created(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusCreated, data)
}