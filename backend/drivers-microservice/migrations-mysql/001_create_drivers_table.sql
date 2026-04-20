-- Migration: 001_create_drivers_table.sql
-- Description: Create drivers table with all necessary fields, indexes and triggers (MySQL version)

CREATE TABLE IF NOT EXISTS drivers (
    fldId CHAR(36) PRIMARY KEY,
    fldFirstName VARCHAR(100) NOT NULL,
    fldLastName VARCHAR(100) NOT NULL,
    fldMiddleName VARCHAR(100),
    fldPhone VARCHAR(20) UNIQUE NOT NULL,
    fldEmail VARCHAR(255) UNIQUE,
    fldBirthDate DATE,
    fldDriverLicenseNumber VARCHAR(50) UNIQUE NOT NULL,
    fldDriverLicenseIssueDate DATE NOT NULL,
    fldDriverLicenseExpiryDate DATE NOT NULL,
    fldDriverLicensePhoto TEXT,
    fldDriverLicenseScan TEXT,
    fldPassportSeries VARCHAR(4),
    fldPassportNumber VARCHAR(6),
    fldPassportIssueDate DATE,
    fldPassportPhoto TEXT,
    fldPassportScan TEXT,
    fldAddress TEXT,
    fldExperienceYears INT DEFAULT 0,
    fldStatus VARCHAR(20) DEFAULT 'active',
    fldCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    -- Check constraints for MySQL 8.0+
    CONSTRAINT chk_experience_years CHECK (fldExperienceYears >= 0),
    CONSTRAINT chk_status CHECK (fldStatus IN ('active', 'inactive', 'blocked'))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create indexes for better query performance
CREATE INDEX idx_drivers_phone ON drivers(fldPhone);
CREATE INDEX idx_drivers_email ON drivers(fldEmail);
CREATE INDEX idx_drivers_license ON drivers(fldDriverLicenseNumber);
CREATE INDEX idx_drivers_status ON drivers(fldStatus);
CREATE INDEX idx_drivers_created_at ON drivers(fldCreatedAt);
CREATE INDEX idx_drivers_last_name ON drivers(fldLastName);
CREATE INDEX idx_drivers_first_name ON drivers(fldFirstName);

-- Add comments for documentation
ALTER TABLE drivers COMMENT = 'Table storing driver information';
ALTER TABLE drivers MODIFY COLUMN fldId CHAR(36) COMMENT 'Unique identifier for the driver';
ALTER TABLE drivers MODIFY COLUMN fldFirstName VARCHAR(100) COMMENT 'Driver first name';
ALTER TABLE drivers MODIFY COLUMN fldLastName VARCHAR(100) COMMENT 'Driver last name';
ALTER TABLE drivers MODIFY COLUMN fldMiddleName VARCHAR(100) COMMENT 'Driver middle name (patronymic)';
ALTER TABLE drivers MODIFY COLUMN fldPhone VARCHAR(20) COMMENT 'Driver phone number (unique)';
ALTER TABLE drivers MODIFY COLUMN fldEmail VARCHAR(255) COMMENT 'Driver email address (unique)';
ALTER TABLE drivers MODIFY COLUMN fldBirthDate DATE COMMENT 'Driver date of birth';
ALTER TABLE drivers MODIFY COLUMN fldDriverLicenseNumber VARCHAR(50) COMMENT 'Driver license number (unique)';
ALTER TABLE drivers MODIFY COLUMN fldDriverLicenseIssueDate DATE COMMENT 'Date when driver license was issued';
ALTER TABLE drivers MODIFY COLUMN fldDriverLicenseExpiryDate DATE COMMENT 'Date when driver license expires';
ALTER TABLE drivers MODIFY COLUMN fldDriverLicensePhoto TEXT COMMENT 'Path to driver license photo file';
ALTER TABLE drivers MODIFY COLUMN fldDriverLicenseScan TEXT COMMENT 'Path to driver license scan file';
ALTER TABLE drivers MODIFY COLUMN fldPassportSeries VARCHAR(4) COMMENT 'Passport series (4 digits)';
ALTER TABLE drivers MODIFY COLUMN fldPassportNumber VARCHAR(6) COMMENT 'Passport number (6 digits)';
ALTER TABLE drivers MODIFY COLUMN fldPassportIssueDate DATE COMMENT 'Date when passport was issued';
ALTER TABLE drivers MODIFY COLUMN fldPassportPhoto TEXT COMMENT 'Path to passport photo file';
ALTER TABLE drivers MODIFY COLUMN fldPassportScan TEXT COMMENT 'Path to passport scan file';
ALTER TABLE drivers MODIFY COLUMN fldAddress TEXT COMMENT 'Driver residential address';
ALTER TABLE drivers MODIFY COLUMN fldExperienceYears INT COMMENT 'Years of driving experience';
ALTER TABLE drivers MODIFY COLUMN fldStatus VARCHAR(20) COMMENT 'Driver status: active, inactive, or blocked';
ALTER TABLE drivers MODIFY COLUMN fldCreatedAt TIMESTAMP COMMENT 'Timestamp when record was created';
ALTER TABLE drivers MODIFY COLUMN fldUpdatedAt TIMESTAMP COMMENT 'Timestamp when record was last updated';

-- Create trigger for UUID generation (similar to PostgreSQL gen_random_uuid())
DELIMITER //
CREATE TRIGGER before_drivers_insert
BEFORE INSERT ON drivers
FOR EACH ROW
BEGIN
    IF NEW.fldId IS NULL OR NEW.fldId = '' THEN
        SET NEW.fldId = UUID();
    END IF;
END//
DELIMITER ;
