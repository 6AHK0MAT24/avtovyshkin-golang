# 🚀 Быстрый старт MySQL на REG.RU

## ✅ Решение проблем с кодировкой и дубликатами

### 📋 Проблемы и решения:

1. **Проблема:** Нечитабельные символы в SQL файлах (╨Я╨╡╤В╤А)
   - **Решение:** Используйте файлы `sql/insert_drivers_clean.sql` и `sql/insert_vehicles_clean.sql`

2. **Проблема:** Ошибка `#1062 - Duplicate entry '' for key 'vehicles.PRIMARY'`
   - **Решение:** Файлы `insert_*_clean.sql` содержат команду `TRUNCATE TABLE` для очистки таблиц перед вставкой

## 🎯 Пошаговая инструкция:

### 1. Подключение к phpMyAdmin:
- **URL:** https://localhost:1500/
- **Пользователь:** `u3424187_root_avtovyshkin`
- **Пароль:** `JavaScript6315`
- **База данных:** `u3424187_avtovyshkin`

### 2. Создание таблиц:

#### Таблица drivers:
```sql
CREATE TABLE drivers (
    fldid VARCHAR(50) PRIMARY KEY,
    fldfirstname VARCHAR(100),
    fldlastname VARCHAR(100),
    fldmiddlename VARCHAR(100),
    fldphone VARCHAR(20),
    fldemail VARCHAR(100),
    fldbirthday DATE,
    fldphoto TEXT,
    flddriverlicensenumber VARCHAR(50),
    flddriverlicenseissuedate DATE,
    flddriverlicenseexpirydate DATE,
    flddriverlicensephoto TEXT,
    flddriverlicensescan TEXT,
    fldpassportseries VARCHAR(10),
    fldpassportnumber VARCHAR(20),
    fldpassportissuedate DATE,
    fldpassportphoto TEXT,
    fldpassportscan TEXT,
    fldaddress TEXT,
    fldexperience INT,
    fldstatus VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

#### Таблица vehicles:
```sql
CREATE TABLE vehicles (
    fldid INT AUTO_INCREMENT PRIMARY KEY,
    fldGarageNumber VARCHAR(50) UNIQUE,
    fldVIN VARCHAR(50),
    fldHeight INT,
    fldType VARCHAR(100),
    fldPower INT,
    fldPrice5 DECIMAL(10,2),
    fldPrice22 DECIMAL(10,2),
    fldDescription TEXT,
    fldBrand VARCHAR(100),
    fldMachine VARCHAR(100),
    fldLength DECIMAL(10,2),
    fldWidth DECIMAL(10,2),
    fldHeightTs DECIMAL(10,2),
    fldWidthWithSupports DECIMAL(10,2),
    fldMass DECIMAL(10,2),
    fldCradleWidthFolded DECIMAL(10,2),
    fldCradleWidthExtended DECIMAL(10,2),
    fldCradleLengthFolded DECIMAL(10,2),
    fldCradleLengthExtended DECIMAL(10,2),
    fldImgArray TEXT,
    fldMainImageIndex INT,
    fldSpecial TEXT,
    fldRostechReg BOOLEAN DEFAULT FALSE,
    fldStatus VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 3. Вставка данных:

#### Вставка водителей:
```sql
-- Выполните содержимое файла sql/insert_drivers_clean.sql
-- Или скопируйте и вставьте весь SQL код из этого файла
```

#### Вставка автовышек:
```sql
-- Выполните содержимое файла sql/insert_vehicles_clean.sql
-- Или скопируйте и вставьте весь SQL код из этого файла
```

### 4. Проверка данных:

```sql
-- Проверка водителей
SELECT COUNT(*) as total_drivers FROM drivers;
SELECT * FROM drivers LIMIT 5;

-- Проверка автовышек
SELECT COUNT(*) as total_vehicles FROM vehicles;
SELECT * FROM vehicles LIMIT 5;
```

## 📦 Содержимое данных:

### Drivers (10 записей):
- d001: Ivan Ivanov (10 лет опыта)
- d002: Petr Petrov (8 лет опыта)
- d003: Sergey Sidorov (6 лет опыта)
- d004: Alexander Kozlov (9 лет опыта)
- d005: Dmitry Morozov (4 года опыта)
- d006: Nikolay Volkov (7 лет опыта)
- d007: Andrey Sokolov (5 лет опыта)
- d008: Maxim Lebedev (11 лет опыта)
- d009: Alexey Kuznetsov (3 года опыта)
- d010: Igor Popov (12 лет опыта)

### Vehicles (8 автовышек):
- А001: 17м, Телескопическая, 800 кг
- А002: 20м, Телескопическая, 800 кг
- А003: 22м, Телескопическая, 800 кг
- А004: 26м, Телескопическая, 300 кг
- А005: 30м, Телескопическая, 300 кг
- А006: 34м, Телескоп + колено, 300 кг
- А007: 37м, Телескоп + колено, 300 кг
- А008: 45м, Телескоп + стрела и рукоять, 450 кг

## 🚀 Запуск приложения:

```powershell
# Настройка production конфигурации
cd backend/drivers-microservice
Copy-Item .env.production .env -Force

cd ../vehicles-microservice
Copy-Item .env.production .env -Force

cd ../frontend
Copy-Item .env.production .env -Force

# Запуск микросервисов
cd ../drivers-microservice
go run cmd/server/main.go

# В другом терминале
cd vehicles-microservice
go run cmd/server/main.go

# В третьем терминале
cd frontend
yarn dev
```

## ✅ Преимущества чистых SQL файлов:

1. **Правильная кодировка UTF-8** - все данные читабельные
2. **Автоматическая очистка** - команда TRUNCATE предотвращает дубликаты
3. **Полная структура** - все поля заполнены корректными данными
4. **Готовые к использованию** - можно сразу запускать приложение

## 🎉 Результат:

После выполнения этих шагов у вас будет полностью рабочая база данных MySQL на REG.RU с:
- ✅ 10 водителями с полными данными
- ✅ 8 автовышками с характеристиками
- ✅ Правильной кодировкой UTF-8
- ✅ Без дубликатов и ошибок

**Проект готов к работе на хостинге REG.RU!** 🚀
