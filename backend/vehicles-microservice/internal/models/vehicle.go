package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Vehicle represents the vehicle entity in the database
type Vehicle struct {
	ID                   uuid.UUID     `json:"id" db:"fldId"`
	GarageNumber         string        `json:"garageNumber" db:"fldGarageNumber"`
	VIN                  string        `json:"vin" db:"fldVIN"`
	Height               float64       `json:"height" db:"fldHeight"`
	Type                 *string       `json:"type,omitempty" db:"fldType"`
	Power                *float64      `json:"power,omitempty" db:"fldPower"`
	Price5               *float64      `json:"price5,omitempty" db:"fldPrice5"`
	Price22              *float64      `json:"price22,omitempty" db:"fldPrice22"`
	Description          *string       `json:"description,omitempty" db:"fldDescription"`
	Brand                *string       `json:"brand,omitempty" db:"fldBrand"`
	Machine              *string       `json:"machine,omitempty" db:"fldMachine"`
	Length               *float64      `json:"length,omitempty" db:"fldLength"`
	Width                *float64      `json:"width,omitempty" db:"fldWidth"`
	HeightTs             *float64      `json:"heightTs,omitempty" db:"fldHeight"`
	WidthWithSupports    *float64      `json:"widthWithSupports,omitempty" db:"fldWidthWithSupports"`
	Mass                 *float64      `json:"mass,omitempty" db:"fldMass"`
	CradleWidthFolded    *float64      `json:"cradleWidthFolded,omitempty" db:"fldCradleWidthFolded"`
	CradleWidthExtended  *float64      `json:"cradleWidthExtended,omitempty" db:"fldCradleWidthExtended"`
	CradleLengthFolded   *float64      `json:"cradleLengthFolded,omitempty" db:"fldCradleLengthFolded"`
	CradleLengthExtended *float64      `json:"cradleLengthExtended,omitempty" db:"fldCradleLengthExtended"`
	ImgArray             StringArray   `json:"imgArray" db:"fldImgArray"`
	MainImageIndex       int           `json:"mainImageIndex" db:"fldMainImageIndex"`
	Special              *string       `json:"special,omitempty" db:"fldSpecial"`
	RostechReg           bool          `json:"rostechReg" db:"fldRostechReg"`
	Status               VehicleStatus `json:"status" db:"fldStatus"`
	CreatedAt            time.Time     `json:"createdAt" db:"fldCreatedAt"`
	UpdatedAt            time.Time     `json:"updatedAt" db:"fldUpdatedAt"`
}

// VehicleStatus represents the status of a vehicle
type VehicleStatus string

const (
	VehicleStatusActive   VehicleStatus = "active"
	VehicleStatusInactive VehicleStatus = "inactive"
	VehicleStatusBlocked  VehicleStatus = "blocked"
)

// Value implements the driver.Valuer interface for VehicleStatus
func (vs VehicleStatus) Value() (driver.Value, error) {
	return string(vs), nil
}

// Scan implements the sql.Scanner interface for VehicleStatus
func (vs *VehicleStatus) Scan(value interface{}) error {
	if value == nil {
		*vs = VehicleStatusActive
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid type for VehicleStatus")
	}
	*vs = VehicleStatus(str)
	return nil
}

// IsValid checks if the vehicle status is valid
func (vs VehicleStatus) IsValid() bool {
	switch vs {
	case VehicleStatusActive, VehicleStatusInactive, VehicleStatusBlocked:
		return true
	default:
		return false
	}
}

// StringArray is a custom type for handling JSON arrays in PostgreSQL
type StringArray []string

// Value implements the driver.Valuer interface for StringArray
func (sa StringArray) Value() (driver.Value, error) {
	if sa == nil {
		return nil, nil
	}
	return json.Marshal(sa)
}

// Scan implements the sql.Scanner interface for StringArray
func (sa *StringArray) Scan(value interface{}) error {
	if value == nil {
		*sa = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for StringArray")
	}
	return json.Unmarshal(bytes, sa)
}

// CreateVehicleRequest represents the request to create a new vehicle
type CreateVehicleRequest struct {
	GarageNumber         string    `json:"garageNumber" validate:"required"`
	VIN                  string    `json:"vin" validate:"required,len=17"`
	Height               float64   `json:"height" validate:"required,gt=0"`
	Type                 *string   `json:"type,omitempty" validate:"omitempty,oneof=Телескопическая Телескоп + колено Телескоп + стрела и рукоять"`
	Power                *float64  `json:"power,omitempty" validate:"omitempty,gt=0"`
	Price5               *float64  `json:"price5,omitempty" validate:"omitempty,gt=0"`
	Price22              *float64  `json:"price22,omitempty" validate:"omitempty,gt=0"`
	Description          *string   `json:"description,omitempty"`
	Brand                *string   `json:"brand,omitempty"`
	Machine              *string   `json:"machine,omitempty"`
	Length               *float64  `json:"length,omitempty" validate:"omitempty,gt=0"`
	Width                *float64  `json:"width,omitempty" validate:"omitempty,gt=0"`
	HeightTs             *float64  `json:"heightTs,omitempty" validate:"omitempty,gt=0"`
	WidthWithSupports    *float64  `json:"widthWithSupports,omitempty" validate:"omitempty,gt=0"`
	Mass                 *float64  `json:"mass,omitempty" validate:"omitempty,gt=0"`
	CradleWidthFolded    *float64  `json:"cradleWidthFolded,omitempty" validate:"omitempty,gt=0"`
	CradleWidthExtended  *float64  `json:"cradleWidthExtended,omitempty" validate:"omitempty,gt=0"`
	CradleLengthFolded   *float64  `json:"cradleLengthFolded,omitempty" validate:"omitempty,gt=0"`
	CradleLengthExtended *float64  `json:"cradleLengthExtended,omitempty" validate:"omitempty,gt=0"`
	Special              *string   `json:"special,omitempty"`
	RostechReg           *bool     `json:"rostechReg,omitempty"`
	Status               *VehicleStatus `json:"status,omitempty" validate:"omitempty,oneof=active inactive blocked"`
}

// UpdateVehicleRequest represents the request to update a vehicle
type UpdateVehicleRequest struct {
	GarageNumber         *string         `json:"garageNumber,omitempty" validate:"omitempty"`
	VIN                  *string         `json:"vin,omitempty" validate:"omitempty,len=17"`
	Height               *float64        `json:"height,omitempty" validate:"omitempty,gt=0"`
	Type                 *string         `json:"type,omitempty" validate:"omitempty,oneof=Телескопическая Телескоп + колено Телескоп + стрела и рукоять"`
	Power                *float64        `json:"power,omitempty" validate:"omitempty,gt=0"`
	Price5               *float64        `json:"price5,omitempty" validate:"omitempty,gt=0"`
	Price22              *float64        `json:"price22,omitempty" validate:"omitempty,gt=0"`
	Description          *string         `json:"description,omitempty"`
	Brand                *string         `json:"brand,omitempty"`
	Machine              *string         `json:"machine,omitempty"`
	Length               *float64        `json:"length,omitempty" validate:"omitempty,gt=0"`
	Width                *float64        `json:"width,omitempty" validate:"omitempty,gt=0"`
	HeightTs             *float64        `json:"heightTs,omitempty" validate:"omitempty,gt=0"`
	WidthWithSupports    *float64        `json:"widthWithSupports,omitempty" validate:"omitempty,gt=0"`
	Mass                 *float64        `json:"mass,omitempty" validate:"omitempty,gt=0"`
	CradleWidthFolded    *float64        `json:"cradleWidthFolded,omitempty" validate:"omitempty,gt=0"`
	CradleWidthExtended  *float64        `json:"cradleWidthExtended,omitempty" validate:"omitempty,gt=0"`
	CradleLengthFolded   *float64        `json:"cradleLengthFolded,omitempty" validate:"omitempty,gt=0"`
	CradleLengthExtended *float64        `json:"cradleLengthExtended,omitempty" validate:"omitempty,gt=0"`
	Special              *string         `json:"special,omitempty"`
	RostechReg           *bool           `json:"rostechReg,omitempty"`
	Status               *VehicleStatus  `json:"status,omitempty" validate:"omitempty,oneof=active inactive blocked"`
	MainImageIndex       *int            `json:"mainImageIndex,omitempty" validate:"omitempty,gte=0"`
}

// VehicleResponse represents the response for a vehicle
type VehicleResponse struct {
	ID                   uuid.UUID     `json:"id"`
	GarageNumber         string        `json:"garageNumber"`
	VIN                  string        `json:"vin"`
	Height               float64       `json:"height"`
	Type                 *string       `json:"type,omitempty"`
	Power                *float64      `json:"power,omitempty"`
	Price5               *float64      `json:"price5,omitempty"`
	Price22              *float64      `json:"price22,omitempty"`
	Description          *string       `json:"description,omitempty"`
	Brand                *string       `json:"brand,omitempty"`
	Machine              *string       `json:"machine,omitempty"`
	Length               *float64      `json:"length,omitempty"`
	Width                *float64      `json:"width,omitempty"`
	HeightTs             *float64      `json:"heightTs,omitempty"`
	WidthWithSupports    *float64      `json:"widthWithSupports,omitempty"`
	Mass                 *float64      `json:"mass,omitempty"`
	CradleWidthFolded    *float64      `json:"cradleWidthFolded,omitempty"`
	CradleWidthExtended  *float64      `json:"cradleWidthExtended,omitempty"`
	CradleLengthFolded   *float64      `json:"cradleLengthFolded,omitempty"`
	CradleLengthExtended *float64      `json:"cradleLengthExtended,omitempty"`
	ImgArray             []string      `json:"imgArray"`
	MainImageIndex       int           `json:"mainImageIndex"`
	Special              *string       `json:"special,omitempty"`
	RostechReg           bool          `json:"rostechReg"`
	Status               VehicleStatus `json:"status"`
	CreatedAt            time.Time     `json:"createdAt"`
	UpdatedAt            time.Time     `json:"updatedAt"`
}

// VehicleListResponse represents the response for a list of vehicles
type VehicleListResponse struct {
	Vehicles []VehicleResponse `json:"vehicles"`
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PerPage  int              `json:"perPage"`
}

// ToResponse converts a Vehicle to VehicleResponse
func (v *Vehicle) ToResponse() VehicleResponse {
	return VehicleResponse{
		ID:                   v.ID,
		GarageNumber:         v.GarageNumber,
		VIN:                  v.VIN,
		Height:               v.Height,
		Type:                 v.Type,
		Power:                v.Power,
		Price5:               v.Price5,
		Price22:              v.Price22,
		Description:          v.Description,
		Brand:                v.Brand,
		Machine:              v.Machine,
		Length:               v.Length,
		Width:                v.Width,
		HeightTs:             v.HeightTs,
		WidthWithSupports:    v.WidthWithSupports,
		Mass:                 v.Mass,
		CradleWidthFolded:    v.CradleWidthFolded,
		CradleWidthExtended:  v.CradleWidthExtended,
		CradleLengthFolded:   v.CradleLengthFolded,
		CradleLengthExtended: v.CradleLengthExtended,
		ImgArray:             v.ImgArray,
		MainImageIndex:       v.MainImageIndex,
		Special:              v.Special,
		RostechReg:           v.RostechReg,
		Status:               v.Status,
		CreatedAt:            v.CreatedAt,
		UpdatedAt:            v.UpdatedAt,
	}
}

// ToVehicle converts CreateVehicleRequest to Vehicle
func (r *CreateVehicleRequest) ToVehicle() *Vehicle {
	status := VehicleStatusActive
	if r.Status != nil {
		status = *r.Status
	}

	rostechReg := false
	if r.RostechReg != nil {
		rostechReg = *r.RostechReg
	}

	return &Vehicle{
		ID:                   uuid.New(),
		GarageNumber:         r.GarageNumber,
		VIN:                  r.VIN,
		Height:               r.Height,
		Type:                 r.Type,
		Power:                r.Power,
		Price5:               r.Price5,
		Price22:              r.Price22,
		Description:          r.Description,
		Brand:                r.Brand,
		Machine:              r.Machine,
		Length:               r.Length,
		Width:                r.Width,
		HeightTs:             r.HeightTs,
		WidthWithSupports:    r.WidthWithSupports,
		Mass:                 r.Mass,
		CradleWidthFolded:    r.CradleWidthFolded,
		CradleWidthExtended:  r.CradleWidthExtended,
		CradleLengthFolded:   r.CradleLengthFolded,
		CradleLengthExtended: r.CradleLengthExtended,
		ImgArray:             StringArray{},
		MainImageIndex:       1,
		Special:              r.Special,
		RostechReg:           rostechReg,
		Status:               status,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
}

// UpdateVehicle updates the vehicle fields from UpdateVehicleRequest
func (v *Vehicle) UpdateVehicle(req *UpdateVehicleRequest) {
	if req.GarageNumber != nil {
		v.GarageNumber = *req.GarageNumber
	}
	if req.VIN != nil {
		v.VIN = *req.VIN
	}
	if req.Height != nil {
		v.Height = *req.Height
	}
	if req.Type != nil {
		v.Type = req.Type
	}
	if req.Power != nil {
		v.Power = req.Power
	}
	if req.Price5 != nil {
		v.Price5 = req.Price5
	}
	if req.Price22 != nil {
		v.Price22 = req.Price22
	}
	if req.Description != nil {
		v.Description = req.Description
	}
	if req.Brand != nil {
		v.Brand = req.Brand
	}
	if req.Machine != nil {
		v.Machine = req.Machine
	}
	if req.Length != nil {
		v.Length = req.Length
	}
	if req.Width != nil {
		v.Width = req.Width
	}
	if req.HeightTs != nil {
		v.HeightTs = req.HeightTs
	}
	if req.WidthWithSupports != nil {
		v.WidthWithSupports = req.WidthWithSupports
	}
	if req.Mass != nil {
		v.Mass = req.Mass
	}
	if req.CradleWidthFolded != nil {
		v.CradleWidthFolded = req.CradleWidthFolded
	}
	if req.CradleWidthExtended != nil {
		v.CradleWidthExtended = req.CradleWidthExtended
	}
	if req.CradleLengthFolded != nil {
		v.CradleLengthFolded = req.CradleLengthFolded
	}
	if req.CradleLengthExtended != nil {
		v.CradleLengthExtended = req.CradleLengthExtended
	}
	if req.Special != nil {
		v.Special = req.Special
	}
	if req.RostechReg != nil {
		v.RostechReg = *req.RostechReg
	}
	if req.Status != nil {
		v.Status = *req.Status
	}
	if req.MainImageIndex != nil {
		v.MainImageIndex = *req.MainImageIndex
	}
	v.UpdatedAt = time.Now()
}
