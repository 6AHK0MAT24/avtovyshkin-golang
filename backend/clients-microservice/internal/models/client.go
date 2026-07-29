package models

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/google/uuid"
)

// ClientType represents the type of client
type ClientType string

const (
	ClientTypeIndividual  ClientType = "individual"
	ClientTypeLegalEntity ClientType = "legal_entity"
)

// Value implements the driver.Valuer interface for ClientType
func (ct ClientType) Value() (driver.Value, error) {
	return string(ct), nil
}

// Scan implements the sql.Scanner interface for ClientType
func (ct *ClientType) Scan(value interface{}) error {
	if value == nil {
		*ct = ClientTypeIndividual
		return nil
	}
	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return errors.New("invalid type for ClientType")
	}
	*ct = ClientType(str)
	return nil
}

// IsValid checks if the client type is valid
func (ct ClientType) IsValid() bool {
	switch ct {
	case ClientTypeIndividual, ClientTypeLegalEntity:
		return true
	default:
		return false
	}
}

// ClientStatus represents the status of a client
type ClientStatus string

const (
	ClientStatusActive   ClientStatus = "active"
	ClientStatusInactive ClientStatus = "inactive"
	ClientStatusBlocked  ClientStatus = "blocked"
)

// Value implements the driver.Valuer interface for ClientStatus
func (cs ClientStatus) Value() (driver.Value, error) {
	return string(cs), nil
}

// Scan implements the sql.Scanner interface for ClientStatus
func (cs *ClientStatus) Scan(value interface{}) error {
	if value == nil {
		*cs = ClientStatusActive
		return nil
	}
	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return errors.New("invalid type for ClientStatus")
	}
	*cs = ClientStatus(str)
	return nil
}

// IsValid checks if the client status is valid
func (cs ClientStatus) IsValid() bool {
	switch cs {
	case ClientStatusActive, ClientStatusInactive, ClientStatusBlocked:
		return true
	default:
		return false
	}
}

// Client represents the client entity in the database
type Client struct {
	ID uuid.UUID `json:"id" db:"fldId"`

	// Общие поля
	ClientType ClientType   `json:"clientType" db:"fldClientType"`
	Phone      string       `json:"phone" db:"fldPhone"`
	Email      *string      `json:"email,omitempty" db:"fldEmail"`
	Address    *string      `json:"address,omitempty" db:"fldAddress"`
	Status     ClientStatus `json:"status" db:"fldStatus"`
	Notes      *string      `json:"notes,omitempty" db:"fldNotes"`

	// Поля для физических лиц
	FirstName         *string    `json:"firstName,omitempty" db:"fldFirstName"`
	LastName          *string    `json:"lastName,omitempty" db:"fldLastName"`
	MiddleName        *string    `json:"middleName,omitempty" db:"fldMiddleName"`
	BirthDate         *time.Time `json:"birthDate,omitempty" db:"fldBirthDate"`
	PassportSeries    *string    `json:"passportSeries,omitempty" db:"fldPassportSeries"`
	PassportNumber    *string    `json:"passportNumber,omitempty" db:"fldPassportNumber"`
	PassportIssueDate *time.Time `json:"passportIssueDate,omitempty" db:"fldPassportIssueDate"`
	PassportIssuedBy  *string    `json:"passportIssuedBy,omitempty" db:"fldPassportIssuedBy"`
	INN               *string    `json:"inn,omitempty" db:"fldINN"`

	// Поля для юридических лиц
	CompanyName          *string `json:"companyName,omitempty" db:"fldCompanyName"`
	CompanyLegalName     *string `json:"companyLegalName,omitempty" db:"fldCompanyLegalName"`
	OGRN                 *string `json:"ogrn,omitempty" db:"fldOGRN"`
	INNLegal             *string `json:"innLegal,omitempty" db:"fldINNLegal"`
	KPP                  *string `json:"kpp,omitempty" db:"fldKPP"`
	LegalAddress         *string `json:"legalAddress,omitempty" db:"fldLegalAddress"`
	ActualAddress        *string `json:"actualAddress,omitempty" db:"fldActualAddress"`
	BankName             *string `json:"bankName,omitempty" db:"fldBankName"`
	BIC                  *string `json:"bic,omitempty" db:"fldBIC"`
	AccountNumber        *string `json:"accountNumber,omitempty" db:"fldAccountNumber"`
	CorrespondentAccount *string `json:"correspondentAccount,omitempty" db:"fldCorrespondentAccount"`
	DirectorName         *string `json:"directorName,omitempty" db:"fldDirectorName"`
	DirectorPosition     *string `json:"directorPosition,omitempty" db:"fldDirectorPosition"`
	ContactPerson        *string `json:"contactPerson,omitempty" db:"fldContactPerson"`
	ContactPersonPhone   *string `json:"contactPersonPhone,omitempty" db:"fldContactPersonPhone"`
	ContactPersonEmail   *string `json:"contactPersonEmail,omitempty" db:"fldContactPersonEmail"`

	// Метаданные
	CreatedAt time.Time `json:"createdAt" db:"fldCreatedAt"`
	UpdatedAt time.Time `json:"updatedAt" db:"fldUpdatedAt"`
}

// CreateClientRequest represents the request to create a new client
type CreateClientRequest struct {
	ClientType ClientType `json:"clientType" validate:"required,oneof=individual legal_entity"`

	// Общие поля
	Phone   string        `json:"phone" validate:"required"`
	Email   *string       `json:"email,omitempty" validate:"omitempty,email"`
	Address *string       `json:"address,omitempty"`
	Status  *ClientStatus `json:"status,omitempty" validate:"omitempty,oneof=active inactive blocked"`
	Notes   *string       `json:"notes,omitempty"`

	// Поля для физических лиц
	FirstName         *string    `json:"firstName,omitempty" validate:"omitempty,max=100"`
	LastName          *string    `json:"lastName,omitempty" validate:"omitempty,max=100"`
	MiddleName        *string    `json:"middleName,omitempty" validate:"omitempty,max=100"`
	BirthDate         *time.Time `json:"birthDate,omitempty"`
	PassportSeries    *string    `json:"passportSeries,omitempty" validate:"omitempty,len=4"`
	PassportNumber    *string    `json:"passportNumber,omitempty" validate:"omitempty,len=6"`
	PassportIssueDate *time.Time `json:"passportIssueDate,omitempty"`
	PassportIssuedBy  *string    `json:"passportIssuedBy,omitempty" validate:"omitempty,max=255"`
	INN               *string    `json:"inn,omitempty" validate:"omitempty,len=12"`

	// Поля для юридических лиц
	CompanyName          *string `json:"companyName,omitempty" validate:"omitempty,max=255"`
	CompanyLegalName     *string `json:"companyLegalName,omitempty" validate:"omitempty,max=255"`
	OGRN                 *string `json:"ogrn,omitempty" validate:"omitempty,len=15"`
	INNLegal             *string `json:"innLegal,omitempty" validate:"omitempty,len=10"`
	KPP                  *string `json:"kpp,omitempty" validate:"omitempty,len=9"`
	LegalAddress         *string `json:"legalAddress,omitempty"`
	ActualAddress        *string `json:"actualAddress,omitempty"`
	BankName             *string `json:"bankName,omitempty" validate:"omitempty,max=255"`
	BIC                  *string `json:"bic,omitempty" validate:"omitempty,len=9"`
	AccountNumber        *string `json:"accountNumber,omitempty" validate:"omitempty,len=20"`
	CorrespondentAccount *string `json:"correspondentAccount,omitempty" validate:"omitempty,len=20"`
	DirectorName         *string `json:"directorName,omitempty" validate:"omitempty,max=255"`
	DirectorPosition     *string `json:"directorPosition,omitempty" validate:"omitempty,max=100"`
	ContactPerson        *string `json:"contactPerson,omitempty" validate:"omitempty,max=255"`
	ContactPersonPhone   *string `json:"contactPersonPhone,omitempty" validate:"omitempty,max=20"`
	ContactPersonEmail   *string `json:"contactPersonEmail,omitempty" validate:"omitempty,email"`
}

// UpdateClientRequest represents the request to update a client
type UpdateClientRequest struct {
	// Общие поля
	Phone   *string       `json:"phone,omitempty"`
	Email   *string       `json:"email,omitempty" validate:"omitempty,email"`
	Address *string       `json:"address,omitempty"`
	Status  *ClientStatus `json:"status,omitempty" validate:"omitempty,oneof=active inactive blocked"`
	Notes   *string       `json:"notes,omitempty"`

	// Поля для физических лиц
	FirstName         *string    `json:"firstName,omitempty" validate:"omitempty,max=100"`
	LastName          *string    `json:"lastName,omitempty" validate:"omitempty,max=100"`
	MiddleName        *string    `json:"middleName,omitempty" validate:"omitempty,max=100"`
	BirthDate         *time.Time `json:"birthDate,omitempty"`
	PassportSeries    *string    `json:"passportSeries,omitempty" validate:"omitempty,len=4"`
	PassportNumber    *string    `json:"passportNumber,omitempty" validate:"omitempty,len=6"`
	PassportIssueDate *time.Time `json:"passportIssueDate,omitempty"`
	PassportIssuedBy  *string    `json:"passportIssuedBy,omitempty" validate:"omitempty,max=255"`
	INN               *string    `json:"inn,omitempty" validate:"omitempty,len=12"`

	// Поля для юридических лиц
	CompanyName          *string `json:"companyName,omitempty" validate:"omitempty,max=255"`
	CompanyLegalName     *string `json:"companyLegalName,omitempty" validate:"omitempty,max=255"`
	OGRN                 *string `json:"ogrn,omitempty" validate:"omitempty,len=15"`
	INNLegal             *string `json:"innLegal,omitempty" validate:"omitempty,len=10"`
	KPP                  *string `json:"kpp,omitempty" validate:"omitempty,len=9"`
	LegalAddress         *string `json:"legalAddress,omitempty"`
	ActualAddress        *string `json:"actualAddress,omitempty"`
	BankName             *string `json:"bankName,omitempty" validate:"omitempty,max=255"`
	BIC                  *string `json:"bic,omitempty" validate:"omitempty,len=9"`
	AccountNumber        *string `json:"accountNumber,omitempty" validate:"omitempty,len=20"`
	CorrespondentAccount *string `json:"correspondentAccount,omitempty" validate:"omitempty,len=20"`
	DirectorName         *string `json:"directorName,omitempty" validate:"omitempty,max=255"`
	DirectorPosition     *string `json:"directorPosition,omitempty" validate:"omitempty,max=100"`
	ContactPerson        *string `json:"contactPerson,omitempty" validate:"omitempty,max=255"`
	ContactPersonPhone   *string `json:"contactPersonPhone,omitempty" validate:"omitempty,max=20"`
	ContactPersonEmail   *string `json:"contactPersonEmail,omitempty" validate:"omitempty,email"`
}

// ClientListResponse represents the response for listing clients
type ClientListResponse struct {
	Clients []Client `json:"clients"`
	Total   int      `json:"total"`
	Page    int      `json:"page"`
	PerPage int      `json:"perPage"`
}

// SearchFilters represents the search filters for clients
type SearchFilters struct {
	Query      *string       `json:"query,omitempty"`
	ClientType *ClientType   `json:"clientType,omitempty"`
	Status     *ClientStatus `json:"status,omitempty"`
	Page       int           `json:"page"`
	PerPage    int           `json:"perPage"`
}
