-- Migration: 001_create_drivers_table.sql
-- Description: Create drivers table with all necessary fields, indexes and triggers

CREATE TABLE IF NOT EXISTS drivers (
    fldId UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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
    fldExperienceYears INTEGER DEFAULT 0 CHECK (fldExperienceYears >= 0),
    fldStatus VARCHAR(20) DEFAULT 'active' CHECK (fldStatus IN ('active', 'inactive', 'blocked')),
    fldCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_drivers_phone ON drivers(fldPhone);
CREATE INDEX IF NOT EXISTS idx_drivers_email ON drivers(fldEmail);
CREATE INDEX IF NOT EXISTS idx_drivers_license ON drivers(fldDriverLicenseNumber);
CREATE INDEX IF NOT EXISTS idx_drivers_status ON drivers(fldStatus);
CREATE INDEX IF NOT EXISTS idx_drivers_created_at ON drivers(fldCreatedAt);
CREATE INDEX IF NOT EXISTS idx_drivers_last_name ON drivers(fldLastName);
CREATE INDEX IF NOT EXISTS idx_drivers_first_name ON drivers(fldFirstName);

-- Create trigger to automatically update fldUpdatedAt
DROP TRIGGER IF EXISTS update_drivers_updated_at ON drivers;
CREATE TRIGGER update_drivers_updated_at
    BEFORE UPDATE ON drivers
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE drivers IS 'Table storing driver information';
COMMENT ON COLUMN drivers.fldId IS 'Unique identifier for the driver';
COMMENT ON COLUMN drivers.fldFirstName IS 'Driver first name';
COMMENT ON COLUMN drivers.fldLastName IS 'Driver last name';
COMMENT ON COLUMN drivers.fldMiddleName IS 'Driver middle name (patronymic)';
COMMENT ON COLUMN drivers.fldPhone IS 'Driver phone number (unique)';
COMMENT ON COLUMN drivers.fldEmail IS 'Driver email address (unique)';
COMMENT ON COLUMN drivers.fldBirthDate IS 'Driver date of birth';
COMMENT ON COLUMN drivers.fldDriverLicenseNumber IS 'Driver license number (unique)';
COMMENT ON COLUMN drivers.fldDriverLicenseIssueDate IS 'Date when driver license was issued';
COMMENT ON COLUMN drivers.fldDriverLicenseExpiryDate IS 'Date when driver license expires';
COMMENT ON COLUMN drivers.fldDriverLicensePhoto IS 'Path to driver license photo file';
COMMENT ON COLUMN drivers.fldDriverLicenseScan IS 'Path to driver license scan file';
COMMENT ON COLUMN drivers.fldPassportSeries IS 'Passport series (4 digits)';
COMMENT ON COLUMN drivers.fldPassportNumber IS 'Passport number (6 digits)';
COMMENT ON COLUMN drivers.fldPassportIssueDate IS 'Date when passport was issued';
COMMENT ON COLUMN drivers.fldPassportPhoto IS 'Path to passport photo file';
COMMENT ON COLUMN drivers.fldPassportScan IS 'Path to passport scan file';
COMMENT ON COLUMN drivers.fldAddress IS 'Driver residential address';
COMMENT ON COLUMN drivers.fldExperienceYears IS 'Years of driving experience';
COMMENT ON COLUMN drivers.fldStatus IS 'Driver status: active, inactive, or blocked';
COMMENT ON COLUMN drivers.fldCreatedAt IS 'Timestamp when record was created';
COMMENT ON COLUMN drivers.fldUpdatedAt IS 'Timestamp when record was last updated';
