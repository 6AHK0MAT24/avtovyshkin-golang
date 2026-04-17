package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// BaseHandler provides common HTTP handlers for all entities
type BaseHandler struct {
	service interface{}
}

// NewBaseHandler creates a new base handler instance
func NewBaseHandler(service interface{}) *BaseHandler {
	return &BaseHandler{
		service: service,
	}
}

// ListResponse represents a paginated list response
type ListResponse struct {
	Data     []interface{} `json:"data"`
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PerPage  int           `json:"perPage"`
}

// CreateRequest represents a generic create request
type CreateRequest interface{}

// UpdateRequest represents a generic update request
type UpdateRequest interface{}

// Create handles POST /api/{entity}
func (h *BaseHandler) Create(w http.ResponseWriter, r *http.Request, entityName string, createFunc func(context.Context, CreateRequest) (interface{}, error)) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	entity, err := createFunc(context.Background(), req)
	if err != nil {
		log.Printf("Error creating %s: %v", entityName, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusCreated, entity)
}

// GetByID handles GET /api/{entity}/{id}
func (h *BaseHandler) GetByID(w http.ResponseWriter, r *http.Request, entityName string, getFunc func(context.Context, uuid.UUID) (interface{}, error)) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/"+entityName+"/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid "+entityName+" ID", http.StatusBadRequest)
		return
	}

	entity, err := getFunc(context.Background(), id)
	if err != nil {
		log.Printf("Error getting %s: %v", entityName, err)
		http.Error(w, entityName+" not found", http.StatusNotFound)
		return
	}

	respondJSON(w, http.StatusOK, entity)
}

// GetAll handles GET /api/{entity}
func (h *BaseHandler) GetAll(w http.ResponseWriter, r *http.Request, entityName string, getAllFunc func(context.Context, int, int) ([]interface{}, int64, error)) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	entities, total, err := getAllFunc(context.Background(), page, pageSize)
	if err != nil {
		log.Printf("Error getting %s: %v", entityName, err)
		http.Error(w, "Failed to get "+entityName, http.StatusInternalServerError)
		return
	}

	response := ListResponse{
		Data:    entities,
		Total:   total,
		Page:    page,
		PerPage: pageSize,
	}

	respondJSON(w, http.StatusOK, response)
}

// Update handles PUT /api/{entity}/{id}
func (h *BaseHandler) Update(w http.ResponseWriter, r *http.Request, entityName string, updateFunc func(context.Context, uuid.UUID, UpdateRequest) (interface{}, error)) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/"+entityName+"/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid "+entityName+" ID", http.StatusBadRequest)
		return
	}

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	entity, err := updateFunc(context.Background(), id, req)
	if err != nil {
		log.Printf("Error updating %s: %v", entityName, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusOK, entity)
}

// Delete handles DELETE /api/{entity}/{id}
func (h *BaseHandler) Delete(w http.ResponseWriter, r *http.Request, entityName string, deleteFunc func(context.Context, uuid.UUID) error) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/"+entityName+"/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid "+entityName+" ID", http.StatusBadRequest)
		return
	}

	if err := deleteFunc(context.Background(), id); err != nil {
		log.Printf("Error deleting %s: %v", entityName, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": entityName + " deleted successfully"})
}

// Search handles GET /api/{entity}/search
func (h *BaseHandler) Search(w http.ResponseWriter, r *http.Request, entityName string, searchFunc func(context.Context, string, int, int) ([]interface{}, int64, error)) {
	query := r.URL.Query().Get("q")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	entities, total, err := searchFunc(context.Background(), query, page, pageSize)
	if err != nil {
		log.Printf("Error searching %s: %v", entityName, err)
		http.Error(w, "Failed to search "+entityName, http.StatusInternalServerError)
		return
	}

	response := ListResponse{
		Data:    entities,
		Total:   total,
		Page:    page,
		PerPage: pageSize,
	}

	respondJSON(w, http.StatusOK, response)
}

// respondJSON sends a JSON response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
