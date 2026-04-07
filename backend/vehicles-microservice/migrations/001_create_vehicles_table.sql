-- Создание таблицы vehicles
CREATE TABLE IF NOT EXISTS vehicles (
    fldId UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fldGarageNumber VARCHAR(50) NOT NULL UNIQUE,
    fldVIN VARCHAR(17) NOT NULL UNIQUE,
    fldHeight DECIMAL(10, 2) NOT NULL,
    fldType VARCHAR(100),
    fldPower INTEGER,
    fldPrice5 DECIMAL(10, 2),
    fldPrice22 DECIMAL(10, 2),
    fldDescription TEXT,
    fldBrand VARCHAR(200),
    fldMachine VARCHAR(200),
    -- Габариты ТС в транспортном положении
    fldLength DECIMAL(10, 2),
    fldWidth DECIMAL(10, 2),
    fldHeightTs DECIMAL(10, 2),
    fldWidthWithSupports DECIMAL(10, 2),
    fldMass DECIMAL(10, 2),
    -- Габариты люльки
    fldCradleWidthFolded DECIMAL(10, 2),
    fldCradleWidthExtended DECIMAL(10, 2),
    fldCradleLengthFolded DECIMAL(10, 2),
    fldCradleLengthExtended DECIMAL(10, 2),
    -- Изображения
    fldImgArray JSONB DEFAULT '[]'::jsonb,
    fldMainImageIndex INTEGER DEFAULT 1,
    -- Дополнительные поля
    fldSpecial TEXT,
    fldRostechReg BOOLEAN DEFAULT false,
    fldStatus VARCHAR(20) DEFAULT 'active',
    fldCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для оптимизации поиска
CREATE INDEX idx_vehicles_garage_number ON vehicles(fldGarageNumber);
CREATE INDEX idx_vehicles_vin ON vehicles(fldVIN);
CREATE INDEX idx_vehicles_height ON vehicles(fldHeight);
CREATE INDEX idx_vehicles_status ON vehicles(fldStatus);
CREATE INDEX idx_vehicles_brand_height ON vehicles(fldBrand, fldHeight);

-- Триггер для автоматического обновления fldUpdatedAt
CREATE OR REPLACE FUNCTION update_vehicles_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.fldUpdatedAt = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_vehicles_updated_at
    BEFORE UPDATE ON vehicles
    FOR EACH ROW
    EXECUTE FUNCTION update_vehicles_updated_at();
