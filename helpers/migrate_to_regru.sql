-- ============================================
-- Миграция базы данных на REG.RU MySQL
-- ============================================

-- Создание таблицы drivers
DROP TABLE IF EXISTS drivers;

CREATE TABLE drivers (
    fldid VARCHAR(36) PRIMARY KEY,
    fldfirstname VARCHAR(100) NOT NULL,
    fldlastname VARCHAR(100) NOT NULL,
    fldpatronymic VARCHAR(100),
    fldphonenumber VARCHAR(20) NOT NULL,
    fldemail VARCHAR(100) NOT NULL,
    fldlicense VARCHAR(50) NOT NULL,
    fldlicenseexpirydate DATE NOT NULL,
    fldstatus VARCHAR(50) DEFAULT 'active',
    fldcreatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fldupdatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    fldphoto VARCHAR(255),
    fldaddress VARCHAR(255),
    fldbirthdate DATE,
    fldpassport VARCHAR(50),
    fldexperience INT DEFAULT 0,
    fldrating DECIMAL(3,2) DEFAULT 5.00,
    fldtotaltrips INT DEFAULT 0,
    fldbalance DECIMAL(10,2) DEFAULT 0.00,
    UNIQUE KEY unique_email (fldemail),
    UNIQUE KEY unique_phone (fldphonenumber),
    UNIQUE KEY unique_license (fldlicense)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Вставка данных водителей
INSERT INTO drivers (fldid, fldfirstname, fldlastname, fldpatronymic, fldphonenumber, fldemail, fldlicense, fldlicenseexpirydate, fldstatus, fldcreatedat, fldupdatedat, fldphoto, fldaddress, fldbirthdate, fldpassport, fldexperience, fldrating, fldtotaltrips, fldbalance) VALUES
('46c5d4b0-4ef6-48f7-82f8-155f0c1803af', 'Иван', 'Иванов', 'Иванович', '+79001234567', 'ivan.ivanov@example.com', '1234567890', '2025-12-31', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/drivers/photo/ivan_ivanov.jpg', 'г. Москва, ул. Ленина, д. 1', '1990-01-15', '4500123456', 5, 4.8, 150, 25000.00),
('436915a7-8f25-4052-badc-d8f7004f4d54', 'Петр', 'Петров', 'Петрович', '+79002345678', 'petr.petrov@example.com', '0987654321', '2026-06-30', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/drivers/photo/petr_petrov.jpg', 'г. Санкт-Петербург, ул. Невского, д. 10', '1985-05-20', '4500987654', 8, 4.9, 320, 45000.00),
('db9423a9-ae22-4aff-b839-5dd55b8c8ff4', 'Сергей', 'Сидоров', 'Сергеевич', '+79003456789', 'sergey.sidorov@example.com', '1122334455', '2025-08-15', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/drivers/photo/sergey_sidorov.jpg', 'г. Казань, ул. Баумана, д. 25', '1988-11-30', '4501122334', 3, 4.7, 85, 18000.00),
('a75ad2e7-acbc-48a8-8dbe-879b77bd8ff2', 'Александр', 'Козлов', 'Александрович', '+79004567890', 'alexander.kozlov@example.com', '5566778899', '2026-03-20', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/drivers/photo/alexander_kozlov.jpg', 'г. Новосибирск, ул. Красный проспект, д. 5', '1992-07-10', '4505566778', 4, 4.6, 120, 22000.00),
('e45661d1-681a-477a-9561-8492500680a2', 'Дмитрий', 'Морозов', 'Дмитриевич', '+79005678901', 'dmitry.morozov@example.com', '9988776655', '2025-10-25', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/drivers/photo/dmitry_morozov.jpg', 'г. Екатеринбург, ул. Мамина-Сибиряка, д. 15', '1987-03-25', '4509988776', 6, 4.8, 200, 35000.00);

-- Создание таблицы vehicles
DROP TABLE IF EXISTS vehicles;

CREATE TABLE vehicles (
    fldid VARCHAR(36) PRIMARY KEY,
    fldbrand VARCHAR(100) NOT NULL,
    fldmodel VARCHAR(100) NOT NULL,
    fldyear INT NOT NULL,
    fldplatenumber VARCHAR(20) NOT NULL,
    fldvin VARCHAR(50) NOT NULL,
    fldcolor VARCHAR(50),
    fldmileage INT DEFAULT 0,
    fldfueltype VARCHAR(50),
    fldtransmissiontype VARCHAR(50),
    fldbodytype VARCHAR(50),
    fldnumberofseats INT DEFAULT 5,
    fldstatus VARCHAR(50) DEFAULT 'active',
    fldcreatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fldupdatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    fldphoto VARCHAR(255),
    flddriverid VARCHAR(36),
    FOREIGN KEY (flddriverid) REFERENCES drivers(fldid) ON DELETE SET NULL,
    UNIQUE KEY unique_plate (fldplatenumber),
    UNIQUE KEY unique_vin (fldvin)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Вставка данных автомобилей
INSERT INTO vehicles (fldid, fldbrand, fldmodel, fldyear, fldplatenumber, fldvin, fldcolor, fldmileage, fldfueltype, fldtransmissiontype, fldbodytype, fldnumberofseats, fldstatus, fldcreatedat, fldupdatedat, fldphoto, flddriverid) VALUES
('1', 'Toyota', 'Camry', '2020', 'А123БВ777', 'JTDKN3DU5A0123456', 'Белый', '45000', 'Бензин', 'Автоматическая', 'Седан', '5', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/vehicles/photo/toyota_camry.jpg', '46c5d4b0-4ef6-48f7-82f8-155f0c1803af'),
('2', 'Volkswagen', 'Polo', '2021', 'В456ГД799', 'WVWZZZAUZPW123456', 'Серый', '32000', 'Бензин', 'Автоматическая', 'Седан', '5', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/vehicles/photo/vw_polo.jpg', '436915a7-8f25-4052-badc-d8f7004f4d54'),
('3', 'Hyundai', 'Solaris', '2019', 'Е789ЖК750', 'KMHCT41BPLU123456', 'Черный', '58000', 'Бензин', 'Автоматическая', 'Седан', '5', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/vehicles/photo/hyundai_solaris.jpg', 'db9423a9-ae22-4aff-b839-5dd55b8c8ff4'),
('4', 'Kia', 'Rio', '2022', 'К012ЛМ777', 'KNADH51ABLA123456', 'Синий', '18000', 'Бензин', 'Автоматическая', 'Седан', '5', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/vehicles/photo/kia_rio.jpg', 'a75ad2e7-acbc-48a8-8dbe-879b77bd8ff2'),
('5', 'Skoda', 'Rapid', '2020', 'М345НО799', 'TMBJH9NE3L0123456', 'Красный', '42000', 'Бензин', 'Автоматическая', 'Седан', '5', 'active', '2026-04-17 07:56:17', '2026-04-17 07:56:17', '/uploads/vehicles/photo/skoda_rapid.jpg', 'e45661d1-681a-477a-9561-8492500680a2');

-- Проверка результатов
SELECT 'Drivers count:' as info, COUNT(*) as count FROM drivers
UNION ALL
SELECT 'Vehicles count:', COUNT(*) FROM vehicles;
