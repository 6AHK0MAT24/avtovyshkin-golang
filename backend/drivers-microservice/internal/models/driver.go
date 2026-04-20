package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Driver represents the driver entity in the database
type Driver struct {
	ID                       uuid.UUID  `json:"id" db:"fldId"`
	FirstName                string     `json:"firstName" db:"fldFirstName"`
	LastName                 string     `json:"lastName" db:"fldLastName"`
	MiddleName               *string    `json:"middleName,omitempty" db:"fldMiddleName"`
	Phone                    string     `json:"phone" db:"fldPhone"`
	Email                    *string    `json:"email,omitempty" db:"fldEmail"`
	BirthDate                *time.Time `json:"birthDate,omitempty" db:"fldBirthDate"`
	Photo                    *string    `json:"photo,omitempty" db:"fldPhoto"`
	DriverLicenseNumber      string     `json:"driverLicenseNumber" db:"fldDriverLicenseNumber"`
	DriverLicenseIssueDate   time.Time  `json:"driverLicenseIssueDate" db:"fldDriverLicenseIssueDate"`
	DriverLicenseExpiryDate  time.Time  `json:"driverLicenseExpiryDate" db:"fldDriverLicensePhoto"`
	DriverLicensePhoto       *string    `json:"driverLicensePhoto,omitempty" db:"fldDriverLicensePhoto"`
	DriverLicenseScan        *string    `json:"driverLicenseScan,omitempty" db:"fldDriverLicenseScan"`
	PassportSeries           *string    `json:"passportSeries,omitempty" db:"fldPassportSeries"`
	PassportNumber           *string    `json:"passportNumber,omitempty" db:"fldPassportNumber"`
	PassportIssueDate        *time.Time `json:"passportIssueDate,omitempty" db:"fldPassportIssueDate"`
	PassportPhoto            *string    `json:"passportPhoto,omitempty" db:"fldPassportPhoto"`
	PassportScan             *string    `json:"passportScan,omitempty" db:"fldPassportScan"`
	Address                  *string    `json:"address,omitempty" db:"fldAddress"`
	ExperienceYears          int        `json:"experienceYears" db:"fldExperienceYears"`
	Status                   DriverStatus `json:"status" db:"fldStatus"`
	CreatedAt                time.Time  `json:"createdAt" db:"fldCreatedAt"`
	UpdatedAt                time.Time  `json:"updatedAt" db:"fldUpdatedAt"`
}

// DriverStatus represents the status of a driver
type DriverStatus string

const (
	DriverStatusActive   DriverStatus = "active"
	DriverStatusInactive DriverStatus = "inactive"
	DriverStatusBlocked  DriverStatus = "blocked"
)

// Value implements the driver.Valuer interface for DriverStatus
func (ds DriverStatus) Value() (driver.Value, error) {
	return string(ds), nil
}

// Scan implements the sql.Scanner interface for DriverStatus
func (ds *DriverStatus) Scan(value interface{}) error {
	if value == nil {
		*ds = DriverStatusActive
		return nil
	}
	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return errors.New("invalid type for DriverStatus")
	}
	*ds = DriverStatus(str)
	return nil
}
// IsValid checks if the driver status is valid
func (ds DriverStatus) IsValid() bool {
	switch ds {
	case DriverStatusActive, DriverStatusInactive, DriverStatusBlocked:
		return true
	default:
		return false
	}
}

// CreateDriverRequest represents the request to create a new driver
type CreateDriverRequest struct {
	FirstName               string     `json:"firstName" validate:"required,min=2,max=100"`
	LastName                string     `json:"lastName" validate:"required,min=2,max=100"`
	MiddleName              *string    `json:"middleName,omitempty" validate:"omitempty,max=100"`
	Phone                   string     `json:"phone" validate:"required"`
	Email                   *string    `json:"email,omitempty" validate:"omitempty,email"`
	BirthDate               *time.Time `json:"birthDate,omitempty"`
	Photo                   *string    `json:"photo,omitempty"`
	DriverLicenseNumber     string     `json:"driverLicenseNumber" validate:"required"`
	DriverLicenseIssueDate  time.Time  `json:"driverLicenseIssueDate" validate:"required"`
	DriverLicenseExpiryDate time.Time  `json:"driverLicenseExpiryDate" validate:"required"`
	PassportSeries          *string    `json:"passportSeries,omitempty" validate:"omitempty,len=4"`
	PassportNumber          *string    `json:"passportNumber,omitempty" validate:"omitempty,len=6"`
	PassportIssueDate       *time.Time `json:"passportIssueDate,omitempty"`
	Address                 *string    `json:"address,omitempty"`
	ExperienceYears         int        `json:"experienceYears" validate:"min=0"`
	Status                  DriverStatus `json:"status" validate:"omitempty,oneof=active inactive blocked"`
}
// UpdateDriverRequest represents the request to update a driver
type UpdateDriverRequest struct {
	FirstName               *string    `json:"firstName,omitempty" validate:"omitempty,min=2,max=100"`
	LastName                *string    `json:"lastName,omitempty" validate:"omitempty,min=2,max=100"`
	MiddleName              *string    `json:"middleName,omitempty" validate:"omitempty,max=100"`
	Phone                   *string    `json:"phone,omitempty" validate:"omitempty"`
	Email                   *string    `json:"email,omitempty" validate:"omitempty,email"`
	BirthDate               *time.Time `json:"birthDate,omitempty"`
	Photo                   *string    `json:"photo,omitempty"`
	DriverLicenseNumber     *string    `json:"driverLicenseNumber,omitempty"`
	DriverLicenseIssueDate  *time.Time `json:"driverLicenseIssueDate,omitempty"`
	DriverLicenseExpiryDate *time.Time `json:"driverLicenseExpiryDate,omitempty"`
	PassportSeries          *string    `json:"passportSeries,omitempty" validate:"omitempty,len=4"`
	PassportNumber          *string    `json:"passportNumber,omitempty" validate:"omitempty,len=6"`
	PassportIssueDate       *time.Time `json:"passportIssueDate,omitempty"`
	Address                 *string    `json:"address,omitempty"`
	ExperienceYears         *int       `json:"experienceYears,omitempty" validate:"omitempty,min=0"`
	Status                  *DriverStatus `json:"status,omitempty" validate:"omitempty,oneof=active inactive blocked"`
}
// DriverResponse represents the response for a driver
type DriverResponse struct {
	ID                       uuid.UUID  `json:"id"`
	FirstName                string     `json:"firstName"`
	LastName                 string     `json:"lastName"`
	MiddleName               *string    `json:"middleName,omitempty"`
	Phone                    string     `json:"phone"`
	Email                    *string    `json:"email,omitempty"`
	BirthDate                *time.Time `json:"birthDate,omitempty"`
	Photo                    *string    `json:"photo,omitempty"`
	DriverLicenseNumber      string     `json:"driverLicenseNumber"`
	DriverLicenseIssueDate   time.Time  `json:"driverLicenseIssueDate"`
	DriverLicenseExpiryDate  time.Time  `json:"driverLicenseExpiryDate"`
	DriverLicensePhoto       *string    `json:"driverLicensePhoto,omitempty"`
	DriverLicenseScan        *string    `json:"driverLicenseScan,omitempty"`
	PassportSeries           *string    `json:"passportSeries,omitempty"`
	PassportNumber           *string    `json:"passportNumber,omitempty"`
	PassportIssueDate        *time.Time `json:"passportIssueDate,omitempty"`
	PassportPhoto            *string    `json:"passportPhoto,omitempty"`
	PassportScan             *string    `json:"passportScan,omitempty"`
	Address                  *string    `json:"address,omitempty"`
	ExperienceYears          int        `json:"experienceYears"`
	Status                   DriverStatus `json:"status"`
	CreatedAt                time.Time  `json:"createdAt"`
	UpdatedAt                time.Time  `json:"updatedAt"`
}
// DriverListResponse represents the response for a list of drivers
type DriverListResponse struct {
	Drivers []DriverResponse `json:"drivers"`
	Total   int64            `json:"total"`
	Page    int              `json:"page"`
	PerPage int              `json:"perPage"`
}

// ToResponse converts a Driver to DriverResponse
func (d *Driver) ToResponse() DriverResponse {
	return DriverResponse{
		ID:                       d.ID,
		FirstName:                d.FirstName,
		LastName:                 d.LastName,
		MiddleName:               d.MiddleName,
		Phone:                    d.Phone,
		Email:                    d.Email,
		BirthDate:                d.BirthDate,
		Photo:                    d.Photo,
		DriverLicenseNumber:      d.DriverLicenseNumber,
		DriverLicenseIssueDate:   d.DriverLicenseIssueDate,
		DriverLicenseExpiryDate:  d.DriverLicenseExpiryDate,
		DriverLicensePhoto:       d.DriverLicensePhoto,
		DriverLicenseScan:        d.DriverLicenseScan,
		PassportSeries:           d.PassportSeries,
		PassportNumber:           d.PassportNumber,
		PassportIssueDate:        d.PassportIssueDate,
		PassportPhoto:            d.PassportPhoto,
		PassportScan:             d.PassportScan,
		Address:                  d.Address,
		ExperienceYears:          d.ExperienceYears,
		Status:                   d.Status,
		CreatedAt:                d.CreatedAt,
		UpdatedAt:                d.UpdatedAt,
	}
}
// ToDriver converts CreateDriverRequest to Driver
func (r *CreateDriverRequest) ToDriver() *Driver {
	status := r.Status
	if status == "" {
		status = DriverStatusActive
	}

	return &Driver{
		ID:                       uuid.New(),
		FirstName:                r.FirstName,
		LastName:                 r.LastName,
		MiddleName:               r.MiddleName,
		Phone:                    r.Phone,
		Email:                    r.Email,
		BirthDate:                r.BirthDate,
		Photo:                    r.Photo,
		DriverLicenseNumber:      r.DriverLicenseNumber,
		DriverLicenseIssueDate:   r.DriverLicenseIssueDate,
		DriverLicenseExpiryDate:  r.DriverLicenseExpiryDate,
		PassportSeries:           r.PassportSeries,
		PassportNumber:           r.PassportNumber,
		PassportIssueDate:        r.PassportIssueDate,
		Address:                  r.Address,
		ExperienceYears:          r.ExperienceYears,
		Status:                   status,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}
}
// UpdateDriver updates the driver fields from UpdateDriverRequest
func (d *Driver) UpdateDriver(req *UpdateDriverRequest) {
	if req.FirstName != nil {
		d.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		d.LastName = *req.LastName
	}
	if req.MiddleName != nil {
		d.MiddleName = req.MiddleName
	}
	if req.Phone != nil {
		d.Phone = *req.Phone
	}
	if req.Email != nil {
		d.Email = req.Email
	}
	if req.BirthDate != nil {
		d.BirthDate = req.BirthDate
	}
	if req.Photo != nil {
		d.Photo = req.Photo
	}
	if req.DriverLicenseNumber != nil {
		d.DriverLicenseNumber = *req.DriverLicenseNumber
	}
	if req.DriverLicenseIssueDate != nil {
		d.DriverLicenseIssueDate = *req.DriverLicenseIssueDate
	}
	if req.DriverLicenseExpiryDate != nil {
		d.DriverLicenseExpiryDate = *req.DriverLicenseExpiryDate
	}
	if req.PassportSeries != nil {
		d.PassportSeries = req.PassportSeries
	}
	if req.PassportNumber != nil {
		d.PassportNumber = req.PassportNumber
	}
	if req.PassportIssueDate != nil {
		d.PassportIssueDate = req.PassportIssueDate
	}
	if req.Address != nil {
		d.Address = req.Address
	}
	if req.ExperienceYears != nil {
		d.ExperienceYears = *req.ExperienceYears
	}
	if req.Status != nil {
		d.Status = *req.Status
	}
	d.UpdatedAt = time.Now()
}
// FileUploadRequest represents the request to upload a file
type FileUploadRequest struct {
	DriverID uuid.UUID `json:"driverId" validate:"required"`
	FileType string    `json:"fileType" validate:"required,oneof=photo license_scan passport_scan"`
}

// FileUploadResponse represents the response after file upload
type FileUploadResponse struct {
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
}

// SearchFilters represents the search filters for drivers
type SearchFilters struct {
	Query    string       `json:"query,omitempty"`
	Status   *DriverStatus `json:"status,omitempty"`
	Page     int          `json:"page"`
	PerPage  int          `json:"perPage"`
}

// WebSocketEvent represents a WebSocket event
type WebSocketEvent struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// DriverEvent represents a driver-related event
type DriverEvent struct {
	DriverID uuid.UUID `json:"driverId"`
	Action   string    `json:"action"` // created, updated, deleted
}

// MarshalJSON implements custom JSON marshaling for DriverEvent
func (de *DriverEvent) MarshalJSON() ([]byte, error) {
	type Alias DriverEvent
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(de),
	})
}

// UnmarshalJSON implements custom JSON unmarshaling for DriverEvent
func (de *DriverEvent) UnmarshalJSON(data []byte) error {
	type Alias DriverEvent
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(de),
	}
	return json.Unmarshal(data, &aux)
}
