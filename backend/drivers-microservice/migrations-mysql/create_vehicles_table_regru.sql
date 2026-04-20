-- Создание таблицы vehicles для REG.RU MySQL
CREATE TABLE IF NOT EXISTS vehicles (
    fldId CHAR(36) PRIMARY KEY,
    fldGarageNumber VARCHAR(50) NOT NULL UNIQUE,
    fldVIN VARCHAR(17) NOT NULL UNIQUE,
    fldHeight DECIMAL(10, 2) NOT NULL,
    fldType VARCHAR(100),
    fldPower INT,
    fldPrice5 DECIMAL(10, 2),
    fldPrice22 DECIMAL(10, 2),
    fldDescription TEXT,
    fldBrand VARCHAR(200),
    fldMachine VARCHAR(200),
    fldLength DECIMAL(10, 2),
    fldWidth DECIMAL(10, 2),
    fldHeightTs DECIMAL(10, 2),
    fldWidthWithSupports DECIMAL(10, 2),
    fldMass DECIMAL(10, 2),
    fldCradleWidthFolded DECIMAL(10, 2),
    fldCradleWidthExtended DECIMAL(10, 2),
    fldCradleLengthFolded DECIMAL(10, 2),
    fldCradleLengthExtended DECIMAL(10, 2),
    fldImgArray JSON DEFAULT ('[]'),
    fldMainImageIndex INT DEFAULT 1,
    fldSpecial TEXT,
    fldRostechReg BOOLEAN DEFAULT false,
    fldStatus VARCHAR(20) DEFAULT 'active',
    fldCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Индексы для оптимизации
CREATE INDEX idx_vehicles_garage_number ON vehicles(fldGarageNumber);
CREATE INDEX idx_vehicles_vin ON vehicles(fldVIN);
CREATE INDEX idx_vehicles_height ON vehicles(fldHeight);
CREATE INDEX idx_vehicles_status ON vehicles(fldStatus);
CREATE INDEX idx_vehicles_brand_height ON vehicles(fldBrand, fldHeight);
