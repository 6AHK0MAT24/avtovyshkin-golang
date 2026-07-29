package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseWriter provides helper functions for writing HTTP responses
type ResponseWriter struct {
	w http.ResponseWriter
}

// NewResponseWriter creates a new ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w: w}
}

// JSON writes a JSON response with status code
func (rw *ResponseWriter) JSON(status int, data interface{}) error {
	rw.w.Header().Set("Content-Type", "application/json")
	rw.w.WriteHeader(status)
	return json.NewEncoder(rw.w).Encode(data)
}

// Success writes a success response
func (rw *ResponseWriter) Success(data interface{}) error {
	return rw.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

// Error writes an error response
func (rw *ResponseWriter) Error(status int, message string) error {
	return rw.JSON(status, map[string]interface{}{
		"success": false,
		"error":   message,
	})
}

// ValidationError writes a validation error response
func (rw *ResponseWriter) ValidationError(errors map[string]string) error {
	return rw.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": false,
		"error":   "Validation failed",
		"errors":  errors,
	})
}

// Created writes a 201 Created response
func (rw *ResponseWriter) Created(data interface{}) error {
	return rw.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

// NoContent writes a 204 No Content response
func (rw *ResponseWriter) NoContent() error {
	rw.w.WriteHeader(http.StatusNoContent)
	return nil
}

// NotFound writes a 404 Not Found response
func (rw *ResponseWriter) NotFound(message string) error {
	if message == "" {
		message = "Resource not found"
	}
	return rw.Error(http.StatusNotFound, message)
}

// BadRequest writes a 400 Bad Request response
func (rw *ResponseWriter) BadRequest(message string) error {
	if message == "" {
		message = "Bad request"
	}
	return rw.Error(http.StatusBadRequest, message)
}

// Unauthorized writes a 401 Unauthorized response
func (rw *ResponseWriter) Unauthorized(message string) error {
	if message == "" {
		message = "Unauthorized"
	}
	return rw.Error(http.StatusUnauthorized, message)
}

// Forbidden writes a 403 Forbidden response
func (rw *ResponseWriter) Forbidden(message string) error {
	if message == "" {
		message = "Forbidden"
	}
	return rw.Error(http.StatusForbidden, message)
}

// InternalServerError writes a 500 Internal Server Error response
func (rw *ResponseWriter) InternalServerError(message string) error {
	if message == "" {
		message = "Internal server error"
	}
	return rw.Error(http.StatusInternalServerError, message)
}
