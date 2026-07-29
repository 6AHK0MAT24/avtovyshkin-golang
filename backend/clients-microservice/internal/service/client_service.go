package service

import (
	"context"
	"fmt"

	"clients-service/internal/models"
	"clients-service/internal/repository"

	"github.com/google/uuid"
)

// ClientService handles business logic for clients
type ClientService struct {
	repo *repository.ClientRepository
}

// NewClientService creates a new client service
func NewClientService(repo *repository.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

// Create creates a new client
func (s *ClientService) Create(ctx context.Context, req *models.CreateClientRequest) (*models.Client, error) {
	// Validate client type
	if !req.ClientType.IsValid() {
		return nil, fmt.Errorf("invalid client type: %s", req.ClientType)
	}

	// Validate fields based on client type
	if req.ClientType == models.ClientTypeIndividual {
		// Validate individual fields
		if req.FirstName == nil || *req.FirstName == "" {
			return nil, fmt.Errorf("first name is required for individual clients")
		}
		if req.LastName == nil || *req.LastName == "" {
			return nil, fmt.Errorf("last name is required for individual clients")
		}
	} else if req.ClientType == models.ClientTypeLegalEntity {
		// Validate legal entity fields
		if req.CompanyName == nil || *req.CompanyName == "" {
			return nil, fmt.Errorf("company name is required for legal entity clients")
		}
		if req.INNLegal == nil || *req.INNLegal == "" {
			return nil, fmt.Errorf("INN is required for legal entity clients")
		}
	}

	// Set default status if not provided
	status := models.ClientStatusActive
	if req.Status != nil {
		status = *req.Status
	}

	client := &models.Client{
		ID:         uuid.New(),
		ClientType: req.ClientType,
		Phone:      req.Phone,
		Email:      req.Email,
		Address:    req.Address,
		Status:     status,
		Notes:      req.Notes,

		// Individual fields
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		MiddleName:        req.MiddleName,
		BirthDate:         req.BirthDate,
		PassportSeries:    req.PassportSeries,
		PassportNumber:    req.PassportNumber,
		PassportIssueDate: req.PassportIssueDate,
		PassportIssuedBy:  req.PassportIssuedBy,
		INN:               req.INN,

		// Legal entity fields
		CompanyName:          req.CompanyName,
		CompanyLegalName:     req.CompanyLegalName,
		OGRN:                 req.OGRN,
		INNLegal:             req.INNLegal,
		KPP:                  req.KPP,
		LegalAddress:         req.LegalAddress,
		ActualAddress:        req.ActualAddress,
		BankName:             req.BankName,
		BIC:                  req.BIC,
		AccountNumber:        req.AccountNumber,
		CorrespondentAccount: req.CorrespondentAccount,
		DirectorName:         req.DirectorName,
		DirectorPosition:     req.DirectorPosition,
		ContactPerson:        req.ContactPerson,
		ContactPersonPhone:   req.ContactPersonPhone,
		ContactPersonEmail:   req.ContactPersonEmail,
	}

	if err := s.repo.Create(ctx, client); err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return client, nil
}

// GetByID retrieves a client by ID
func (s *ClientService) GetByID(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	client, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %w", err)
	}
	return client, nil
}

// GetAll retrieves all clients with pagination
func (s *ClientService) GetAll(ctx context.Context, page, perPage int) (*models.ClientListResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	clients, total, err := s.repo.GetAll(ctx, page, perPage)
	if err != nil {
		return nil, fmt.Errorf("failed to get clients: %w", err)
	}

	return &models.ClientListResponse{
		Clients: clients,
		Total:   total,
		Page:    page,
		PerPage: perPage,
	}, nil
}

// GetByType retrieves clients by type with pagination
func (s *ClientService) GetByType(ctx context.Context, clientType models.ClientType, page, perPage int) (*models.ClientListResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	clients, total, err := s.repo.GetByType(ctx, clientType, page, perPage)
	if err != nil {
		return nil, fmt.Errorf("failed to get clients by type: %w", err)
	}

	return &models.ClientListResponse{
		Clients: clients,
		Total:   total,
		Page:    page,
		PerPage: perPage,
	}, nil
}

// Update updates a client
func (s *ClientService) Update(ctx context.Context, id uuid.UUID, req *models.UpdateClientRequest) (*models.Client, error) {
	// Get existing client
	client, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %w", err)
	}

	// Update fields if provided
	if req.Phone != nil {
		client.Phone = *req.Phone
	}
	if req.Email != nil {
		client.Email = req.Email
	}
	if req.Address != nil {
		client.Address = req.Address
	}
	if req.Status != nil {
		client.Status = *req.Status
	}
	if req.Notes != nil {
		client.Notes = req.Notes
	}

	// Update individual fields if provided
	if req.FirstName != nil {
		client.FirstName = req.FirstName
	}
	if req.LastName != nil {
		client.LastName = req.LastName
	}
	if req.MiddleName != nil {
		client.MiddleName = req.MiddleName
	}
	if req.BirthDate != nil {
		client.BirthDate = req.BirthDate
	}
	if req.PassportSeries != nil {
		client.PassportSeries = req.PassportSeries
	}
	if req.PassportNumber != nil {
		client.PassportNumber = req.PassportNumber
	}
	if req.PassportIssueDate != nil {
		client.PassportIssueDate = req.PassportIssueDate
	}
	if req.PassportIssuedBy != nil {
		client.PassportIssuedBy = req.PassportIssuedBy
	}
	if req.INN != nil {
		client.INN = req.INN
	}

	// Update legal entity fields if provided
	if req.CompanyName != nil {
		client.CompanyName = req.CompanyName
	}
	if req.CompanyLegalName != nil {
		client.CompanyLegalName = req.CompanyLegalName
	}
	if req.OGRN != nil {
		client.OGRN = req.OGRN
	}
	if req.INNLegal != nil {
		client.INNLegal = req.INNLegal
	}
	if req.KPP != nil {
		client.KPP = req.KPP
	}
	if req.LegalAddress != nil {
		client.LegalAddress = req.LegalAddress
	}
	if req.ActualAddress != nil {
		client.ActualAddress = req.ActualAddress
	}
	if req.BankName != nil {
		client.BankName = req.BankName
	}
	if req.BIC != nil {
		client.BIC = req.BIC
	}
	if req.AccountNumber != nil {
		client.AccountNumber = req.AccountNumber
	}
	if req.CorrespondentAccount != nil {
		client.CorrespondentAccount = req.CorrespondentAccount
	}
	if req.DirectorName != nil {
		client.DirectorName = req.DirectorName
	}
	if req.DirectorPosition != nil {
		client.DirectorPosition = req.DirectorPosition
	}
	if req.ContactPerson != nil {
		client.ContactPerson = req.ContactPerson
	}
	if req.ContactPersonPhone != nil {
		client.ContactPersonPhone = req.ContactPersonPhone
	}
	if req.ContactPersonEmail != nil {
		client.ContactPersonEmail = req.ContactPersonEmail
	}

	if err := s.repo.Update(ctx, client); err != nil {
		return nil, fmt.Errorf("failed to update client: %w", err)
	}

	return client, nil
}

// Delete deletes a client
func (s *ClientService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete client: %w", err)
	}
	return nil
}

// Search searches clients by filters
func (s *ClientService) Search(ctx context.Context, filters *models.SearchFilters) (*models.ClientListResponse, error) {
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.PerPage < 1 || filters.PerPage > 100 {
		filters.PerPage = 10
	}

	clients, total, err := s.repo.Search(ctx, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to search clients: %w", err)
	}

	return &models.ClientListResponse{
		Clients: clients,
		Total:   total,
		Page:    filters.Page,
		PerPage: filters.PerPage,
	}, nil
}
