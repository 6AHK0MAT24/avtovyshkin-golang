-- Migration: 001_create_clients_table.sql
-- Description: Create clients table for both individuals and legal entities (MySQL version)

CREATE TABLE IF NOT EXISTS clients (
    fldId CHAR(36) PRIMARY KEY,
    fldClientType VARCHAR(20) NOT NULL,
    
    -- Общие поля для всех клиентов
    fldPhone VARCHAR(20) NOT NULL,
    fldEmail VARCHAR(255),
    fldAddress TEXT,
    fldStatus VARCHAR(20) DEFAULT 'active',
    fldNotes TEXT,
    
    -- Поля для физических лиц
    fldFirstName VARCHAR(100),
    fldLastName VARCHAR(100),
    fldMiddleName VARCHAR(100),
    fldBirthDate DATE,
    fldPassportSeries VARCHAR(4),
    fldPassportNumber VARCHAR(6),
    fldPassportIssueDate DATE,
    fldPassportIssuedBy VARCHAR(255),
    fldINN VARCHAR(12),
    
    -- Поля для юридических лиц
    fldCompanyName VARCHAR(255),
    fldCompanyLegalName VARCHAR(255),
    fldOGRN VARCHAR(15),
    fldINNLegal VARCHAR(10),
    fldKPP VARCHAR(9),
    fldLegalAddress TEXT,
    fldActualAddress TEXT,
    fldBankName VARCHAR(255),
    fldBIC VARCHAR(9),
    fldAccountNumber VARCHAR(20),
    fldCorrespondentAccount VARCHAR(20),
    fldDirectorName VARCHAR(255),
    fldDirectorPosition VARCHAR(100),
    fldContactPerson VARCHAR(255),
    fldContactPersonPhone VARCHAR(20),
    fldContactPersonEmail VARCHAR(255),
    
    -- Метаданные
    fldCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    -- Constraints
    CONSTRAINT chk_client_type CHECK (fldClientType IN ('individual', 'legal_entity')),
    CONSTRAINT chk_client_status CHECK (fldStatus IN ('active', 'inactive', 'blocked'))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Создание индексов для оптимизации запросов
CREATE INDEX idx_clients_type ON clients(fldClientType);
CREATE INDEX idx_clients_phone ON clients(fldPhone);
CREATE INDEX idx_clients_email ON clients(fldEmail);
CREATE INDEX idx_clients_status ON clients(fldStatus);
CREATE INDEX idx_clients_created_at ON clients(fldCreatedAt);
CREATE INDEX idx_clients_last_name ON clients(fldLastName);
CREATE INDEX idx_clients_first_name ON clients(fldFirstName);
CREATE INDEX idx_clients_company_name ON clients(fldCompanyName);
CREATE INDEX idx_clients_inn_individual ON clients(fldINN);
CREATE INDEX idx_clients_inn_legal ON clients(fldINNLegal);
CREATE INDEX idx_clients_ogrn ON clients(fldOGRN);
