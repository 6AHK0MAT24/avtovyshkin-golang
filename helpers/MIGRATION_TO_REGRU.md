# Инструкция по миграции данных из PostgreSQL в MySQL на REG.RU

## Подготовка

1. **Убедитесь, что PostgreSQL контейнер запущен:**
   ```powershell
   docker ps | findstr postgres
   ```

2. **Запустите скрипт миграции:**
   ```powershell
   .\migrate-to-regru.ps1
   ```

## Выполнение миграции на REG.RU

### Шаг 1: Подключение к базе данных

1. Откройте phpMyAdmin на REG.RU: https://localhost:1500/
2. Войдите с учетными данными:
   - Пользователь: `u3424187_root_avtovyshkin`
   - Пароль: `JavaScript6315`
   - База данных: `u3424187_avtovyshkin`

### Шаг 2: Создание таблиц

1. В phpMyAdmin выберите базу данных `u3424187_avtovyshkin`
2. Перейдите в раздел "SQL"
3. Выполните содержимое файла `sql/create_drivers_table_regru.sql`
4. Выполните содержимое файла `sql/create_vehicles_table_regru.sql`

### Шаг 3: Импорт данных

**Вариант А: Через phpMyAdmin**

1. Для таблицы `drivers`:
   - Выберите таблицу `drivers`
   - Нажмите "Import" (Импорт)
   - Выберите файл `drivers_export.csv`
   - Установите формат: CSV
   - Разделитель: `,`
   - Обрамление: `"`
   - Нажмите "Go"

2. Для таблицы `vehicles`:
   - Выберите таблицу `vehicles`
   - Нажмите "Import" (Импорт)
   - Выберите файл `vehicles_export.csv`
   - Установите формат: CSV
   - Разделитель: `,`
   - Обрамление: `"`
   - Нажмите "Go"

**Вариант Б: Через SQL команды**

```sql
-- Импорт данных drivers
LOAD DATA LOCAL INFILE 'drivers_export.csv' 
INTO TABLE drivers
FIELDS TERMINATED BY ',' 
ENCLOSED BY '"' 
LINES TERMINATED BY '\n' 
IGNORE 1 ROWS;

-- Импорт данных vehicles
LOAD DATA LOCAL INFILE 'vehicles_export.csv' 
INTO TABLE vehicles
FIELDS TERMINATED BY ',' 
ENCLOSED BY '"' 
LINES TERMINATED BY '\n' 
IGNORE 1 ROWS;
```

### Шаг 4: Проверка данных

```sql
-- Проверка количества записей
SELECT COUNT(*) as drivers_count FROM drivers;
SELECT COUNT(*) as vehicles_count FROM vehicles;

-- Проверка данных drivers
SELECT * FROM drivers LIMIT 5;

-- Проверка данных vehicles
SELECT * FROM vehicles LIMIT 5;
```

## Обновление конфигурации приложения

Конфигурационные файлы уже обновлены для работы с REG.RU:

- `backend/drivers-microservice/.env.production`
- `backend/vehicles-microservice/.env.production`
- `frontend/.env.production`

## Запуск приложения с REG.RU базой данных

### Локальная разработка с REG.RU базой:

```powershell
# Drivers microservice
cd backend/drivers-microservice
cp .env.production .env
go run cmd/server/main.go

# Vehicles microservice
cd ../vehicles-microservice
cp .env.production .env
go run cmd/server/main.go

# Frontend
cd ../../frontend
cp .env.production .env
yarn dev
```

### Production на REG.RU:

1. Загрузите файлы проекта на сервер
2. Установите зависимости:
   ```bash
   cd backend/drivers-microservice
   go mod download
   go build -o drivers-server cmd/server/main.go
   
   cd ../vehicles-microservice
   go mod download
   go build -o vehicles-server cmd/server/main.go
   
   cd ../../frontend
   yarn install
   yarn build
   ```

3. Запустите сервисы:
   ```bash
   # Drivers
   ./backend/drivers-microservice/drivers-server
   
   # Vehicles
   ./backend/vehicles-microservice/vehicles-server
   ```

## Проверка подключения

```powershell
# Тест подключения к MySQL на REG.RU
Test-NetConnection -ComputerName localhost -Port 3306
```

## Решение проблем

### Ошибка подключения

1. Проверьте firewall на REG.RU
2. Убедитесь, что IP адрес добавлен в список разрешенных
3. Проверьте правильность учетных данных

### Ошибка импорта данных

1. Проверьте формат CSV файлов
2. Убедитесь, что кодировка UTF-8
3. Проверьте соответствие колонок

### Проблемы с JSON полями

MySQL использует тип `JSON` вместо `JSONB`. Если возникнут проблемы:

```sql
-- Проверка JSON поля
SELECT fldImgArray FROM vehicles LIMIT 1;

-- Обновление JSON поля
UPDATE vehicles SET fldImgArray = '[]' WHERE fldImgArray IS NULL;
```

## Завершение

После успешной миграции:

1. ✅ Данные перенесены в MySQL на REG.RU
2. ✅ Таблицы созданы и проиндексированы
3. ✅ Приложение настроено для работы с новой базой
4. ✅ Бекап PostgreSQL сохранен

Для отката используйте бекап `backup_postgres_20260420_170319.sql`
