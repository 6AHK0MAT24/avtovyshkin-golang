package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// BaseService provides common business logic for all entities
type BaseService interface {
	Create(ctx context.Context, req interface{}) (interface{}, error)
	GetByID(ctx context.Context, id uuid.UUID) (interface{}, error)
	GetAll(ctx context.Context, page, pageSize int) ([]interface{}, int64, error)
	Update(ctx context.Context, id uuid.UUID, req interface{}) (interface{}, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Search(ctx context.Context, query string, page, pageSize int) ([]interface{}, int64, error)
}

// baseService implements BaseService interface
type baseService struct {
	repo interface{}
}

// NewBaseService creates a new base service instance
func NewBaseService(repo interface{}) BaseService {
	return &baseService{
		repo: repo,
	}
}

// Create creates a new entity with validation
func (s *baseService) Create(ctx context.Context, req interface{}) (interface{}, error) {
	// This is a base implementation - should be overridden by specific services
	return nil, fmt.Errorf("Create method must be implemented by specific service")
}

// GetByID retrieves an entity by ID
func (s *baseService) GetByID(ctx context.Context, id uuid.UUID) (interface{}, error) {
	// This is a base implementation - should be overridden by specific services
	return nil, fmt.Errorf("GetByID method must be implemented by specific service")
}

// GetAll retrieves all entities with pagination
func (s *baseService) GetAll(ctx context.Context, page, pageSize int) ([]interface{}, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	_ = (page - 1) * pageSize // Calculate offset for future use

	// This is a base implementation - should be overridden by specific services
	return nil, 0, fmt.Errorf("GetAll method must be implemented by specific service")
}
// Update updates an existing entity with validation
func (s *baseService) Update(ctx context.Context, id uuid.UUID, req interface{}) (interface{}, error) {
	// This is a base implementation - should be overridden by specific services
	return nil, fmt.Errorf("Update method must be implemented by specific service")
}

// Delete deletes an entity by ID
func (s *baseService) Delete(ctx context.Context, id uuid.UUID) error {
	// This is a base implementation - should be overridden by specific services
	return fmt.Errorf("Delete method must be implemented by specific service")
}

// Search searches entities by query
func (s *baseService) Search(ctx context.Context, query string, page, pageSize int) ([]interface{}, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	_ = (page - 1) * pageSize // Calculate offset for future use

	// This is a base implementation - should be overridden by specific services
	return nil, 0, fmt.Errorf("Search method must be implemented by specific service")
}
