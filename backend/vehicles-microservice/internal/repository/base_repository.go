package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

// BaseEntity represents common fields for all entities
type BaseEntity struct {
	ID        uuid.UUID `db:"fldid"`
	CreatedAt string    `db:"fldcreatedat"`
	UpdatedAt string    `db:"fldupdatedat"`
}

// BaseRepository provides common CRUD operations for all entities
type BaseRepository interface {
	Create(ctx context.Context, entity interface{}) error
	GetByID(ctx context.Context, id uuid.UUID, dest interface{}) error
	GetAll(ctx context.Context, limit, offset int, dest interface{}) (int64, error)
	Update(ctx context.Context, entity interface{}) error
	Delete(ctx context.Context, id uuid.UUID) error
	Search(ctx context.Context, query string, limit, offset int, dest interface{}) (int64, error)
}

// baseRepository implements BaseRepository interface
type baseRepository struct {
	db        *sqlx.DB
	tableName string
}

// NewBaseRepository creates a new base repository instance
func NewBaseRepository(db *sqlx.DB, tableName string) BaseRepository {
	return &baseRepository{
		db:        db,
		tableName: tableName,
	}
}

// Create inserts a new entity into the database
func (r *baseRepository) Create(ctx context.Context, entity interface{}) error {
	query := fmt.Sprintf("INSERT INTO %s DEFAULT VALUES", r.tableName)
	_, err := r.db.NamedExecContext(ctx, query, entity)
	if err != nil {
		return fmt.Errorf("failed to create entity: %w", err)
	}
	return nil
}

// GetByID retrieves an entity by its ID
func (r *baseRepository) GetByID(ctx context.Context, id uuid.UUID, dest interface{}) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE fldid = $1", r.tableName)
	err := r.db.GetContext(ctx, dest, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("entity not found")
		}
		return fmt.Errorf("failed to get entity: %w", err)
	}
	return nil
}

// GetAll retrieves all entities with pagination
func (r *baseRepository) GetAll(ctx context.Context, limit, offset int, dest interface{}) (int64, error) {
	// Get total count
	var total int64
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", r.tableName)
	err := r.db.GetContext(ctx, &total, countQuery)
	if err != nil {
		return 0, fmt.Errorf("failed to count entities: %w", err)
	}

	// Get entities with pagination
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY fldcreatedat DESC LIMIT $1 OFFSET $2", r.tableName)
	err = r.db.SelectContext(ctx, dest, query, limit, offset)
	if err != nil {
		return 0, fmt.Errorf("failed to get entities: %w", err)
	}

	return total, nil
}

// Update updates an existing entity in the database
func (r *baseRepository) Update(ctx context.Context, entity interface{}) error {
	query := fmt.Sprintf("UPDATE %s SET fldupdatedat = NOW() WHERE fldid = :fldid", r.tableName)
	result, err := r.db.NamedExecContext(ctx, query, entity)
	if err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("entity not found")
	}

	return nil
}

// Delete removes an entity from the database
func (r *baseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE fldid = $1", r.tableName)
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete entity: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("entity not found")
	}

	return nil
}

// Search searches entities by query
func (r *baseRepository) Search(ctx context.Context, query string, limit, offset int, dest interface{}) (int64, error) {
	searchQuery := fmt.Sprintf("%%%s%%", query)
	whereClause := fmt.Sprintf("WHERE CAST(fldid AS VARCHAR) ILIKE $1", r.tableName)

	// Get total count
	var total int64
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", r.tableName, whereClause)
	err := r.db.GetContext(ctx, &total, countQuery, searchQuery)
	if err != nil {
		return 0, fmt.Errorf("failed to count search results: %w", err)
	}

	// Get entities with pagination
	selectQuery := fmt.Sprintf("SELECT * FROM %s %s ORDER BY fldcreatedat DESC LIMIT $2 OFFSET $3", r.tableName, whereClause)
	err = r.db.SelectContext(ctx, dest, selectQuery, searchQuery, limit, offset)
	if err != nil {
		return 0, fmt.Errorf("failed to search entities: %w", err)
	}

	return total, nil
}
