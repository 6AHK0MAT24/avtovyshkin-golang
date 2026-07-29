package models

import "time"

// BaseEntity contains common fields for all entities
// This provides a consistent structure across all microservices
type BaseEntity struct {
	ID        string    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Timestamps represents creation and update timestamps
type Timestamps struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Status represents the status of an entity
type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusBlocked  Status = "blocked"
)

// IsValidStatus checks if a status is valid
func IsValidStatus(status Status) bool {
	switch status {
	case StatusActive, StatusInactive, StatusBlocked:
		return true
	default:
		return false
	}
}

// EntityWithStatus represents an entity with status
type EntityWithStatus struct {
	BaseEntity
	Status Status `json:"status" db:"status"`
}

// PaginationRequest represents pagination parameters
type PaginationRequest struct {
	Page    int `json:"page" query:"page"`
	PerPage int `json:"per_page" query:"per_page"`
}

// PaginationResponse represents pagination metadata
type PaginationResponse struct {
	Total   int `json:"total"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

// ListResponse represents a paginated list response
type ListResponse[T any] struct {
	Data       []T                `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
