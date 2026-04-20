package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
	"vehicles-service/internal/models"
)

// VehicleRepository defines the interface for vehicle data operations
type VehicleRepository interface {
	Create(ctx context.Context, vehicle *models.Vehicle) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Vehicle, error)
	GetAll(ctx context.Context, limit, offset int) ([]*models.Vehicle, int64, error)
	Update(ctx context.Context, vehicle *models.Vehicle) error
	Delete(ctx context.Context, id uuid.UUID) error
	Search(ctx context.Context, query string, limit, offset int) ([]*models.Vehicle, int64, error)
	GetByGarageNumber(ctx context.Context, garageNumber string) (*models.Vehicle, error)
	GetByVIN(ctx context.Context, vin string) (*models.Vehicle, error)
	Close() error
}

// vehicleRepository implements VehicleRepository interface
type vehicleRepository struct {
	db *sqlx.DB
}

// NewVehicleRepository creates a new vehicle repository instance
func NewVehicleRepository(db *sqlx.DB) VehicleRepository {
	return &vehicleRepository{
		db: db,
	}
}

// Create inserts a new vehicle into the database
func (r *vehicleRepository) Create(ctx context.Context, vehicle *models.Vehicle) error {
	query := `
		INSERT INTO vehicles (
			fldid, fldgaragenumber, fldvin, fldheight, fldtype, fldpower,
			fldprice5, fldprice22, flddescription, fldbrand, fldmachine,
			fldlength, fldwidth, fldheightts, fldwidthwithsupports, fldmass,
			fldcradlewidthfolded, fldcradlewidthextended, fldcradlelengthfolded, fldcradlelengthextended,
			fldimgarray, fldmainimageindex, fldspecial, fldrostechreg, fldstatus,
			fldcreatedat, fldupdatedat
		) VALUES (
			:fldid, :fldgaragenumber, :fldvin, :fldheight, :fldtype, :fldpower,
			:fldprice5, :fldprice22, :flddescription, :fldbrand, :fldmachine,
			:fldlength, :fldwidth, :fldheightts, :fldwidthwithsupports, :fldmass,
			:fldcradlewidthfolded, :fldcradlewidthextended, :fldcradlelengthfolded, :fldcradlelengthextended,
			:fldimgarray, :fldmainimageindex, :fldspecial, :fldrostechreg, :fldstatus,
			:fldcreatedat, :fldupdatedat
		)
	`

	_, err := r.db.NamedExecContext(ctx, query, vehicle)
	if err != nil {
		return fmt.Errorf("failed to create vehicle: %w", err)
	}

	return nil
}
// GetByID retrieves a vehicle by its ID
func (r *vehicleRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Vehicle, error) {
	query := `
		SELECT
			fldid, fldgaragenumber, fldvin, fldheight, fldtype, fldpower,
			fldprice5, fldprice22, flddescription, fldbrand, fldmachine,
			fldlength, fldwidth, fldheightts, fldwidthwithsupports, fldmass,
			fldcradlewidthfolded, fldcradlewidthextended, fldcradlelengthfolded, fldcradlelengthextended,
			fldimgarray, fldmainimageindex, fldspecial, fldrostechreg, fldstatus,
			fldcreatedat, fldupdatedat
		FROM vehicles
		WHERE fldid = ?
	`

	var vehicle models.Vehicle
	err := r.db.GetContext(ctx, &vehicle, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("vehicle not found")
		}
		return nil, fmt.Errorf("failed to get vehicle: %w", err)
	}

	return &vehicle, nil
}
// GetAll retrieves all vehicles with pagination
func (r *vehicleRepository) GetAll(ctx context.Context, limit, offset int) ([]*models.Vehicle, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM vehicles`
	err := r.db.GetContext(ctx, &total, countQuery)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count vehicles: %w", err)
	}

	// Get vehicles with pagination
	query := `
		SELECT
			fldid, fldgaragenumber, fldvin, fldheight, fldtype, fldpower,
			fldprice5, fldprice22, flddescription, fldbrand, fldmachine,
			fldlength, fldwidth, fldheightts, fldwidthwithsupports, fldmass,
			fldcradlewidthfolded, fldcradlewidthextended, fldcradlelengthfolded, fldcradlelengthextended,
			fldimgarray, fldmainimageindex, fldspecial, fldrostechreg, fldstatus,
			fldcreatedat, fldupdatedat
		FROM vehicles
		ORDER BY fldcreatedat DESC
		LIMIT ? OFFSET ?
	`
	var vehicles []*models.Vehicle
	err = r.db.SelectContext(ctx, &vehicles, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get vehicles: %w", err)
	}

	return vehicles, total, nil
}

// Update updates an existing vehicle in the database
func (r *vehicleRepository) Update(ctx context.Context, vehicle *models.Vehicle) error {
	query := `
		UPDATE vehicles SET
			fldgaragenumber = :fldgaragenumber,
			fldvin = :fldvin,
			fldheight = :fldheight,
			fldtype = :fldtype,
			fldpower = :fldpower,
			fldprice5 = :fldprice5,
			fldprice22 = :fldprice22,
			flddescription = :flddescription,
			fldbrand = :fldbrand,
			fldmachine = :fldmachine,
			fldlength = :fldlength,
			fldwidth = :fldwidth,
			fldheightts = :fldheightts,
			fldwidthwithsupports = :fldwidthwithsupports,
			fldmass = :fldmass,
			fldcradlewidthfolded = :fldcradlewidthfolded,
			fldcradlewidthextended = :fldcradlewidthextended,
			fldcradlelengthfolded = :fldcradlelengthfolded,
			fldcradlelengthextended = :fldcradlelengthextended,
			fldimgarray = :fldimgarray,
			fldmainimageindex = :fldmainimageindex,
			fldspecial = :fldspecial,
			fldrostechreg = :fldrostechreg,
			fldstatus = :fldstatus,
			fldupdatedat = :fldupdatedat
		WHERE fldid = :fldid
	`

	result, err := r.db.NamedExecContext(ctx, query, vehicle)
	if err != nil {
		return fmt.Errorf("failed to update vehicle: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("vehicle not found")
	}

	return nil
}
// Delete removes a vehicle from the database
func (r *vehicleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM vehicles WHERE fldid = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete vehicle: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("vehicle not found")
	}

	return nil
}
// Search searches vehicles by height and garage number
func (r *vehicleRepository) Search(ctx context.Context, query string, limit, offset int) ([]*models.Vehicle, int64, error) {
	// Parse query - can be height or garage number
	searchQuery := strings.TrimSpace(query)
	if searchQuery == "" {
		return r.GetAll(ctx, limit, offset)
	}

	// Build search query
	whereClause := `WHERE (fldgaragenumber LIKE ? OR CAST(fldheight AS CHAR) LIKE ?)`
	searchParam := "%" + searchQuery + "%"

	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM vehicles ` + whereClause
	err := r.db.GetContext(ctx, &total, countQuery, searchParam, searchParam)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count search results: %w", err)
	}

	// Get vehicles with pagination
	selectQuery := `
		SELECT
			fldid, fldgaragenumber, fldvin, fldheight, fldtype, fldpower,
			fldprice5, fldprice22, flddescription, fldbrand, fldmachine,
			fldlength, fldwidth, fldheightts, fldwidthwithsupports, fldmass,
			fldcradlewidthfolded, fldcradlewidthextended, fldcradlelengthfolded, fldcradlelengthextended,
			fldimgarray, fldmainimageindex, fldspecial, fldrostechreg, fldstatus,
			fldcreatedat, fldupdatedat
		FROM vehicles
		` + whereClause + `
		ORDER BY fldcreatedat DESC
		LIMIT ? OFFSET ?
	`

	var vehicles []*models.Vehicle
	err = r.db.SelectContext(ctx, &vehicles, selectQuery, searchParam, searchParam, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to search vehicles: %w", err)
	}
	return vehicles, total, nil
}
// GetByGarageNumber retrieves a vehicle by its garage number
func (r *vehicleRepository) GetByGarageNumber(ctx context.Context, garageNumber string) (*models.Vehicle, error) {
	query := `
		SELECT
			fldid, fldgaragenumber, fldvin, fldheight, fldtype, fldpower,
			fldprice5, fldprice22, flddescription, fldbrand, fldmachine,
			fldlength, fldwidth, fldheightts, fldwidthwithsupports, fldmass,
			fldcradlewidthfolded, fldcradlewidthextended, fldcradlelengthfolded, fldcradlelengthextended,
			fldimgarray, fldmainimageindex, fldspecial, fldrostechreg, fldstatus,
			fldcreatedat, fldupdatedat
		FROM vehicles
		WHERE fldgaragenumber = ?
	`

	var vehicle models.Vehicle
	err := r.db.GetContext(ctx, &vehicle, query, garageNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get vehicle by garage number: %w", err)
	}

	return &vehicle, nil
}

// GetByVIN retrieves a vehicle by its VIN
func (r *vehicleRepository) GetByVIN(ctx context.Context, vin string) (*models.Vehicle, error) {
	query := `		SELECT
			fldid, fldgaragenumber, fldvin, fldheight, fldtype, fldpower,
			fldprice5, fldprice22, flddescription, fldbrand, fldmachine,
			fldlength, fldwidth, fldheightts, fldwidthwithsupports, fldmass,
			fldcradlewidthfolded, fldcradlewidthextended, fldcradlelengthfolded, fldcradlelengthextended,
			fldimgarray, fldmainimageindex, fldspecial, fldrostechreg, fldstatus,
			fldcreatedat, fldupdatedat
		FROM vehicles
		WHERE fldvin = ?
	`

	var vehicle models.Vehicle
	err := r.db.GetContext(ctx, &vehicle, query, vin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get vehicle by VIN: %w", err)
	}

	return &vehicle, nil
}

// Close closes the database connection
func (r *vehicleRepository) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}
