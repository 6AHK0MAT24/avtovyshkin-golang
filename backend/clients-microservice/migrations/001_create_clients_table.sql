-- Migration: 001_create_clients_table.sql
-- Description: Create clients table for both individuals and legal entities

CREATE TABLE IF NOT EXISTS clients (
    fldId UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fldClientType VARCHAR(20) NOT NULL CHECK (fldClientType IN ('individual', 'legal_entity')),
    
    -- Общие поля для всех клиентов
    fldPhone VARCHAR(20) NOT NULL,
    fldEmail VARCHAR(255),
    fldAddress TEXT,
    fldStatus VARCHAR(20) DEFAULT 'active' CHECK (fldStatus IN ('active', 'inactive', 'blocked')),
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
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание индексов для оптимизации запросов
CREATE INDEX IF NOT EXISTS idx_clients_type ON clients(fldClientType);
CREATE INDEX IF NOT EXISTS idx_clients_phone ON clients(fldPhone);
CREATE INDEX IF NOT EXISTS idx_clients_email ON clients(fldEmail);
CREATE INDEX IF NOT EXISTS idx_clients_status ON clients(fldStatus);
CREATE INDEX IF NOT EXISTS idx_clients_created_at ON clients(fldCreatedAt);
CREATE INDEX IF NOT EXISTS idx_clients_last_name ON clients(fldLastName);
CREATE INDEX IF NOT EXISTS idx_clients_first_name ON clients(fldFirstName);
CREATE INDEX IF NOT EXISTS idx_clients_company_name ON clients(fldCompanyName);
CREATE INDEX IF NOT EXISTS idx_clients_inn_individual ON clients(fldINN);
CREATE INDEX IF NOT EXISTS idx_clients_inn_legal ON clients(fldINNLegal);
CREATE INDEX IF NOT EXISTS idx_clients_ogrn ON clients(fldOGRN);

-- Создание триггера для автоматического обновления fldUpdatedAt
DROP TRIGGER IF EXISTS update_clients_updated_at ON clients;
CREATE TRIGGER update_clients_updated_at
    BEFORE UPDATE ON clients
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Добавление комментариев для документации
COMMENT ON TABLE clients IS 'Таблица хранения информации о клиентах (физические и юридические лица)';
COMMENT ON COLUMN clients.fldId IS 'Уникальный идентификатор клиента';
COMMENT ON COLUMN clients.fldClientType IS 'Тип клиента: individual (физическое лицо) или legal_entity (юридическое лицо)';
COMMENT ON COLUMN clients.fldPhone IS 'Номер телефона клиента';
COMMENT ON COLUMN clients.fldEmail IS 'Email адрес клиента';
COMMENT ON COLUMN clients.fldAddress IS 'Адрес клиента';
COMMENT ON COLUMN clients.fldStatus IS 'Статус клиента: active, inactive или blocked';
COMMENT ON COLUMN clients.fldNotes IS 'Дополнительные заметки о клиенте';

-- Комментарии для полей физических лиц
COMMENT ON COLUMN clients.fldFirstName IS 'Имя физического лица';
COMMENT ON COLUMN clients.fldLastName IS 'Фамилия физического лица';
COMMENT ON COLUMN clients.fldMiddleName IS 'Отчество физического лица';
COMMENT ON COLUMN clients.fldBirthDate IS 'Дата рождения физического лица';
COMMENT ON COLUMN clients.fldPassportSeries IS 'Серия паспорта (4 цифры)';
COMMENT ON COLUMN clients.fldPassportNumber IS 'Номер паспорта (6 цифр)';
COMMENT ON COLUMN clients.fldPassportIssueDate IS 'Дата выдачи паспорта';
COMMENT ON COLUMN clients.fldPassportIssuedBy IS 'Кем выдан паспорт';
COMMENT ON COLUMN clients.fldINN IS 'ИНН физического лица (12 цифр)';

-- Комментарии для полей юридических лиц
COMMENT ON COLUMN clients.fldCompanyName IS 'Краткое название компании';
COMMENT ON COLUMN clients.fldCompanyLegalName IS 'Полное юридическое название компании';
COMMENT ON COLUMN clients.fldOGRN IS 'ОГРН юридического лица (15 цифр)';
COMMENT ON COLUMN clients.fldINNLegal IS 'ИНН юридического лица (10 цифр)';
COMMENT ON COLUMN clients.fldKPP IS 'КПП юридического лица (9 цифр)';
COMMENT ON COLUMN clients.fldLegalAddress IS 'Юридический адрес компании';
COMMENT ON COLUMN clients.fldActualAddress IS 'Фактический адрес компании';
COMMENT ON COLUMN clients.fldBankName IS 'Название банка';
COMMENT ON COLUMN clients.fldBIC IS 'БИК банка (9 цифр)';
COMMENT ON COLUMN clients.fldAccountNumber IS 'Расчетный счет (20 цифр)';
COMMENT ON COLUMN clients.fldCorrespondentAccount IS 'Корреспондентский счет (20 цифр)';
COMMENT ON COLUMN clients.fldDirectorName IS 'ФИО директора';
COMMENT ON COLUMN clients.fldDirectorPosition IS 'Должность директора';
COMMENT ON COLUMN clients.fldContactPerson IS 'Контактное лицо';
COMMENT ON COLUMN clients.fldContactPersonPhone IS 'Телефон контактного лица';
COMMENT ON COLUMN clients.fldContactPersonEmail IS 'Email контактного лица';
COMMENT ON COLUMN clients.fldCreatedAt IS 'Время создания записи';
COMMENT ON COLUMN clients.fldUpdatedAt IS 'Время последнего обновления записи';
