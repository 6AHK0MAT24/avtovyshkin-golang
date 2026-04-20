package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"vehicles-service/internal/models"
	"vehicles-service/internal/repository"
	"vehicles-service/internal/utils"
)// VehicleService defines interface for vehicle business logic
type VehicleService interface {
	CreateVehicle(ctx context.Context, req *models.CreateVehicleRequest) (*models.Vehicle, error)
	GetVehicle(ctx context.Context, id uuid.UUID) (*models.Vehicle, error)
	GetVehicles(ctx context.Context, page, pageSize int) ([]*models.Vehicle, int64, error)
	UpdateVehicle(ctx context.Context, id uuid.UUID, req *models.UpdateVehicleRequest) (*models.Vehicle, error)
	DeleteVehicle(ctx context.Context, id uuid.UUID) error
	SearchVehicles(ctx context.Context, query string, page, pageSize int) ([]*models.Vehicle, int64, error)
	UploadVehicleImages(ctx context.Context, id uuid.UUID, files []FileData) ([]string, error)
	DeleteVehicleImage(ctx context.Context, id uuid.UUID, index int) error
	SetMainImageIndex(ctx context.Context, id uuid.UUID, index int) error
}

// FileData represents uploaded file data
type FileData struct {
	Data     []byte
	Filename string
}

// vehicleService implements VehicleService interface
type vehicleService struct {
	repo repository.VehicleRepository
}

// NewVehicleService creates a new vehicle service
func NewVehicleService(repo repository.VehicleRepository) VehicleService {
	return &vehicleService{repo: repo}
}

// CreateVehicle creates a new vehicle with validation
func (s *vehicleService) CreateVehicle(ctx context.Context, req *models.CreateVehicleRequest) (*models.Vehicle, error) {
	// Validate VIN (17 characters, alphanumeric without I, O, Q)
	if !utils.ValidateVIN(req.VIN) {
		return nil, fmt.Errorf("invalid VIN format: must be 17 alphanumeric characters without I, O, Q")
	}

	// Validate garage number
	if !utils.ValidateGarageNumber(req.GarageNumber) {
		return nil, fmt.Errorf("invalid garage number format")
	}

	// Validate height (> 0)
	if req.Height <= 0 {
		return nil, fmt.Errorf("height must be greater than 0")
	}

	// Validate type if provided
	if req.Type != nil {
		validTypes := map[string]bool{
			"Телескопическая":          true,
			"Телескоп + колено":        true,
			"Телескоп + стрела и рукоять": true,
		}
		if !validTypes[*req.Type] {
			return nil, fmt.Errorf("invalid vehicle type: must be one of Телескопическая, Телескоп + колено, Телескоп + стрела и рукоять")
		}
	}

	// Validate status if provided
	if req.Status != nil && !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid vehicle status")
	}

	// Check if garage number already exists
	existingVehicle, err := s.repo.GetByGarageNumber(ctx, req.GarageNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing vehicle: %w", err)
	}
	if existingVehicle != nil {
		return nil, fmt.Errorf("vehicle with this garage number already exists")
	}

	// Check if VIN already exists
	existingVehicle, err = s.repo.GetByVIN(ctx, req.VIN)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing vehicle: %w", err)
	}
	if existingVehicle != nil {
		return nil, fmt.Errorf("vehicle with this VIN already exists")
	}

	// Create vehicle
	vehicle := req.ToVehicle()
	vehicle.CreatedAt = time.Now()
	vehicle.UpdatedAt = time.Now()

	if err := s.repo.Create(ctx, vehicle); err != nil {
		return nil, fmt.Errorf("failed to create vehicle: %w", err)
	}

	return vehicle, nil
}

// GetVehicles retrieves all vehicles with pagination
func (s *vehicleService) GetVehicles(ctx context.Context, page, pageSize int) ([]*models.Vehicle, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	vehicles, total, err := s.repo.GetAll(ctx, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get vehicles: %w", err)
	}

	return vehicles, total, nil
}

// UpdateVehicle updates an existing vehicle with validation
func (s *vehicleService) UpdateVehicle(ctx context.Context, id uuid.UUID, req *models.UpdateVehicleRequest) (*models.Vehicle, error) {
	// Get existing vehicle
	vehicle, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get vehicle: %w", err)
	}

	// Validate VIN if provided
	if req.VIN != nil {
		if !utils.ValidateVIN(*req.VIN) {
			return nil, fmt.Errorf("invalid VIN format: must be 17 alphanumeric characters without I, O, Q")
		}
	}

	// Validate garage number if provided
	if req.GarageNumber != nil && !utils.ValidateGarageNumber(*req.GarageNumber) {
		return nil, fmt.Errorf("invalid garage number format")
	}

	// Validate height if provided
	if req.Height != nil && *req.Height <= 0 {
		return nil, fmt.Errorf("height must be greater than 0")
	}

	// Validate type if provided
	if req.Type != nil {
		validTypes := map[string]bool{
			"Телескопическая":          true,
			"Телескоп + колено":        true,
			"Телескоп + стрела и рукоять": true,
		}
		if !validTypes[*req.Type] {
			return nil, fmt.Errorf("invalid vehicle type: must be one of Телескопическая, Телескоп + колено, Телескоп + стрела и рукоять")
		}
	}

	// Validate status if provided
	if req.Status != nil && !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid vehicle status")
	}

	// Validate main image index if provided
	if req.MainImageIndex != nil {
		if *req.MainImageIndex < 0 {
			return nil, fmt.Errorf("main image index must be non-negative")
		}
		if len(vehicle.ImgArray) > 0 && *req.MainImageIndex >= len(vehicle.ImgArray) {
			return nil, fmt.Errorf("main image index out of range")
		}
	}

	// Check if garage number already exists (excluding current vehicle)
	if req.GarageNumber != nil && *req.GarageNumber != vehicle.GarageNumber {
		existingVehicle, err := s.repo.GetByGarageNumber(ctx, *req.GarageNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to check existing vehicle: %w", err)
		}
		if existingVehicle != nil && existingVehicle.ID != id {
			return nil, fmt.Errorf("vehicle with this garage number already exists")
		}
	}

	// Check if VIN already exists (excluding current vehicle)
	if req.VIN != nil && *req.VIN != vehicle.VIN {
		existingVehicle, err := s.repo.GetByVIN(ctx, *req.VIN)
		if err != nil {
			return nil, fmt.Errorf("failed to check existing vehicle: %w", err)
		}
		if existingVehicle != nil && existingVehicle.ID != id {
			return nil, fmt.Errorf("vehicle with this VIN already exists")
		}
	}

	// Update vehicle
	vehicle.UpdateVehicle(req)
	vehicle.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, vehicle); err != nil {
		return nil, fmt.Errorf("failed to update vehicle: %w", err)
	}

	return vehicle, nil
}

// DeleteVehicle deletes a vehicle by ID and removes associated files
func (s *vehicleService) DeleteVehicle(ctx context.Context, id uuid.UUID) error {
	// Get vehicle data to delete files
	vehicle, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get vehicle: %w", err)
	}

	// Delete all vehicle images
	for _, imagePath := range vehicle.ImgArray {
		if err := os.Remove(imagePath); err != nil && !os.IsNotExist(err) {
			// Log error but continue with deletion
			fmt.Printf("warning: failed to delete image %s: %v\n", imagePath, err)
		}
	}

	// Delete vehicle from database
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete vehicle: %w", err)
	}

	return nil
}

// SearchVehicles searches vehicles by height and garage number
func (s *vehicleService) SearchVehicles(ctx context.Context, query string, page, pageSize int) ([]*models.Vehicle, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	vehicles, total, err := s.repo.Search(ctx, query, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to search vehicles: %w", err)
	}

	return vehicles, total, nil
}

// UploadVehicleImages uploads multiple images for a vehicle
func (s *vehicleService) UploadVehicleImages(ctx context.Context, id uuid.UUID, files []FileData) ([]string, error) {
	// Get vehicle
	vehicle, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get vehicle: %w", err)
	}

	// Determine upload directory using vehicle ID
	uploadDir := filepath.Join("uploads", "cars", vehicle.ID.String())
	// Create directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Upload files
	var newImages []string
	allowedTypes := []string{"jpg", "jpeg", "png", "gif", "webp"}
	for _, file := range files {
		// Validate file type
		if !utils.ValidateFileType(file.Filename, allowedTypes) {
			return nil, fmt.Errorf("invalid file type: %s", file.Filename)
		}
		// Generate unique filename
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s_%s%s", uuid.New().String(), strings.TrimSuffix(file.Filename, ext), ext)
		filePath := filepath.Join(uploadDir, filename)

		// Save file
		if err := os.WriteFile(filePath, file.Data, 0644); err != nil {
			return nil, fmt.Errorf("failed to save file: %w", err)
		}

		// Convert to relative path with forward slashes and remove duplicates
		relativePath := strings.ReplaceAll(filePath, "\\", "/")
		// Remove duplicate slashes
		for strings.Contains(relativePath, "//") {
			relativePath = strings.ReplaceAll(relativePath, "//", "/")
		}
		if !strings.HasPrefix(relativePath, "/") {
			relativePath = "/" + relativePath
		}

		newImages = append(newImages, relativePath)
	}
	// Update vehicle with new images
	vehicle.ImgArray = append(vehicle.ImgArray, newImages...)
	vehicle.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, vehicle); err != nil {
		// Rollback: delete uploaded files
		for _, imagePath := range newImages {
			os.Remove(imagePath)
		}
		return nil, fmt.Errorf("failed to update vehicle: %w", err)
	}

	return newImages, nil
}

// DeleteVehicleImage deletes a vehicle image by index
func (s *vehicleService) DeleteVehicleImage(ctx context.Context, id uuid.UUID, index int) error {
	// Get vehicle
	vehicle, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get vehicle: %w", err)
	}

	// Validate index
	if index < 0 || index >= len(vehicle.ImgArray) {
		return fmt.Errorf("image index out of range")
	}

	// Delete file
	imagePath := vehicle.ImgArray[index]
	if err := os.Remove(imagePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete image file: %w", err)
	}

	// Remove image from array
	vehicle.ImgArray = append(vehicle.ImgArray[:index], vehicle.ImgArray[index+1:]...)

	// Adjust main image index if necessary
	if vehicle.MainImageIndex >= len(vehicle.ImgArray) {
		vehicle.MainImageIndex = len(vehicle.ImgArray) - 1
	}
	if len(vehicle.ImgArray) == 0 {
		vehicle.MainImageIndex = 0
	}

	vehicle.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, vehicle); err != nil {
		return fmt.Errorf("failed to update vehicle: %w", err)
	}

	return nil
}

// SetMainImageIndex sets the main image index for a vehicle
func (s *vehicleService) SetMainImageIndex(ctx context.Context, id uuid.UUID, index int) error {
	// Get vehicle
	vehicle, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get vehicle: %w", err)
	}

	// Validate index
	if index < 0 || index >= len(vehicle.ImgArray) {
		return fmt.Errorf("image index out of range")
	}

	// Set main image index
	vehicle.MainImageIndex = index
	vehicle.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, vehicle); err != nil {
		return fmt.Errorf("failed to update vehicle: %w", err)
	}

	return nil
}
