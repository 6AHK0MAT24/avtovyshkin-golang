package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"drivers-service/internal/config"
	"drivers-service/internal/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// DriverRepository defines the interface for driver data operations
type DriverRepository interface {
	Create(ctx context.Context, driver *models.Driver) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Driver, error)
	GetAll(ctx context.Context, limit, offset int) ([]models.Driver, int64, error)
	Update(ctx context.Context, driver *models.Driver) error
	Delete(ctx context.Context, id uuid.UUID) error
	Search(ctx context.Context, query string, status *models.DriverStatus, limit, offset int) ([]models.Driver, int64, error)
	GetByPhone(ctx context.Context, phone string) (*models.Driver, error)
	GetByLicenseNumber(ctx context.Context, licenseNumber string) (*models.Driver, error)
	GetByEmail(ctx context.Context, email string) (*models.Driver, error)
	Close() error
}

// driverRepository implements DriverRepository interface
type driverRepository struct {
	db *sql.DB
}

// NewDriverRepository creates a new driver repository
func NewDriverRepository(cfg *config.DatabaseConfig) (DriverRepository, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &driverRepository{db: db}, nil
}

// Create creates a new driver in the database
func (r *driverRepository) Create(ctx context.Context, driver *models.Driver) error {
	query := `
		INSERT INTO drivers (
			fldId, fldFirstName, fldLastName, fldMiddleName, fldPhone, fldEmail,
			fldBirthDate, fldPhoto, fldDriverLicenseNumber, fldDriverLicenseIssueDate,
			fldDriverLicenseExpiryDate, fldDriverLicensePhoto, fldDriverLicenseScan,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate,
			fldPassportPhoto, fldPassportScan, fldAddress, fldExperienceYears,
			fldStatus, fldCreatedAt, fldUpdatedAt
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
	`

	_, err := r.db.ExecContext(ctx, query,
		driver.ID, driver.FirstName, driver.LastName, driver.MiddleName, driver.Phone, driver.Email,
		driver.BirthDate, driver.Photo, driver.DriverLicenseNumber, driver.DriverLicenseIssueDate,
		driver.DriverLicenseExpiryDate, driver.DriverLicensePhoto, driver.DriverLicenseScan,
		driver.PassportSeries, driver.PassportNumber, driver.PassportIssueDate,
		driver.PassportPhoto, driver.PassportScan, driver.Address, driver.ExperienceYears,
		driver.Status, driver.CreatedAt, driver.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create driver: %w", err)
	}

	return nil
}
// GetByID retrieves a driver by ID
func (r *driverRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Driver, error) {
	query := `
		SELECT
			fldId, fldFirstName, fldLastName, fldMiddleName, fldPhone, fldEmail,
			fldBirthDate, fldPhoto, fldDriverLicenseNumber, fldDriverLicenseIssueDate,
			fldDriverLicenseExpiryDate, fldDriverLicensePhoto, fldDriverLicenseScan,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate,
			fldPassportPhoto, fldPassportScan, fldAddress, fldExperienceYears,
			fldStatus, fldCreatedAt, fldUpdatedAt
		FROM drivers
		WHERE fldId = $1
	`

	driver := &models.Driver{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&driver.ID, &driver.FirstName, &driver.LastName, &driver.MiddleName, &driver.Phone, &driver.Email,
		&driver.BirthDate, &driver.Photo, &driver.DriverLicenseNumber, &driver.DriverLicenseIssueDate,
		&driver.DriverLicenseExpiryDate, &driver.DriverLicensePhoto, &driver.DriverLicenseScan,
		&driver.PassportSeries, &driver.PassportNumber, &driver.PassportIssueDate,
		&driver.PassportPhoto, &driver.PassportScan, &driver.Address, &driver.ExperienceYears,
		&driver.Status, &driver.CreatedAt, &driver.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("driver not found")
		}
		return nil, fmt.Errorf("failed to get driver: %w", err)
	}

	return driver, nil
}
// GetAll retrieves all drivers with pagination
func (r *driverRepository) GetAll(ctx context.Context, limit, offset int) ([]models.Driver, int64, error) {
	// Get total count
	var total int64
	countQuery := "SELECT COUNT(*) FROM drivers"
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count drivers: %w", err)
	}

	// Get drivers
	query := `
		SELECT
			fldId, fldFirstName, fldLastName, fldMiddleName, fldPhone, fldEmail,
			fldBirthDate, fldPhoto, fldDriverLicenseNumber, fldDriverLicenseIssueDate,
			fldDriverLicenseExpiryDate, fldDriverLicensePhoto, fldDriverLicenseScan,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate,
			fldPassportPhoto, fldPassportScan, fldAddress, fldExperienceYears,
			fldStatus, fldCreatedAt, fldUpdatedAt
		FROM drivers
		ORDER BY fldCreatedAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query drivers: %w", err)
	}
	defer rows.Close()

	drivers := []models.Driver{}
	for rows.Next() {
		driver := models.Driver{}
		err := rows.Scan(
			&driver.ID, &driver.FirstName, &driver.LastName, &driver.MiddleName, &driver.Phone, &driver.Email,
			&driver.BirthDate, &driver.Photo, &driver.DriverLicenseNumber, &driver.DriverLicenseIssueDate,
			&driver.DriverLicenseExpiryDate, &driver.DriverLicensePhoto, &driver.DriverLicenseScan,
			&driver.PassportSeries, &driver.PassportNumber, &driver.PassportIssueDate,
			&driver.PassportPhoto, &driver.PassportScan, &driver.Address, &driver.ExperienceYears,
			&driver.Status, &driver.CreatedAt, &driver.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan driver: %w", err)
		}
		drivers = append(drivers, driver)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating drivers: %w", err)
	}

	return drivers, total, nil
}
// Update updates an existing driver
func (r *driverRepository) Update(ctx context.Context, driver *models.Driver) error {
	query := `
		UPDATE drivers SET
			fldFirstName = $1,
			fldLastName = $2,
			fldMiddleName = $3,
			fldPhone = $4,
			fldEmail = $5,
			fldBirthDate = $6,
			fldPhoto = $7,
			fldDriverLicenseNumber = $8,
			fldDriverLicenseIssueDate = $9,
			fldDriverLicenseExpiryDate = $10,
			fldDriverLicensePhoto = $11,
			fldDriverLicenseScan = $12,
			fldPassportSeries = $13,
			fldPassportNumber = $14,
			fldPassportIssueDate = $15,
			fldPassportPhoto = $16,
			fldPassportScan = $17,
			fldAddress = $18,
			fldExperienceYears = $19,
			fldStatus = $20,
			fldUpdatedAt = $21
		WHERE fldId = $22
	`

	result, err := r.db.ExecContext(ctx, query,
		driver.FirstName, driver.LastName, driver.MiddleName, driver.Phone, driver.Email,
		driver.BirthDate, driver.Photo, driver.DriverLicenseNumber, driver.DriverLicenseIssueDate,
		driver.DriverLicenseExpiryDate, driver.DriverLicensePhoto, driver.DriverLicenseScan,
		driver.PassportSeries, driver.PassportNumber, driver.PassportIssueDate,
		driver.PassportPhoto, driver.PassportScan, driver.Address, driver.ExperienceYears,
		driver.Status, driver.UpdatedAt, driver.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update driver: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("driver not found")
	}

	return nil
}
// Delete deletes a driver by ID
func (r *driverRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM drivers WHERE fldId = $1"

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete driver: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("driver not found")
	}

	return nil
}

// Search searches for drivers by query and optional status filter
func (r *driverRepository) Search(ctx context.Context, query string, status *models.DriverStatus, limit, offset int) ([]models.Driver, int64, error) {
	// Build search query
	searchQuery := `
		SELECT
			fldId, fldFirstName, fldLastName, fldMiddleName, fldPhone, fldEmail,
			fldBirthDate, fldPhoto, fldDriverLicenseNumber, fldDriverLicenseIssueDate,
			fldDriverLicenseExpiryDate, fldDriverLicensePhoto, fldDriverLicenseScan,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate,
			fldPassportPhoto, fldPassportScan, fldAddress, fldExperienceYears,
			fldStatus, fldCreatedAt, fldUpdatedAt
		FROM drivers
		WHERE ($1 = '' OR
			LOWER(fldFirstName) LIKE LOWER($1) OR
			LOWER(fldLastName) LIKE LOWER($1) OR
			LOWER(fldMiddleName) LIKE LOWER($1) OR
			LOWER(fldPhone) LIKE LOWER($1) OR
			LOWER(fldEmail) LIKE LOWER($1) OR
			LOWER(fldDriverLicenseNumber) LIKE LOWER($1))
	`

	args := []interface{}{"%" + query + "%"}
	argCount := 1

	if status != nil {
		searchQuery += fmt.Sprintf(" AND fldStatus = $%d", argCount+1)
		args = append(args, *status)
		argCount++
	}

	// Get total count
	countQuery := "SELECT COUNT(*) FROM drivers WHERE " + searchQuery[45:] // Remove SELECT ... FROM drivers WHERE
	var total int64
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count search results: %w", err)
	}

	// Add ordering and pagination
	searchQuery += " ORDER BY fldCreatedAt DESC"
	searchQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCount+1, argCount+2)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, searchQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to search drivers: %w", err)
	}
	defer rows.Close()

	drivers := []models.Driver{}
	for rows.Next() {
		driver := models.Driver{}
		err := rows.Scan(
			&driver.ID, &driver.FirstName, &driver.LastName, &driver.MiddleName, &driver.Phone, &driver.Email,
			&driver.BirthDate, &driver.Photo, &driver.DriverLicenseNumber, &driver.DriverLicenseIssueDate,
			&driver.DriverLicenseExpiryDate, &driver.DriverLicensePhoto, &driver.DriverLicenseScan,
			&driver.PassportSeries, &driver.PassportNumber, &driver.PassportIssueDate,
			&driver.PassportPhoto, &driver.PassportScan, &driver.Address, &driver.ExperienceYears,
			&driver.Status, &driver.CreatedAt, &driver.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan driver: %w", err)
		}
		drivers = append(drivers, driver)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating drivers: %w", err)
	}

	return drivers, total, nil
}
// GetByPhone retrieves a driver by phone number
func (r *driverRepository) GetByPhone(ctx context.Context, phone string) (*models.Driver, error) {
	query := `
		SELECT
			fldId, fldFirstName, fldLastName, fldMiddleName, fldPhone, fldEmail,
			fldBirthDate, fldPhoto, fldDriverLicenseNumber, fldDriverLicenseIssueDate,
			fldDriverLicenseExpiryDate, fldDriverLicensePhoto, fldDriverLicenseScan,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate,
			fldPassportPhoto, fldPassportScan, fldAddress, fldExperienceYears,
			fldStatus, fldCreatedAt, fldUpdatedAt
		FROM drivers
		WHERE fldPhone = $1
	`

	driver := &models.Driver{}
	err := r.db.QueryRowContext(ctx, query, phone).Scan(
		&driver.ID, &driver.FirstName, &driver.LastName, &driver.MiddleName, &driver.Phone, &driver.Email,
		&driver.BirthDate, &driver.Photo, &driver.DriverLicenseNumber, &driver.DriverLicenseIssueDate,
		&driver.DriverLicenseExpiryDate, &driver.DriverLicensePhoto, &driver.DriverLicenseScan,
		&driver.PassportSeries, &driver.PassportNumber, &driver.PassportIssueDate,
		&driver.PassportPhoto, &driver.PassportScan, &driver.Address, &driver.ExperienceYears,
		&driver.Status, &driver.CreatedAt, &driver.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get driver by phone: %w", err)
	}

	return driver, nil
}
// GetByLicenseNumber retrieves a driver by driver license number
func (r *driverRepository) GetByLicenseNumber(ctx context.Context, licenseNumber string) (*models.Driver, error) {
	query := `
		SELECT
			fldId, fldFirstName, fldLastName, fldMiddleName, fldPhone, fldEmail,
			fldBirthDate, fldPhoto, fldDriverLicenseNumber, fldDriverLicenseIssueDate,
			fldDriverLicenseExpiryDate, fldDriverLicensePhoto, fldDriverLicenseScan,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate,
			fldPassportPhoto, fldPassportScan, fldAddress, fldExperienceYears,
			fldStatus, fldCreatedAt, fldUpdatedAt
		FROM drivers
		WHERE fldDriverLicenseNumber = $1
	`

	driver := &models.Driver{}
	err := r.db.QueryRowContext(ctx, query, licenseNumber).Scan(
		&driver.ID, &driver.FirstName, &driver.LastName, &driver.MiddleName, &driver.Phone, &driver.Email,
		&driver.BirthDate, &driver.Photo, &driver.DriverLicenseNumber, &driver.DriverLicenseIssueDate,
		&driver.DriverLicenseExpiryDate, &driver.DriverLicensePhoto, &driver.DriverLicenseScan,
		&driver.PassportSeries, &driver.PassportNumber, &driver.PassportIssueDate,
		&driver.PassportPhoto, &driver.PassportScan, &driver.Address, &driver.ExperienceYears,
		&driver.Status, &driver.CreatedAt, &driver.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get driver by license number: %w", err)
	}

	return driver, nil
}
// GetByEmail retrieves a driver by email
func (r *driverRepository) GetByEmail(ctx context.Context, email string) (*models.Driver, error) {
	query := `
		SELECT
			fldId, fldFirstName, fldLastName, fldMiddleName, fldPhone, fldEmail,
			fldBirthDate, fldPhoto, fldDriverLicenseNumber, fldDriverLicenseIssueDate,
			fldDriverLicenseExpiryDate, fldDriverLicensePhoto, fldDriverLicenseScan,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate,
			fldPassportPhoto, fldPassportScan, fldAddress, fldExperienceYears,
			fldStatus, fldCreatedAt, fldUpdatedAt
		FROM drivers
		WHERE fldEmail = $1
	`

	driver := &models.Driver{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&driver.ID, &driver.FirstName, &driver.LastName, &driver.MiddleName, &driver.Phone, &driver.Email,
		&driver.BirthDate, &driver.Photo, &driver.DriverLicenseNumber, &driver.DriverLicenseIssueDate,
		&driver.DriverLicenseExpiryDate, &driver.DriverLicensePhoto, &driver.DriverLicenseScan,
		&driver.PassportSeries, &driver.PassportNumber, &driver.PassportIssueDate,
		&driver.PassportPhoto, &driver.PassportScan, &driver.Address, &driver.ExperienceYears,
		&driver.Status, &driver.CreatedAt, &driver.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get driver by email: %w", err)
	}

	return driver, nil
}
// Close closes the database connection
func (r *driverRepository) Close() error {
	return r.db.Close()
}
