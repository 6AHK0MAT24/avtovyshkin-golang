package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"drivers-service/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDriverRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	driver := &models.Driver{
		ID:                       uuid.New(),
		FirstName:                "Иван",
		LastName:                 "Иванов",
		MiddleName:               "Иванович",
		Phone:                    "+79001234567",
		Email:                    "ivan@example.com",
		BirthDate:                sql.NullTime{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		DriverLicenseNumber:      "1234567890",
		DriverLicenseIssueDate:   sql.NullTime{Time: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		DriverLicenseExpiryDate:  sql.NullTime{Time: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		DriverLicensePhoto:       "/uploads/licenses/photo_001.jpg",
		DriverLicenseScan:        "/uploads/licenses/scan_001.jpg",
		PassportSeries:           "1234",
		PassportNumber:           "567890",
		PassportIssueDate:        sql.NullTime{Time: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		PassportPhoto:            "/uploads/passports/photo_001.jpg",
		PassportScan:             "/uploads/passports/scan_001.jpg",
		Address:                  "г. Москва, ул. Тестовая, д. 1",
		ExperienceYears:          10,
		Status:                   models.StatusActive,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	mock.ExpectExec("INSERT INTO drivers").
		WithArgs(
			driver.ID, driver.FirstName, driver.LastName, driver.MiddleName, driver.Phone, driver.Email,
			driver.BirthDate, driver.DriverLicenseNumber, driver.DriverLicenseIssueDate,
			driver.DriverLicenseExpiryDate, driver.DriverLicensePhoto, driver.DriverLicenseScan,
			driver.PassportSeries, driver.PassportNumber, driver.PassportIssueDate,
			driver.PassportPhoto, driver.PassportScan, driver.Address, driver.ExperienceYears,
			driver.Status, driver.CreatedAt, driver.UpdatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(context.Background(), driver)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	driverID := uuid.New()
	expectedDriver := &models.Driver{
		ID:                       driverID,
		FirstName:                "Иван",
		LastName:                 "Иванов",
		MiddleName:               "Иванович",
		Phone:                    "+79001234567",
		Email:                    "ivan@example.com",
		BirthDate:                sql.NullTime{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		DriverLicenseNumber:      "1234567890",
		DriverLicenseIssueDate:   sql.NullTime{Time: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		DriverLicenseExpiryDate:  sql.NullTime{Time: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		DriverLicensePhoto:       "/uploads/licenses/photo_001.jpg",
		DriverLicenseScan:        "/uploads/licenses/scan_001.jpg",
		PassportSeries:           "1234",
		PassportNumber:           "567890",
		PassportIssueDate:        sql.NullTime{Time: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		PassportPhoto:            "/uploads/passports/photo_001.jpg",
		PassportScan:             "/uploads/passports/scan_001.jpg",
		Address:                  "г. Москва, ул. Тестовая, д. 1",
		ExperienceYears:          10,
		Status:                   models.StatusActive,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	rows := sqlmock.NewRows([]string{
		"fldId", "fldFirstName", "fldLastName", "fldMiddleName", "fldPhone", "fldEmail",
		"fldBirthDate", "fldDriverLicenseNumber", "fldDriverLicenseIssueDate",
		"fldDriverLicenseExpiryDate", "fldDriverLicensePhoto", "fldDriverLicenseScan",
		"fldPassportSeries", "fldPassportNumber", "fldPassportIssueDate",
		"fldPassportPhoto", "fldPassportScan", "fldAddress", "fldExperienceYears",
		"fldStatus", "fldCreatedAt", "fldUpdatedAt",
	}).AddRow(
		expectedDriver.ID, expectedDriver.FirstName, expectedDriver.LastName, expectedDriver.MiddleName,
		expectedDriver.Phone, expectedDriver.Email, expectedDriver.BirthDate,
		expectedDriver.DriverLicenseNumber, expectedDriver.DriverLicenseIssueDate,
		expectedDriver.DriverLicenseExpiryDate, expectedDriver.DriverLicensePhoto, expectedDriver.DriverLicenseScan,
		expectedDriver.PassportSeries, expectedDriver.PassportNumber, expectedDriver.PassportIssueDate,
		expectedDriver.PassportPhoto, expectedDriver.PassportScan, expectedDriver.Address,
		expectedDriver.ExperienceYears, expectedDriver.Status, expectedDriver.CreatedAt, expectedDriver.UpdatedAt,
	)

	mock.ExpectQuery("SELECT (.+) FROM drivers WHERE fldId = \\$1").
		WithArgs(driverID).
		WillReturnRows(rows)

	driver, err := repo.GetByID(context.Background(), driverID)
	assert.NoError(t, err)
	assert.Equal(t, expectedDriver.ID, driver.ID)
	assert.Equal(t, expectedDriver.FirstName, driver.FirstName)
	assert.Equal(t, expectedDriver.LastName, driver.LastName)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_GetByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	driverID := uuid.New()

	mock.ExpectQuery("SELECT (.+) FROM drivers WHERE fldId = \\$1").
		WithArgs(driverID).
		WillReturnError(sql.ErrNoRows)

	driver, err := repo.GetByID(context.Background(), driverID)
	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Contains(t, err.Error(), "driver not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	driverID1 := uuid.New()
	driverID2 := uuid.New()

	countRows := sqlmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM drivers").
		WillReturnRows(countRows)

	driverRows := sqlmock.NewRows([]string{
		"fldId", "fldFirstName", "fldLastName", "fldMiddleName", "fldPhone", "fldEmail",
		"fldBirthDate", "fldDriverLicenseNumber", "fldDriverLicenseIssueDate",
		"fldDriverLicenseExpiryDate", "fldDriverLicensePhoto", "fldDriverLicenseScan",
		"fldPassportSeries", "fldPassportNumber", "fldPassportIssueDate",
		"fldPassportPhoto", "fldPassportScan", "fldAddress", "fldExperienceYears",
		"fldStatus", "fldCreatedAt", "fldUpdatedAt",
	}).
		AddRow(
			driverID1, "Иван", "Иванов", "Иванович", "+79001234567", "ivan@example.com",
			sql.NullTime{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
			"1234567890", sql.NullTime{Time: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
			sql.NullTime{Time: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
			"/uploads/licenses/photo_001.jpg", "/uploads/licenses/scan_001.jpg",
			"1234", "567890", sql.NullTime{Time: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
			"/uploads/passports/photo_001.jpg", "/uploads/passports/scan_001.jpg",
			"г. Москва, ул. Тестовая, д. 1", 10, models.StatusActive, time.Now(), time.Now(),
		).
		AddRow(
			driverID2, "Петр", "Петров", "Петрович", "+79009876543", "petr@example.com",
			sql.NullTime{Time: time.Date(1985, 5, 15, 0, 0, 0, 0, time.UTC), Valid: true},
			"0987654321", sql.NullTime{Time: time.Date(2005, 5, 15, 0, 0, 0, 0, time.UTC), Valid: true},
			sql.NullTime{Time: time.Date(2025, 5, 15, 0, 0, 0, 0, time.UTC), Valid: true},
			"/uploads/licenses/photo_002.jpg", "/uploads/licenses/scan_002.jpg",
			"4321", "098765", sql.NullTime{Time: time.Date(2003, 5, 15, 0, 0, 0, 0, time.UTC), Valid: true},
			"/uploads/passports/photo_002.jpg", "/uploads/passports/scan_002.jpg",
			"г. Санкт-Петербург, ул. Тестовая, д. 2", 15, models.StatusActive, time.Now(), time.Now(),
		)

	mock.ExpectQuery("SELECT (.+) FROM drivers ORDER BY fldCreatedAt DESC LIMIT \\$1 OFFSET \\$2").
		WithArgs(10, 0).
		WillReturnRows(driverRows)

	drivers, total, err := repo.GetAll(context.Background(), 10, 0)
	assert.NoError(t, err)
	assert.Len(t, drivers, 2)
	assert.Equal(t, int64(2), total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	driver := &models.Driver{
		ID:                       uuid.New(),
		FirstName:                "Иван",
		LastName:                 "Иванов",
		MiddleName:               "Иванович",
		Phone:                    "+79001234567",
		Email:                    "ivan@example.com",
		BirthDate:                sql.NullTime{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		DriverLicenseNumber:      "1234567890",
		DriverLicenseIssueDate:   sql.NullTime{Time: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		DriverLicenseExpiryDate:  sql.NullTime{Time: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		DriverLicensePhoto:       "/uploads/licenses/photo_001.jpg",
		DriverLicenseScan:        "/uploads/licenses/scan_001.jpg",
		PassportSeries:           "1234",
		PassportNumber:           "567890",
		PassportIssueDate:        sql.NullTime{Time: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		PassportPhoto:            "/uploads/passports/photo_001.jpg",
		PassportScan:             "/uploads/passports/scan_001.jpg",
		Address:                  "г. Москва, ул. Тестовая, д. 1",
		ExperienceYears:          10,
		Status:                   models.StatusActive,
		UpdatedAt:                time.Now(),
	}

	mock.ExpectExec("UPDATE drivers SET").
		WithArgs(
			driver.FirstName, driver.LastName, driver.MiddleName, driver.Phone, driver.Email,
			driver.BirthDate, driver.DriverLicenseNumber, driver.DriverLicenseIssueDate,
			driver.DriverLicenseExpiryDate, driver.DriverLicensePhoto, driver.DriverLicenseScan,
			driver.PassportSeries, driver.PassportNumber, driver.PassportIssueDate,
			driver.PassportPhoto, driver.PassportScan, driver.Address, driver.ExperienceYears,
			driver.Status, driver.UpdatedAt, driver.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(context.Background(), driver)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_Update_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	driver := &models.Driver{
		ID:        uuid.New(),
		FirstName: "Иван",
		LastName:  "Иванов",
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec("UPDATE drivers SET").
		WithArgs(
			driver.FirstName, driver.LastName, driver.MiddleName, driver.Phone, driver.Email,
			driver.BirthDate, driver.DriverLicenseNumber, driver.DriverLicenseIssueDate,
			driver.DriverLicenseExpiryDate, driver.DriverLicensePhoto, driver.DriverLicenseScan,
			driver.PassportSeries, driver.PassportNumber, driver.PassportIssueDate,
			driver.PassportPhoto, driver.PassportScan, driver.Address, driver.ExperienceYears,
			driver.Status, driver.UpdatedAt, driver.ID,
		).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.Update(context.Background(), driver)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "driver not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	driverID := uuid.New()

	mock.ExpectExec("DELETE FROM drivers WHERE fldId = \\$1").
		WithArgs(driverID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(context.Background(), driverID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_Delete_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	driverID := uuid.New()

	mock.ExpectExec("DELETE FROM drivers WHERE fldId = \\$1").
		WithArgs(driverID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.Delete(context.Background(), driverID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "driver not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_GetByPhone(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	phone := "+79001234567"
	driverID := uuid.New()

	rows := sqlmock.NewRows([]string{
		"fldId", "fldFirstName", "fldLastName", "fldMiddleName", "fldPhone", "fldEmail",
		"fldBirthDate", "fldDriverLicenseNumber", "fldDriverLicenseIssueDate",
		"fldDriverLicenseExpiryDate", "fldDriverLicensePhoto", "fldDriverLicenseScan",
		"fldPassportSeries", "fldPassportNumber", "fldPassportIssueDate",
		"fldPassportPhoto", "fldPassportScan", "fldAddress", "fldExperienceYears",
		"fldStatus", "fldCreatedAt", "fldUpdatedAt",
	}).AddRow(
		driverID, "Иван", "Иванов", "Иванович", phone, "ivan@example.com",
		sql.NullTime{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		"1234567890", sql.NullTime{Time: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		sql.NullTime{Time: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		"/uploads/licenses/photo_001.jpg", "/uploads/licenses/scan_001.jpg",
		"1234", "567890", sql.NullTime{Time: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		"/uploads/passports/photo_001.jpg", "/uploads/passports/scan_001.jpg",
		"г. Москва, ул. Тестовая, д. 1", 10, models.StatusActive, time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT (.+) FROM drivers WHERE fldPhone = \\$1").
		WithArgs(phone).
		WillReturnRows(rows)

	driver, err := repo.GetByPhone(context.Background(), phone)
	assert.NoError(t, err)
	assert.Equal(t, phone, driver.Phone)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_GetByLicenseNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	licenseNumber := "1234567890"
	driverID := uuid.New()

	rows := sqlmock.NewRows([]string{
		"fldId", "fldFirstName", "fldLastName", "fldMiddleName", "fldPhone", "fldEmail",
		"fldBirthDate", "fldDriverLicenseNumber", "fldDriverLicenseIssueDate",
		"fldDriverLicenseExpiryDate", "fldDriverLicensePhoto", "fldDriverLicenseScan",
		"fldPassportSeries", "fldPassportNumber", "fldPassportIssueDate",
		"fldPassportPhoto", "fldPassportScan", "fldAddress", "fldExperienceYears",
		"fldStatus", "fldCreatedAt", "fldUpdatedAt",
	}).AddRow(
		driverID, "Иван", "Иванов", "Иванович", "+79001234567", "ivan@example.com",
		sql.NullTime{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		licenseNumber, sql.NullTime{Time: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		sql.NullTime{Time: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		"/uploads/licenses/photo_001.jpg", "/uploads/licenses/scan_001.jpg",
		"1234", "567890", sql.NullTime{Time: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		"/uploads/passports/photo_001.jpg", "/uploads/passports/scan_001.jpg",
		"г. Москва, ул. Тестовая, д. 1", 10, models.StatusActive, time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT (.+) FROM drivers WHERE fldDriverLicenseNumber = \\$1").
		WithArgs(licenseNumber).
		WillReturnRows(rows)

	driver, err := repo.GetByLicenseNumber(context.Background(), licenseNumber)
	assert.NoError(t, err)
	assert.Equal(t, licenseNumber, driver.DriverLicenseNumber)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDriverRepository_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &driverRepository{db: db}

	email := "ivan@example.com"
	driverID := uuid.New()

	rows := sqlmock.NewRows([]string{
		"fldId", "fldFirstName", "fldLastName", "fldMiddleName", "fldPhone", "fldEmail",
		"fldBirthDate", "fldDriverLicenseNumber", "fldDriverLicenseIssueDate",
		"fldDriverLicenseExpiryDate", "fldDriverLicensePhoto", "fldDriverLicenseScan",
		"fldPassportSeries", "fldPassportNumber", "fldPassportIssueDate",
		"fldPassportPhoto", "fldPassportScan", "fldAddress", "fldExperienceYears",
		"fldStatus", "fldCreatedAt", "fldUpdatedAt",
	}).AddRow(
		driverID, "Иван", "Иванов", "Иванович", "+79001234567", email,
		sql.NullTime{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		"1234567890", sql.NullTime{Time: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		sql.NullTime{Time: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		"/uploads/licenses/photo_001.jpg", "/uploads/licenses/scan_001.jpg",
		"1234", "567890", sql.NullTime{Time: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
		"/uploads/passports/photo_001.jpg", "/uploads/passports/scan_001.jpg",
		"г. Москва, ул. Тестовая, д. 1", 10, models.StatusActive, time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT (.+) FROM drivers WHERE fldEmail = \\$1").
		WithArgs(email).
		WillReturnRows(rows)

	driver, err := repo.GetByEmail(context.Background(), email)
	assert.NoError(t, err)
	assert.Equal(t, email, driver.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}
