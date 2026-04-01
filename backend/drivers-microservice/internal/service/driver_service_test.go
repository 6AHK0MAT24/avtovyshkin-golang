package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"drivers-service/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
// MockDriverRepository is a mock implementation of DriverRepository
type MockDriverRepository struct {
	mock.Mock
}

func (m *MockDriverRepository) Create(ctx context.Context, driver *models.Driver) error {
	args := m.Called(ctx, driver)
	return args.Error(0)
}

func (m *MockDriverRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Driver, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Driver), args.Error(1)
}

func (m *MockDriverRepository) GetAll(ctx context.Context, limit, offset int) ([]models.Driver, int64, error) {
	args := m.Called(ctx, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]models.Driver), args.Get(1).(int64), args.Error(2)
}

func (m *MockDriverRepository) Update(ctx context.Context, driver *models.Driver) error {
	args := m.Called(ctx, driver)
	return args.Error(0)
}

func (m *MockDriverRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockDriverRepository) Search(ctx context.Context, query string, status *models.DriverStatus, limit, offset int) ([]models.Driver, int64, error) {
	args := m.Called(ctx, query, status, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]models.Driver), args.Get(1).(int64), args.Error(2)
}

func (m *MockDriverRepository) GetByPhone(ctx context.Context, phone string) (*models.Driver, error) {
	args := m.Called(ctx, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Driver), args.Error(1)
}

func (m *MockDriverRepository) GetByLicenseNumber(ctx context.Context, licenseNumber string) (*models.Driver, error) {
	args := m.Called(ctx, licenseNumber)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Driver), args.Error(1)
}

func (m *MockDriverRepository) GetByEmail(ctx context.Context, email string) (*models.Driver, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Driver), args.Error(1)
}

func (m *MockDriverRepository) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestDriverService_CreateDriver_Success(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	licenseIssueDate := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)
	licenseExpiryDate := time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)
	passportIssueDate := time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC)

	middleName := "Иванович"
	req := &models.CreateDriverRequest{
		FirstName:                "Иван",
		LastName:                 "Иванов",
		MiddleName:               &middleName,
		Phone:                    "+79001234567",
		Email:                    stringPtr("ivan@example.com"),
		BirthDate:                &birthDate,
		DriverLicenseNumber:      "1234567890",
		DriverLicenseIssueDate:   licenseIssueDate,
		DriverLicenseExpiryDate:  licenseExpiryDate,
		PassportSeries:           stringPtr("1234"),
		PassportNumber:           stringPtr("567890"),
		PassportIssueDate:        &passportIssueDate,
		Address:                  stringPtr("г. Москва, ул. Тестовая, д. 1"),
		ExperienceYears:          10,
		Status:                   models.DriverStatusActive,
	}

	mockRepo.On("GetByPhone", mock.Anything, req.Phone).Return(nil, nil)
	mockRepo.On("GetByEmail", mock.Anything, *req.Email).Return(nil, nil)
	mockRepo.On("GetByLicenseNumber", mock.Anything, req.DriverLicenseNumber).Return(nil, nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Driver")).Return(nil)

	driver, err := service.CreateDriver(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, driver)
	assert.Equal(t, req.FirstName, driver.FirstName)
	assert.Equal(t, req.LastName, driver.LastName)
	mockRepo.AssertExpectations(t)
}

func TestDriverService_CreateDriver_InvalidPhone(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	req := &models.CreateDriverRequest{
		FirstName:           "Иван",
		LastName:            "Иванов",
		Phone:               "invalid-phone",
		DriverLicenseNumber: "1234567890",
	}

	driver, err := service.CreateDriver(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Contains(t, err.Error(), "invalid phone number format")
}

func TestDriverService_CreateDriver_InvalidEmail(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	req := &models.CreateDriverRequest{
		FirstName:           "Иван",
		LastName:            "Иванов",
		Phone:               "+79001234567",
		Email:               stringPtr("invalid-email"),
		DriverLicenseNumber: "1234567890",
	}

	driver, err := service.CreateDriver(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Contains(t, err.Error(), "invalid email format")
}

func TestDriverService_CreateDriver_DuplicatePhone(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	req := &models.CreateDriverRequest{
		FirstName:           "Иван",
		LastName:            "Иванов",
		Phone:               "+79001234567",
		DriverLicenseNumber: "1234567890",
	}

	existingDriver := &models.Driver{ID: uuid.New()}
	mockRepo.On("GetByPhone", mock.Anything, req.Phone).Return(existingDriver, nil)

	driver, err := service.CreateDriver(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Contains(t, err.Error(), "driver with this phone number already exists")
	mockRepo.AssertExpectations(t)
}

func TestDriverService_CreateDriver_DuplicateEmail(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	req := &models.CreateDriverRequest{
		FirstName:           "Иван",
		LastName:            "Иванов",
		Phone:               "+79001234567",
		Email:               stringPtr("ivan@example.com"),
		DriverLicenseNumber: "1234567890",
	}

	mockRepo.On("GetByPhone", mock.Anything, req.Phone).Return(nil, nil)
	existingDriver := &models.Driver{ID: uuid.New()}
	mockRepo.On("GetByEmail", mock.Anything, *req.Email).Return(existingDriver, nil)

	driver, err := service.CreateDriver(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Contains(t, err.Error(), "driver with this email already exists")
	mockRepo.AssertExpectations(t)
}

func TestDriverService_CreateDriver_DuplicateLicense(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	req := &models.CreateDriverRequest{
		FirstName:           "Иван",
		LastName:            "Иванов",
		Phone:               "+79001234567",
		DriverLicenseNumber: "1234567890",
	}

	mockRepo.On("GetByPhone", mock.Anything, req.Phone).Return(nil, nil)
	existingDriver := &models.Driver{ID: uuid.New()}
	mockRepo.On("GetByLicenseNumber", mock.Anything, req.DriverLicenseNumber).Return(existingDriver, nil)

	driver, err := service.CreateDriver(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Contains(t, err.Error(), "driver with this license number already exists")
	mockRepo.AssertExpectations(t)
}

func TestDriverService_GetDriver_Success(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	driverID := uuid.New()
	expectedDriver := &models.Driver{
		ID:        driverID,
		FirstName: "Иван",
		LastName:  "Иванов",
	}

	mockRepo.On("GetByID", mock.Anything, driverID).Return(expectedDriver, nil)

	driver, err := service.GetDriver(context.Background(), driverID)

	assert.NoError(t, err)
	assert.NotNil(t, driver)
	assert.Equal(t, expectedDriver.ID, driver.ID)
	mockRepo.AssertExpectations(t)
}

func TestDriverService_GetDriver_NotFound(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	driverID := uuid.New()
	mockRepo.On("GetByID", mock.Anything, driverID).Return(nil, errors.New("driver not found"))

	driver, err := service.GetDriver(context.Background(), driverID)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Contains(t, err.Error(), "failed to get driver")
	mockRepo.AssertExpectations(t)
}

func TestDriverService_GetDrivers_Success(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	drivers := []models.Driver{
		{ID: uuid.New(), FirstName: "Иван", LastName: "Иванов"},
		{ID: uuid.New(), FirstName: "Петр", LastName: "Петров"},
	}

	mockRepo.On("GetAll", mock.Anything, 10, 0).Return(drivers, int64(2), nil)

	result, total, err := service.GetDrivers(context.Background(), 1, 10)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(2), total)
	mockRepo.AssertExpectations(t)
}

func TestDriverService_GetDrivers_DefaultPagination(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	mockRepo.On("GetAll", mock.Anything, 10, 0).Return([]models.Driver{}, int64(0), nil)

	_, _, err := service.GetDrivers(context.Background(), 0, 0)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDriverService_UpdateDriver_Success(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	driverID := uuid.New()
	existingDriver := &models.Driver{
		ID:        driverID,
		FirstName: "Иван",
		LastName:  "Иванов",
		Phone:     "+79001234567",
	}

	newPhone := "+79009876543"
	req := &models.UpdateDriverRequest{
		FirstName: stringPtr("Петр"),
		Phone:     &newPhone,
	}

	mockRepo.On("GetByID", mock.Anything, driverID).Return(existingDriver, nil)
	mockRepo.On("GetByPhone", mock.Anything, newPhone).Return(nil, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*models.Driver")).Return(nil)

	driver, err := service.UpdateDriver(context.Background(), driverID, req)

	assert.NoError(t, err)
	assert.NotNil(t, driver)
	mockRepo.AssertExpectations(t)
}

func TestDriverService_UpdateDriver_NotFound(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	driverID := uuid.New()
	req := &models.UpdateDriverRequest{
		FirstName: stringPtr("Петр"),
	}

	mockRepo.On("GetByID", mock.Anything, driverID).Return(nil, errors.New("driver not found"))

	driver, err := service.UpdateDriver(context.Background(), driverID, req)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Contains(t, err.Error(), "failed to get driver")
	mockRepo.AssertExpectations(t)
}

func TestDriverService_DeleteDriver_Success(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	driverID := uuid.New()
	existingDriver := &models.Driver{ID: driverID}

	mockRepo.On("GetByID", mock.Anything, driverID).Return(existingDriver, nil)
	mockRepo.On("Delete", mock.Anything, driverID).Return(nil)

	err := service.DeleteDriver(context.Background(), driverID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDriverService_DeleteDriver_NotFound(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	driverID := uuid.New()
	mockRepo.On("GetByID", mock.Anything, driverID).Return(nil, errors.New("driver not found"))

	err := service.DeleteDriver(context.Background(), driverID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get driver")
	mockRepo.AssertExpectations(t)
}

func TestDriverService_SearchDrivers_Success(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	query := "Иван"
	status := models.DriverStatusActive
	drivers := []models.Driver{
		{ID: uuid.New(), FirstName: "Иван", LastName: "Иванов"},
	}

	mockRepo.On("Search", mock.Anything, query, &status, 10, 0).Return(drivers, int64(1), nil)
	result, total, err := service.SearchDrivers(context.Background(), query, &status, 1, 10)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), total)
	mockRepo.AssertExpectations(t)
}

func TestDriverService_SearchDrivers_WithoutStatus(t *testing.T) {
	mockRepo := new(MockDriverRepository)
	service := NewDriverService(mockRepo)

	query := "Иван"
	drivers := []models.Driver{
		{ID: uuid.New(), FirstName: "Иван", LastName: "Иванов"},
	}

	mockRepo.On("Search", mock.Anything, query, (*models.DriverStatus)(nil), 10, 0).Return(drivers, int64(1), nil)

	result, total, err := service.SearchDrivers(context.Background(), query, nil, 1, 10)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), total)
	mockRepo.AssertExpectations(t)
}

// Helper function
func stringPtr(s string) *string {
	return &s
}
