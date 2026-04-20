-- Создание таблицы drivers для REG.RU MySQL
CREATE TABLE IF NOT EXISTS drivers (
    fldId CHAR(36) PRIMARY KEY,
    fldFirstName VARCHAR(100) NOT NULL,
    fldLastName VARCHAR(100) NOT NULL,
    fldMiddleName VARCHAR(100),
    fldPhone VARCHAR(20) UNIQUE NOT NULL,
    fldEmail VARCHAR(255) UNIQUE,
    fldBirthDate DATE,
    fldPhoto TEXT,
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
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Индексы для оптимизации
CREATE INDEX idx_drivers_phone ON drivers(fldPhone);
CREATE INDEX idx_drivers_email ON drivers(fldEmail);
CREATE INDEX idx_drivers_license ON drivers(fldDriverLicenseNumber);
CREATE INDEX idx_drivers_status ON drivers(fldStatus);
CREATE INDEX idx_drivers_created_at ON drivers(fldCreatedAt);
CREATE INDEX idx_drivers_last_name ON drivers(fldLastName);
CREATE INDEX idx_drivers_first_name ON drivers(fldFirstName);
