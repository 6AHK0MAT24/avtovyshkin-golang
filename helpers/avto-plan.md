# План реализации микросервиса автовышек (avto-microservice)

## Обзор
Создание бекенд микросервиса для управления автовышками на основе архитектуры микросервиса водителей.

## Структура таблицы vehicles

### Основные поля
- `fldId` - UUID первичный ключ
- `fldGarageNumber` - Гаражный номер (обязательный)
- `fldVIN` - VIN код (обязательный)
- `fldHeight` - Высота автовышки (обязательный)
- `fldType` - Тип автовышки (Телескопическая, Телескоп + колено, Телескоп + стрела и рукоять)
- `fldPower` - Грузоподъемность (кг)
- `fldPrice5` - Цена за 5 часов
- `fldPrice22` - Цена за 22 часа
- `fldDescription` - Описание
- `fldBrand` - Марка ТС
- `fldMachine` - Модель подъемника

### Габариты ТС (разделить sizeTs)
- `fldLength` - Длина ТС в транспортном положении
- `fldWidth` - Ширина ТС в транспортном положении
- `fldHeight` - Высота ТС в транспортном положении
- `fldWidthWithSupports` - Ширина с опорами (отдельное поле width из constants.ts)
- `fldMass` - Масса ТС (тонн)
### Габариты люльки (разделить sizeCradle)
- `fldCradleWidthFolded` - Ширина люльки в сложенном состоянии
- `fldCradleWidthExtended` - Ширина люльки в разложенном состоянии
- `fldCradleLengthFolded` - Длина люльки в сложенном состоянии
- `fldCradleLengthExtended` - Длина люльки в разложенном состоянии

### Изображения
- `fldImgArray` - JSON массив путей к изображениям
- `fldMainImageIndex` - Индекс основной картинки (по умолчанию 1)

### Дополнительные поля
- `fldSpecial` - Особенности
- `fldRostechReg` - Регистрация в Ростехнадзоре (boolean)
- `fldStatus` - Статус (active/inactive/blocked)
- `fldCreatedAt` - Дата создания
- `fldUpdatedAt` - Дата обновления

## Шаги реализации

### 1. Создание структуры проекта ✅ ВЫПОЛНЕНО
- Скопировать структуру из drivers-microservice
- Переименовать пакеты с drivers-service на vehicles-service
- Создать директории: cmd, internal, migrations, uploads

### 2. Миграции БД ✅ ВЫПОЛНЕНО
- Создать `001_create_vehicles_table.sql` с таблицей vehicles
- Создать `002_insert_initial_vehicles.sql` с данными из constants.ts
- Добавить индексы: garage_number, vin, height, status
- Добавить триггер для автоматического обновления fldUpdatedAt

### 3. Модели (internal/models/vehicle.go) ✅ ВЫПОЛНЕНО
- Создать структуру Vehicle
- Создать VehicleStatus enum (active/inactive/blocked)
- Создать CreateVehicleRequest (обязательные: garageNumber, vin, height)
- Создать UpdateVehicleRequest
- Создать VehicleResponse
- Создать VehicleListResponse
- Реализовать методы ToResponse(), ToVehicle(), UpdateVehicle()
- Реализовать JSON marshaling/unmarshaling для fldImgArray

### 4. Repository (internal/repository/vehicle_repository.go) ✅ ВЫПОЛНЕНО
- Создать интерфейс VehicleRepository
- Реализовать методы:
  - Create(ctx, vehicle)
  - GetByID(ctx, id)
  - GetAll(ctx, limit, offset)
  - Update(ctx, vehicle)
  - Delete(ctx, id)
  - Search(ctx, query, limit, offset) - поиск по высоте и номеру
  - GetByGarageNumber(ctx, garageNumber)
  - GetByVIN(ctx, vin)
  - Close()
  
### 5. Service (internal/service/vehicle_service.go) ✅ ВЫПОЛНЕНО
- Создать интерфейс VehicleService
- Реализовать бизнес-логику:
  - Валидация VIN (17 символов)
  - Валидация гаражного номера
  - Валидация высоты (> 0)
  - Проверка уникальности garageNumber и VIN
  - CreateVehicle с валидацией
  - GetVehicle, GetVehicles с пагинацией
  - UpdateVehicle с валидацией
  - DeleteVehicle с удалением файлов
  - SearchVehicles по высоте и номеру
  - UploadVehicleImages (загрузка нескольких изображений)
  - DeleteVehicleImage
  - SetMainImageIndex

### 6. Handler (internal/handler/vehicle_handler.go) ✅ ВЫПОЛНЕНО
- Создать VehicleHandler
- Реализовать HTTP обработчики:
  - CreateVehicle - POST /api/vehicles
  - GetVehicle - GET /api/vehicles/{id}
  - GetVehicles - GET /api/vehicles (с пагинацией)
  - UpdateVehicle - PUT /api/vehicles/{id}
  - DeleteVehicle - DELETE /api/vehicles/{id}
  - SearchVehicles - GET /api/vehicles/search (по высоте и номеру)
  - UploadVehicleImages - POST /api/vehicles/{id}/images (multipart)
  - DeleteVehicleImage - DELETE /api/vehicles/{id}/images/{index}
  - SetMainImage - PUT /api/vehicles/{id}/main-image

### 7. WebSocket (internal/websocket/) ✅ ВЫПОЛНЕНО
- Адаптировать WebSocket из drivers-microservice
- Создать события: VehicleCreated, VehicleUpdated, VehicleDeleted
- Реализовать broadcast для автовышек

### 8. Middleware (internal/middleware/) ✅ ВЫПОЛНЕНО
- Скопировать middleware из drivers-microservice:
  - CORSMiddleware
  - LoggingMiddleware
  
  - RecoveryMiddleware
### 9. Config (internal/config/config.go) ✅ ВЫПОЛНЕНО
- Адаптировать конфигурацию из drivers-microservice
- Изменить имя сервиса на vehicles-service
- Настроить подключение к той же БД

### 10. Utils (internal/utils/) ✅ ВЫПОЛНЕНО
- Скопировать и адаптировать валидацию:
  - ValidateVIN (17 символов, алфавитно-цифровой)
  - ValidateGarageNumber
  - ValidateFileType (для изображений)
- Адаптировать file_utils для загрузки изображений

### 11. File Upload (internal/utils/file_utils.go) ✅
- Реализовать загрузку нескольких изображений
- Путь сохранения: `uploads/cars/{brand}_{height}/`
- Генерация уникальных имен файлов
- Обновление fldImgArray в БД
- Удаление файлов при удалении автовышки

### 12. Main (cmd/server/main.go) ✅ ВЫПОЛНЕНО
- Инициализация конфигурации
- Инициализация repository
- Инициализация WebSocket manager
- Инициализация service и handler
- Настройка роутера:
  - /api/vehicles - CRUD операции
  - /api/vehicles/search - поиск
  - /api/vehicles/{id}/images - загрузка/удаление изображений
  - /api/vehicles/{id}/main-image - установка основной картинки
  - /uploads/ - статические файлы
  - /ws - WebSocket
  - /health - health check
- Graceful shutdown

### 13. Dockerfile ✅ ВЫПОЛНЕНО
- Создать Dockerfile на основе drivers-microservice
- Настроить multi-stage build
- Экспозировать порт (отличный от drivers, например 8081)

### 14. Заполнение данными (migrations/002_insert_initial_vehicles.sql) ✅ ВЫПОЛНЕНО
- Парсинг данных из constants.ts
- Преобразование sizeCradle в отдельные поля (cradleWidthFolded, cradleWidthExtended, cradleLengthFolded, cradleLengthExtended)
- Преобразование sizeTs в length, width, height (габариты в транспортном положении)
- Добавление widthWithSupports (ширина с опорами)
- Установка mainImageIndex = 1 для всех записей
- Использование существующих путей изображений из constants.ts (поля img)### 15. Тестирование
- Unit тесты для service layer
- Unit тесты для repository layer
- Интеграционные тесты для API endpoints
- Тестирование загрузки изображений
- Тестирование WebSocket событий

## Поля для валидации

### Обязательные при создании
- garageNumber (гаражный номер)
- vin (VIN код)
- height (высота автовышки)

### Валидация VIN
- 17 символов
- Только алфавитно-цифровые символы (без I, O, Q)
- Уникальность в БД

### Валидация гаражного номера
- Не пустой
- Уникальность в БД

### Валидация высоты
- Число > 0

## Поиск

Поля для поиска:
- height (высота автовышки) - точное совпадение или диапазон
- garageNumber (гаражный номер) - частичное совпадение

## Пути сохранения изображений

Формат: `uploads/cars/{brand}_{height}/{uuid}_{original_filename}`

Пример: `uploads/cars/MITSUBISHI CANTER_17/550e8400-e29b-41d4-a716-446655440000_17.png`

## API Endpoints

### CRUD
- POST /api/vehicles - создание автовышки
- GET /api/vehicles - список с пагинацией
- GET /api/vehicles/{id} - получение по ID
- PUT /api/vehicles/{id} - обновление
- DELETE /api/vehicles/{id} - удаление

### Поиск
- GET /api/vehicles/search?height=17&garageNumber=А123 - поиск

### Изображения
- POST /api/vehicles/{id}/images - загрузка нескольких изображений
- DELETE /api/vehicles/{id}/images/{index} - удаление изображения по индексу
- PUT /api/vehicles/{id}/main-image - установка основной картинки (body: {"index": 2})

### WebSocket
- WS /ws - WebSocket соединение для real-time обновлений

### Статические файлы
- GET /uploads/* - отдача загруженных изображений

## Риски

1. **Конфликт портов** - drivers-microservice использует 8080, нужно использовать другой порт (8081)
2. **Общая БД** - обе службы используют одну БД, нужно убедиться в отсутствии конфликтов
3. **Миграции** - нужно аккуратно создавать миграции, чтобы не повредить данные drivers
4. **Загрузка файлов** - нужно обеспечить уникальность путей для изображений
5. **Валидация VIN** - сложный формат, нужна тщательная проверка

## Открытые вопросы

Нет - все требования уточнены.
