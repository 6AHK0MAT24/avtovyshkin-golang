package service

import (
	"context"
	"fmt"
	"time"

	"drivers-service/internal/models"
	"drivers-service/internal/repository"
	"drivers-service/internal/utils"
	"github.com/google/uuid"
)// DriverService defines interface for driver business logic
type DriverService interface {
	CreateDriver(ctx context.Context, req *models.CreateDriverRequest) (*models.Driver, error)
	GetDriver(ctx context.Context, id uuid.UUID) (*models.Driver, error)
	GetDrivers(ctx context.Context, page, pageSize int) ([]models.Driver, int64, error)
	UpdateDriver(ctx context.Context, id uuid.UUID, req *models.UpdateDriverRequest) (*models.Driver, error)
	DeleteDriver(ctx context.Context, id uuid.UUID) error
	SearchDrivers(ctx context.Context, query string, status *models.DriverStatus, page, pageSize int) ([]models.Driver, int64, error)
	UploadDriverPhoto(ctx context.Context, id uuid.UUID, fileData []byte, filename string) (string, error)
	UploadDriverLicense(ctx context.Context, id uuid.UUID, fileData []byte, filename string, fileType string) (string, error)
	UploadDriverPassport(ctx context.Context, id uuid.UUID, fileData []byte, filename string, fileType string) (string, error)
	DeleteDriverPhoto(ctx context.Context, id uuid.UUID) error
}
// driverService implements DriverService interface
type driverService struct {
	repo repository.DriverRepository
}

// NewDriverService creates a new driver service
func NewDriverService(repo repository.DriverRepository) DriverService {
	return &driverService{repo: repo}
}

// CreateDriver creates a new driver with validation
func (s *driverService) CreateDriver(ctx context.Context, req *models.CreateDriverRequest) (*models.Driver, error) {
	// Validate phone format
	if !utils.ValidatePhone(req.Phone) {
		return nil, fmt.Errorf("invalid phone number format")
	}

	// Validate email format if provided
	if req.Email != nil && !utils.ValidateEmail(*req.Email) {
		return nil, fmt.Errorf("invalid email format")
	}

	// Validate driver license number
	if !utils.ValidateDriverLicense(req.DriverLicenseNumber) {
		return nil, fmt.Errorf("invalid driver license number format")
	}

	// Validate passport if provided
	if req.PassportSeries != nil || req.PassportNumber != nil {
		series := ""
		number := ""
		if req.PassportSeries != nil {
			series = *req.PassportSeries
		}
		if req.PassportNumber != nil {
			number = *req.PassportNumber
		}
		if !utils.ValidatePassport(series, number) {
			return nil, fmt.Errorf("invalid passport format")
		}
	}

	// Validate age (minimum 18 years)
	if req.BirthDate != nil && !utils.ValidateAge(req.BirthDate, 18) {
		return nil, fmt.Errorf("driver must be at least 18 years old")
	}

	// Validate status
	if req.Status != "" && !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid driver status")
	}

	// Check if phone already exists
	existingDriver, err := s.repo.GetByPhone(ctx, req.Phone)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing driver: %w", err)
	}
	if existingDriver != nil {
		return nil, fmt.Errorf("driver with this phone number already exists")
	}

	// Check if email already exists if provided
	if req.Email != nil {
		existingDriver, err := s.repo.GetByEmail(ctx, *req.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to check existing driver: %w", err)
		}
		if existingDriver != nil {
			return nil, fmt.Errorf("driver with this email already exists")
		}
	}

	// Check if driver license number already exists
	existingDriver, err = s.repo.GetByLicenseNumber(ctx, req.DriverLicenseNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing driver: %w", err)
	}
	if existingDriver != nil {
		return nil, fmt.Errorf("driver with this license number already exists")
	}

	// Create driver
	driver := req.ToDriver()
	if err := s.repo.Create(ctx, driver); err != nil {
		return nil, fmt.Errorf("failed to create driver: %w", err)
	}

	return driver, nil
}

// GetDriver retrieves a driver by ID
func (s *driverService) GetDriver(ctx context.Context, id uuid.UUID) (*models.Driver, error) {
	driver, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get driver: %w", err)
	}
	return driver, nil
}

// GetDrivers retrieves all drivers with pagination
func (s *driverService) GetDrivers(ctx context.Context, page, pageSize int) ([]models.Driver, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	drivers, total, err := s.repo.GetAll(ctx, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get drivers: %w", err)
	}

	return drivers, total, nil
}

// UpdateDriver updates an existing driver with validation
func (s *driverService) UpdateDriver(ctx context.Context, id uuid.UUID, req *models.UpdateDriverRequest) (*models.Driver, error) {
	// Get existing driver
	driver, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get driver: %w", err)
	}

	// Validate phone if provided
	if req.Phone != nil && !utils.ValidatePhone(*req.Phone) {
		return nil, fmt.Errorf("invalid phone number format")
	}

	// Validate email if provided
	if req.Email != nil && !utils.ValidateEmail(*req.Email) {
		return nil, fmt.Errorf("invalid email format")
	}

	// Validate driver license number if provided
	if req.DriverLicenseNumber != nil && !utils.ValidateDriverLicense(*req.DriverLicenseNumber) {
		return nil, fmt.Errorf("invalid driver license number format")
	}

	// Validate passport if provided
	if req.PassportSeries != nil || req.PassportNumber != nil {
		series := ""
		number := ""
		if req.PassportSeries != nil {
			series = *req.PassportSeries
		}
		if req.PassportNumber != nil {
			number = *req.PassportNumber
		}
		if !utils.ValidatePassport(series, number) {
			return nil, fmt.Errorf("invalid passport format")
		}
	}

	// Validate age if birth date is provided
	if req.BirthDate != nil && !utils.ValidateAge(req.BirthDate, 18) {
		return nil, fmt.Errorf("driver must be at least 18 years old")
	}

	// Validate status if provided
	if req.Status != nil && !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid driver status")
	}

	// Check if phone already exists (excluding current driver)
	if req.Phone != nil && *req.Phone != driver.Phone {
		existingDriver, err := s.repo.GetByPhone(ctx, *req.Phone)
		if err != nil {
			return nil, fmt.Errorf("failed to check existing driver: %w", err)
		}
		if existingDriver != nil && existingDriver.ID != id {
			return nil, fmt.Errorf("driver with this phone number already exists")
		}
	}

	// Check if email already exists (excluding current driver)
	if req.Email != nil && (*req.Email != "" || driver.Email != nil) {
		newEmail := *req.Email
		oldEmail := ""
		if driver.Email != nil {
			oldEmail = *driver.Email
		}
		if newEmail != "" && newEmail != oldEmail {
			existingDriver, err := s.repo.GetByEmail(ctx, newEmail)
			if err != nil {
				return nil, fmt.Errorf("failed to check existing driver: %w", err)
			}
			if existingDriver != nil && existingDriver.ID != id {
				return nil, fmt.Errorf("driver with this email already exists")
			}
		}
	}

	// Check if driver license number already exists (excluding current driver)
	if req.DriverLicenseNumber != nil && *req.DriverLicenseNumber != driver.DriverLicenseNumber {
		existingDriver, err := s.repo.GetByLicenseNumber(ctx, *req.DriverLicenseNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to check existing driver: %w", err)
		}
		if existingDriver != nil && existingDriver.ID != id {
			return nil, fmt.Errorf("driver with this license number already exists")
		}
	}

	// Update driver
	driver.UpdateDriver(req)

	if err := s.repo.Update(ctx, driver); err != nil {
		return nil, fmt.Errorf("failed to update driver: %w", err)
	}

	return driver, nil
}

// DeleteDriver deletes a driver by ID
func (s *driverService) DeleteDriver(ctx context.Context, id uuid.UUID) error {
	// Get driver data to delete files
	driver, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get driver: %w", err)
	}

	// Delete driver files
	uploadDir := "./uploads"
	if err := utils.DeleteDriverFiles(uploadDir, driver.ID.String()); err != nil {
		// Log error but continue with deletion
		// Files will be cleaned up later by a cleanup job if needed
		fmt.Printf("Warning: failed to delete driver files: %v\n", err)
	}

	// Delete driver from database
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete driver: %w", err)
	}

	return nil
}
// SearchDrivers searches for drivers by query and optional status filter
func (s *driverService) SearchDrivers(ctx context.Context, query string, status *models.DriverStatus, page, pageSize int) ([]models.Driver, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	drivers, total, err := s.repo.Search(ctx, query, status, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to search drivers: %w", err)
	}

	return drivers, total, nil
}

// UploadDriverPhoto uploads driver photo
func (s *driverService) UploadDriverPhoto(ctx context.Context, id uuid.UUID, fileData []byte, filename string) (string, error) {
	// Get driver
	driver, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return "", fmt.Errorf("failed to get driver: %w", err)
	}

	// Use driver-specific photo directory
	photoDir := "./uploads/drivers/photo"

	// Delete old photo if exists
	if driver.Photo != nil {
		if err := utils.DeleteFile(*driver.Photo); err != nil {
			return "", fmt.Errorf("failed to delete old photo: %w", err)
		}
	}

	// Generate unique filename using UUID
	uniqueFilename := utils.GenerateUniqueFilename(filename)

	// Save new photo to driver-specific directory
	filePath, err := utils.SaveFileFromBytes(fileData, uniqueFilename, photoDir)
	if err != nil {
		return "", fmt.Errorf("failed to save photo: %w", err)
	}

	// Update driver
	driver.Photo = &filePath
	driver.UpdatedAt = time.Now()
	if err := s.repo.Update(ctx, driver); err != nil {
	// Rollback: delete uploaded file
		utils.DeleteFile(filePath)
		return "", fmt.Errorf("failed to update driver: %w", err)
	}

	return filePath, nil
}
// UploadDriverLicense uploads driver license document
func (s *driverService) UploadDriverLicense(ctx context.Context, id uuid.UUID, fileData []byte, filename string, fileType string) (string, error) {
	// Get driver
	driver, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return "", fmt.Errorf("failed to get driver: %w", err)
	}

	// Use shared license directory
	licenseDir := "./uploads/drivers/license"

	// Delete old file if exists
	var oldFilePath *string
	if fileType == "photo" {
		oldFilePath = driver.DriverLicensePhoto
	} else if fileType == "scan" {
		oldFilePath = driver.DriverLicenseScan
	}

	if oldFilePath != nil {
		if err := utils.DeleteFile(*oldFilePath); err != nil {
			return "", fmt.Errorf("failed to delete old file: %w", err)
		}
	}

	// Generate unique filename
	uniqueFilename := utils.GenerateUniqueFilename(filename)

	// Save new file
	filePath, err := utils.SaveFileFromBytes(fileData, uniqueFilename, licenseDir)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Update driver
	if fileType == "photo" {
		driver.DriverLicensePhoto = &filePath
	} else if fileType == "scan" {
		driver.DriverLicenseScan = &filePath
	}
	driver.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, driver); err != nil {
		// Rollback: delete uploaded file
		utils.DeleteFile(filePath)
		return "", fmt.Errorf("failed to update driver: %w", err)
	}

	return filePath, nil
}
// UploadDriverPassport uploads driver passport document
func (s *driverService) UploadDriverPassport(ctx context.Context, id uuid.UUID, fileData []byte, filename string, fileType string) (string, error) {
	// Get driver
	driver, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return "", fmt.Errorf("failed to get driver: %w", err)
	}

	// Use shared passport directory
	passportDir := "./uploads/drivers/passport"

	// Delete old file if exists
	var oldFilePath *string
	if fileType == "photo" {
		oldFilePath = driver.PassportPhoto
	} else if fileType == "scan" {
		oldFilePath = driver.PassportScan
	}

	if oldFilePath != nil {
		if err := utils.DeleteFile(*oldFilePath); err != nil {
			return "", fmt.Errorf("failed to delete old file: %w", err)
		}
	}

	// Generate unique filename
	uniqueFilename := utils.GenerateUniqueFilename(filename)

	// Save new file
	filePath, err := utils.SaveFileFromBytes(fileData, uniqueFilename, passportDir)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Update driver
	if fileType == "photo" {
		driver.PassportPhoto = &filePath
	} else if fileType == "scan" {
		driver.PassportScan = &filePath
	}
	driver.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, driver); err != nil {
		// Rollback: delete uploaded file
		utils.DeleteFile(filePath)
		return "", fmt.Errorf("failed to update driver: %w", err)
	}

	return filePath, nil
}
// DeleteDriverPhoto deletes a driver's photo
func (s *driverService) DeleteDriverPhoto(ctx context.Context, id uuid.UUID) error {
	// Get driver
	driver, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get driver: %w", err)
	}

	// Check if driver has a photo
	if driver.Photo == nil {
		return fmt.Errorf("driver has no photo")
	}

	// Delete file from disk
	if err := utils.DeleteFile(*driver.Photo); err != nil {
		return fmt.Errorf("failed to delete photo file: %w", err)
	}

	// Update driver to remove photo reference
	driver.Photo = nil
	driver.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, driver); err != nil {
		return fmt.Errorf("failed to update driver: %w", err)
	}

	return nil
}
