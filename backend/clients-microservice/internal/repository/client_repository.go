package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"clients-service/internal/config"
	"clients-service/internal/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// ClientRepository handles database operations for clients
type ClientRepository struct {
	db *sql.DB
}

// NewClientRepository creates a new client repository
func NewClientRepository(cfg *config.DatabaseConfig) (*ClientRepository, error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &ClientRepository{db: db}, nil
}

// Close closes the database connection
func (r *ClientRepository) Close() error {
	return r.db.Close()
}

// Create creates a new client
func (r *ClientRepository) Create(ctx context.Context, client *models.Client) error {
	// Generate UUID if not provided
	if client.ID == uuid.Nil {
		client.ID = uuid.New()
	}

	query := `
		INSERT INTO clients (
			fldId, fldClientType, fldPhone, fldEmail, fldAddress, fldStatus, fldNotes,
			fldFirstName, fldLastName, fldMiddleName, fldBirthDate,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate, fldPassportIssuedBy, fldINN,
			fldCompanyName, fldCompanyLegalName, fldOGRN, fldINNLegal, fldKPP,
			fldLegalAddress, fldActualAddress, fldBankName, fldBIC,
			fldAccountNumber, fldCorrespondentAccount, fldDirectorName, fldDirectorPosition,
			fldContactPerson, fldContactPersonPhone, fldContactPersonEmail
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query,
		client.ID, client.ClientType, client.Phone, client.Email, client.Address, client.Status, client.Notes,
		client.FirstName, client.LastName, client.MiddleName, client.BirthDate,
		client.PassportSeries, client.PassportNumber, client.PassportIssueDate, client.PassportIssuedBy, client.INN,
		client.CompanyName, client.CompanyLegalName, client.OGRN, client.INNLegal, client.KPP,
		client.LegalAddress, client.ActualAddress, client.BankName, client.BIC,
		client.AccountNumber, client.CorrespondentAccount, client.DirectorName, client.DirectorPosition,
		client.ContactPerson, client.ContactPersonPhone, client.ContactPersonEmail,
	)

	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	// Get the created timestamps
	err = r.db.QueryRowContext(ctx, "SELECT fldCreatedAt, fldUpdatedAt FROM clients WHERE fldId = ?", client.ID).
		Scan(&client.CreatedAt, &client.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to get created timestamps: %w", err)
	}

	return nil
}

// GetByID retrieves a client by ID
func (r *ClientRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	query := `
		SELECT
			fldId, fldClientType, fldPhone, fldEmail, fldAddress, fldStatus, fldNotes,
			fldFirstName, fldLastName, fldMiddleName, fldBirthDate,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate, fldPassportIssuedBy, fldINN,
			fldCompanyName, fldCompanyLegalName, fldOGRN, fldINNLegal, fldKPP,
			fldLegalAddress, fldActualAddress, fldBankName, fldBIC,
			fldAccountNumber, fldCorrespondentAccount, fldDirectorName, fldDirectorPosition,
			fldContactPerson, fldContactPersonPhone, fldContactPersonEmail,
			fldCreatedAt, fldUpdatedAt
		FROM clients
		WHERE fldId = ?
	`

	client := &models.Client{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&client.ID, &client.ClientType, &client.Phone, &client.Email, &client.Address, &client.Status, &client.Notes,
		&client.FirstName, &client.LastName, &client.MiddleName, &client.BirthDate,
		&client.PassportSeries, &client.PassportNumber, &client.PassportIssueDate, &client.PassportIssuedBy, &client.INN,
		&client.CompanyName, &client.CompanyLegalName, &client.OGRN, &client.INNLegal, &client.KPP,
		&client.LegalAddress, &client.ActualAddress, &client.BankName, &client.BIC,
		&client.AccountNumber, &client.CorrespondentAccount, &client.DirectorName, &client.DirectorPosition,
		&client.ContactPerson, &client.ContactPersonPhone, &client.ContactPersonEmail,
		&client.CreatedAt, &client.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("client not found")
		}
		return nil, fmt.Errorf("failed to get client: %w", err)
	}

	return client, nil
}

// GetAll retrieves all clients with pagination
func (r *ClientRepository) GetAll(ctx context.Context, page, perPage int) ([]models.Client, int, error) {
	offset := (page - 1) * perPage

	// Get total count
	var total int
	countQuery := "SELECT COUNT(*) FROM clients"
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count clients: %w", err)
	}

	// Get clients
	query := `
		SELECT
			fldId, fldClientType, fldPhone, fldEmail, fldAddress, fldStatus, fldNotes,
			fldFirstName, fldLastName, fldMiddleName, fldBirthDate,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate, fldPassportIssuedBy, fldINN,
			fldCompanyName, fldCompanyLegalName, fldOGRN, fldINNLegal, fldKPP,
			fldLegalAddress, fldActualAddress, fldBankName, fldBIC,
			fldAccountNumber, fldCorrespondentAccount, fldDirectorName, fldDirectorPosition,
			fldContactPerson, fldContactPersonPhone, fldContactPersonEmail,
			fldCreatedAt, fldUpdatedAt
		FROM clients
		ORDER BY fldCreatedAt DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, perPage, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get clients: %w", err)
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		err := rows.Scan(
			&client.ID, &client.ClientType, &client.Phone, &client.Email, &client.Address, &client.Status, &client.Notes,
			&client.FirstName, &client.LastName, &client.MiddleName, &client.BirthDate,
			&client.PassportSeries, &client.PassportNumber, &client.PassportIssueDate, &client.PassportIssuedBy, &client.INN,
			&client.CompanyName, &client.CompanyLegalName, &client.OGRN, &client.INNLegal, &client.KPP,
			&client.LegalAddress, &client.ActualAddress, &client.BankName, &client.BIC,
			&client.AccountNumber, &client.CorrespondentAccount, &client.DirectorName, &client.DirectorPosition,
			&client.ContactPerson, &client.ContactPersonPhone, &client.ContactPersonEmail,
			&client.CreatedAt, &client.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan client: %w", err)
		}
		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating clients: %w", err)
	}

	return clients, total, nil
}

// Update updates a client
func (r *ClientRepository) Update(ctx context.Context, client *models.Client) error {
	query := `
		UPDATE clients SET
			fldPhone = ?, fldEmail = ?, fldAddress = ?, fldStatus = ?, fldNotes = ?,
			fldFirstName = ?, fldLastName = ?, fldMiddleName = ?, fldBirthDate = ?,
			fldPassportSeries = ?, fldPassportNumber = ?, fldPassportIssueDate = ?, fldPassportIssuedBy = ?, fldINN = ?,
			fldCompanyName = ?, fldCompanyLegalName = ?, fldOGRN = ?, fldINNLegal = ?, fldKPP = ?,
			fldLegalAddress = ?, fldActualAddress = ?, fldBankName = ?, fldBIC = ?,
			fldAccountNumber = ?, fldCorrespondentAccount = ?, fldDirectorName = ?, fldDirectorPosition = ?,
			fldContactPerson = ?, fldContactPersonPhone = ?, fldContactPersonEmail = ?,
			fldUpdatedAt = CURRENT_TIMESTAMP
		WHERE fldId = ?
	`

	result, err := r.db.ExecContext(ctx, query,
		client.Phone, client.Email, client.Address, client.Status, client.Notes,
		client.FirstName, client.LastName, client.MiddleName, client.BirthDate,
		client.PassportSeries, client.PassportNumber, client.PassportIssueDate, client.PassportIssuedBy, client.INN,
		client.CompanyName, client.CompanyLegalName, client.OGRN, client.INNLegal, client.KPP,
		client.LegalAddress, client.ActualAddress, client.BankName, client.BIC,
		client.AccountNumber, client.CorrespondentAccount, client.DirectorName, client.DirectorPosition,
		client.ContactPerson, client.ContactPersonPhone, client.ContactPersonEmail,
		client.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update client: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("client not found")
	}

	// Get the updated timestamp
	err = r.db.QueryRowContext(ctx, "SELECT fldUpdatedAt FROM clients WHERE fldId = ?", client.ID).
		Scan(&client.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to get updated timestamp: %w", err)
	}

	return nil
}

// Delete deletes a client
func (r *ClientRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM clients WHERE fldId = ?"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete client: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("client not found")
	}

	return nil
}

// Search searches clients by query
func (r *ClientRepository) Search(ctx context.Context, filters *models.SearchFilters) ([]models.Client, int, error) {
	offset := (filters.Page - 1) * filters.PerPage

	// Build WHERE clause
	var whereClauses []string
	var args []interface{}
	argIndex := 1

	if filters.Query != nil && *filters.Query != "" {
		whereClauses = append(whereClauses, fmt.Sprintf(
			"(fldFirstName LIKE ? OR fldLastName LIKE ? OR fldMiddleName LIKE ? OR fldCompanyName LIKE ? OR fldPhone LIKE ? OR fldEmail LIKE ?)",
		))
		queryPattern := "%" + *filters.Query + "%"
		args = append(args, queryPattern, queryPattern, queryPattern, queryPattern, queryPattern, queryPattern)
		argIndex += 6
	}

	if filters.ClientType != nil {
		whereClauses = append(whereClauses, "fldClientType = ?")
		args = append(args, *filters.ClientType)
		argIndex++
	}

	if filters.Status != nil {
		whereClauses = append(whereClauses, "fldStatus = ?")
		args = append(args, *filters.Status)
		argIndex++
	}

	whereClause := ""
	if len(whereClauses) > 0 {
		whereClause = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Get total count
	var total int
	countQuery := "SELECT COUNT(*) FROM clients " + whereClause
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count clients: %w", err)
	}

	// Get clients
	query := `
		SELECT
			fldId, fldClientType, fldPhone, fldEmail, fldAddress, fldStatus, fldNotes,
			fldFirstName, fldLastName, fldMiddleName, fldBirthDate,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate, fldPassportIssuedBy, fldINN,
			fldCompanyName, fldCompanyLegalName, fldOGRN, fldINNLegal, fldKPP,
			fldLegalAddress, fldActualAddress, fldBankName, fldBIC,
			fldAccountNumber, fldCorrespondentAccount, fldDirectorName, fldDirectorPosition,
			fldContactPerson, fldContactPersonPhone, fldContactPersonEmail,
			fldCreatedAt, fldUpdatedAt
		FROM clients
		` + whereClause + `
		ORDER BY fldCreatedAt DESC
		LIMIT ? OFFSET ?
	`

	args = append(args, filters.PerPage, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to search clients: %w", err)
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		err := rows.Scan(
			&client.ID, &client.ClientType, &client.Phone, &client.Email, &client.Address, &client.Status, &client.Notes,
			&client.FirstName, &client.LastName, &client.MiddleName, &client.BirthDate,
			&client.PassportSeries, &client.PassportNumber, &client.PassportIssueDate, &client.PassportIssuedBy, &client.INN,
			&client.CompanyName, &client.CompanyLegalName, &client.OGRN, &client.INNLegal, &client.KPP,
			&client.LegalAddress, &client.ActualAddress, &client.BankName, &client.BIC,
			&client.AccountNumber, &client.CorrespondentAccount, &client.DirectorName, &client.DirectorPosition,
			&client.ContactPerson, &client.ContactPersonPhone, &client.ContactPersonEmail,
			&client.CreatedAt, &client.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan client: %w", err)
		}
		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating clients: %w", err)
	}

	return clients, total, nil
}

// GetByType retrieves clients by type with pagination
func (r *ClientRepository) GetByType(ctx context.Context, clientType models.ClientType, page, perPage int) ([]models.Client, int, error) {
	offset := (page - 1) * perPage

	// Get total count
	var total int
	countQuery := "SELECT COUNT(*) FROM clients WHERE fldClientType = ?"
	if err := r.db.QueryRowContext(ctx, countQuery, clientType).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count clients: %w", err)
	}

	// Get clients
	query := `
		SELECT
			fldId, fldClientType, fldPhone, fldEmail, fldAddress, fldStatus, fldNotes,
			fldFirstName, fldLastName, fldMiddleName, fldBirthDate,
			fldPassportSeries, fldPassportNumber, fldPassportIssueDate, fldPassportIssuedBy, fldINN,
			fldCompanyName, fldCompanyLegalName, fldOGRN, fldINNLegal, fldKPP,
			fldLegalAddress, fldActualAddress, fldBankName, fldBIC,
			fldAccountNumber, fldCorrespondentAccount, fldDirectorName, fldDirectorPosition,
			fldContactPerson, fldContactPersonPhone, fldContactPersonEmail,
			fldCreatedAt, fldUpdatedAt
		FROM clients
		WHERE fldClientType = ?
		ORDER BY fldCreatedAt DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, clientType, perPage, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get clients: %w", err)
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		err := rows.Scan(
			&client.ID, &client.ClientType, &client.Phone, &client.Email, &client.Address, &client.Status, &client.Notes,
			&client.FirstName, &client.LastName, &client.MiddleName, &client.BirthDate,
			&client.PassportSeries, &client.PassportNumber, &client.PassportIssueDate, &client.PassportIssuedBy, &client.INN,
			&client.CompanyName, &client.CompanyLegalName, &client.OGRN, &client.INNLegal, &client.KPP,
			&client.LegalAddress, &client.ActualAddress, &client.BankName, &client.BIC,
			&client.AccountNumber, &client.CorrespondentAccount, &client.DirectorName, &client.DirectorPosition,
			&client.ContactPerson, &client.ContactPersonPhone, &client.ContactPersonEmail,
			&client.CreatedAt, &client.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan client: %w", err)
		}
		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating clients: %w", err)
	}

	return clients, total, nil
}
