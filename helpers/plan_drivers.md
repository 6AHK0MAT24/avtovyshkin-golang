# План реализации CRUD для водителей (Drivers Service)

## Обзор проекта

**Технологический стек:**
- **Backend:** Go 1.26.1
- **Frontend:** React 19+, Vite, TypeScript, Ant Design
- **Database:** PostgreSQL 13.3 (порт 54320)
- **Real-time:** WebSocket (BEST PRACTICE для мгновенных обновлений)
- **File Storage:** Локальное хранение с организованной структурой (BEST PRACTICE)
- **Architecture:** Микросервисная архитектура

---

## Архитектура проекта

```
avtovyshkin-golang/
├── drivers-service/              # Микросервис водителей
│   ├── cmd/
│   │   └── server/
│   │       └── main.go          # Точка входа
│   ├── internal/
│   │   ├── config/              # Конфигурация
│   │   ├── models/              # Модели данных
│   │   ├── repository/          # Репозитории (Data Access Layer)
│   │   ├── service/             # Бизнес-логика
│   │   ├── handler/             # HTTP handlers
│   │   ├── websocket/           # WebSocket для real-time
│   │   ├── middleware/          # Middleware (logging, CORS, etc.)
│   │   └── utils/               # Утилиты
│   ├── migrations/              # SQL миграции
│   ├── uploads/                 # Хранилище файлов
│   │   ├── drivers/
│   │   │   ├── photos/          # Фото водителей
│   │   │   ├── licenses/        # Сканы водительских прав
│   │   │   └── passports/       # Сканы паспортов
│   │   └── temp/                # Временные файлы
│   ├── docs/                    # Swagger документация
│   ├── go.mod
│   └── go.sum
├── frontend/                     # React Frontend
│   ├── src/
│   │   ├── components/          # React компоненты
│   │   │   ├── drivers/
│   │   │   │   ├── DriverList.tsx
│   │   │   │   ├── DriverForm.tsx
│   │   │   │   ├── DriverCard.tsx
│   │   │   │   ├── DriverSearch.tsx
│   │   │   │   └── FileUpload.tsx
│   │   │   ├── common/
│   │   │   │   ├── Layout.tsx
│   │   │   │   ├── Header.tsx
│   │   │   │   └── Notification.tsx
│   │   │   └── ui/              # Обертки над Ant Design
│   │   ├── services/            # API сервисы
│   │   │   ├── api.ts           # Базовый API клиент
│   │   │   ├── driverService.ts # API для водителей
│   │   │   └── websocket.ts     # WebSocket клиент
│   │   ├── hooks/               # Custom React hooks
│   │   │   ├── useDrivers.ts
│   │   │   └── useWebSocket.ts
│   │   ├── types/               # TypeScript типы
│   │   │   └── driver.ts
│   │   ├── utils/               # Утилиты
│   │   ├── App.tsx
│   │   └── main.tsx
│   ├── public/
│   ├── package.json
│   ├── tsconfig.json
│   ├── vite.config.ts
│   └── tailwind.config.js
├── docker-compose.yml            # Docker для локальной разработки
├── .gitignore
└── README.md
```

---

## Детальный план реализации

### Этап 1: Настройка проекта и окружения

#### 1.1 Инициализация Go модуля
- [x] Создать структуру директорий для drivers-service
- [x] Инициализировать go.mod
- [x] Установить зависимости:
  - `github.com/gorilla/mux` - HTTP роутер
  - `github.com/lib/pq` - PostgreSQL драйвер
  - `github.com/gorilla/websocket` - WebSocket
  - `github.com/joho/godotenv` - Управление переменными окружения
  - `github.com/google/uuid` - Генерация UUID
  - `github.com/swaggo/gin-swagger` - Swagger документация
  - `github.com/rs/cors` - CORS middleware

#### 1.2 Настройка конфигурации
- [x] Создать файл `.env` с настройками:
  ```
  DB_HOST=localhost
  DB_PORT=54320
  DB_USER=postgres
  DB_PASSWORD=postgres
  DB_NAME=avtovyshkin_db_go
  SERVER_PORT=8080
  UPLOAD_DIR=./uploads
  MAX_UPLOAD_SIZE=10485760
  ```
- [x] Создать `internal/config/config.go` для загрузки конфигурации
#### 1.3 Инициализация React проекта
- [x] Создать React проект с Vite: `yarn create vite frontend --template react-ts`
- [x] Установить зависимости:
  - `antd` - UI библиотека
  - `@ant-design/icons` - Иконки
  - `axios` - HTTP клиент
  - `dayjs` - Работа с датами
  - `react-router-dom` - Роутинг
  - `zustand` - State management (BEST PRACTICE для React 19)
  - `@tanstack/react-query` - Server state management
- [x] Настроить TypeScript (strict mode)
- [x] Настроить Vite proxy для API запросов
---

### Этап 2: База данных

#### 2.1 Создание таблицы drivers
- [x] Создать миграцию `migrations/001_create_drivers_table.sql`:
  ```sql
  CREATE TABLE drivers (
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
      fldExperienceYears INTEGER DEFAULT 0,
      fldStatus VARCHAR(20) DEFAULT 'active' CHECK (fldStatus IN ('active', 'inactive', 'blocked')),
      fldCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

  CREATE INDEX idx_drivers_phone ON drivers(fldPhone);
  CREATE INDEX idx_drivers_email ON drivers(fldEmail);
  CREATE INDEX idx_drivers_license ON drivers(fldDriverLicenseNumber);
  CREATE INDEX idx_drivers_status ON drivers(fldStatus);

  CREATE TRIGGER update_drivers_updated_at
      BEFORE UPDATE ON drivers
      FOR EACH ROW
      EXECUTE FUNCTION update_updated_at_column();
  ```

- [x] Создать функцию для автоматического обновления `fldUpdatedAt`
- [x] Применить миграцию к базе данных

#### 2.2 Создание функции для обновления timestamp
- [x] Создать миграцию `migrations/000_create_update_timestamp_function.sql`
---

### Этап 3: Backend - Go Implementation

#### 3.1 Models (Модели данных)
- [x] Создать `internal/models/driver.go`:
  - Структура `Driver` с полями из БД
  - Методы для валидации
  - Теги JSON для API
  - Теги DB для ORM

- [x] Создать DTO (Data Transfer Objects):
  - `CreateDriverRequest` - для создания
  - `UpdateDriverRequest` - для обновления
  - `DriverResponse` - для ответа API
  - `DriverListResponse` - для списка с пагинацией
#### 3.2 Repository (Data Access Layer)
- [x] Создать `internal/repository/driver_repository.go`:
  - `Create(driver *Driver) error`
  - `GetByID(id string) (*Driver, error)`
  - `GetAll(limit, offset int) ([]Driver, int64, error)`
  - `Update(driver *Driver) error`
  - `Delete(id string) error`
  - `Search(query string, limit, offset int) ([]Driver, int64, error)`
  - `GetByPhone(phone string) (*Driver, error)`
  - `GetByLicenseNumber(licenseNumber string) (*Driver, error)`

- [x] Создать интерфейс `DriverRepository` для тестирования
- [x] Реализовать подключение к PostgreSQL
- [x] Добавить connection pooling

#### 3.3 Service (Бизнес-логика)
- [x] Создать `internal/service/driver_service.go`:
  - `CreateDriver(req CreateDriverRequest) (*Driver, error)`
  - `GetDriver(id string) (*Driver, error)`
  - `GetDrivers(page, pageSize int) ([]Driver, int64, error)`
  - `UpdateDriver(id string, req UpdateDriverRequest) (*Driver, error)`
  - `DeleteDriver(id string) error`
  - `SearchDrivers(query string, page, pageSize int) ([]Driver, int64, error)`
  - `UploadDriverPhoto(driverID string, file []byte, fileType string) (string, error)`

- [x] Добавить валидацию данных:
  - Проверка формата телефона
  - Проверка формата email
  - Проверка формата водительских прав
  - Проверка формата паспорта
  - Проверка возраста (минимум 18 лет)
  - Проверка срока действия прав

#### 3.4 WebSocket (Real-time updates)
- [x] Создать `internal/websocket/hub.go`:
  - Структура `Hub` для управления подключениями
  - Методы: `Register`, `Unregister`, `Broadcast`
  - Каналы для сообщений

- [x] Создать `internal/websocket/client.go`:
  - Структура `Client` для каждого подключения
  - Методы: `ReadPump`, `WritePump`

- [x] Создать `internal/websocket/manager.go`:
  - Глобальный менеджер WebSocket
  - Типы событий: `driver.created`, `driver.updated`, `driver.deleted`

- [x] Интегрировать WebSocket в Service слой для отправки уведомлений

#### 3.5 Handler (HTTP handlers)
- [x] Создать `internal/handler/driver_handler.go`:
  - `CreateDriver(w, r)` - POST /api/drivers
  - `GetDriver(w, r)` - GET /api/drivers/{id}
  - `GetDrivers(w, r)` - GET /api/drivers
  - `UpdateDriver(w, r)` - PUT /api/drivers/{id}
  - `DeleteDriver(w, r)` - DELETE /api/drivers/{id}
  - `SearchDrivers(w, r)` - GET /api/drivers/search
  - `UploadDriverPhoto(w, r)` - POST /api/drivers/{id}/photo
  - `UploadDriverLicense(w, r)` - POST /api/drivers/{id}/license
  - `UploadDriverPassport(w, r)` - POST /api/drivers/{id}/passport

- [x] Создать `internal/handler/websocket_handler.go`:
  - `HandleWebSocket(w, r)` - WS /api/ws

#### 3.6 Middleware
- [x] Создать `internal/middleware/logging.go` - Логирование запросов
- [x] Создать `internal/middleware/cors.go` - CORS настройки
- [x] Создать `internal/middleware/recovery.go` - Обработка паник
- [ ] Создать `internal/middleware/content_type.go` - Проверка Content-Type

#### 3.7 Utils (Утилиты)
- [x] Создать `internal/utils/file_utils.go`:
  - `SaveFile(file []byte, filename string) (string, error)`
  - `DeleteFile(filepath string) error`
  - `GenerateUniqueFilename(originalName string) string`
  - `ValidateFileType(filename string, allowedTypes []string) bool`

- [x] Создать `internal/utils/validation.go`:
  - `ValidatePhone(phone string) bool`
  - `ValidateEmail(email string) bool`
  - `ValidateDriverLicense(license string) bool`
  - `ValidatePassport(series, number string) bool`

#### 3.8 Main (Точка входа)
- [x] Создать `cmd/server/main.go`:
  - Загрузка конфигурации
  - Подключение к БД  - Инициализация репозиториев
  - Инициализация сервисов
  - Инициализация WebSocket Hub
  - Настройка роутера
  - Запуск HTTP сервера
  - Graceful shutdown

---

### Этап 4: Frontend - React Implementation

#### 4.1 Типы TypeScript
- [x] Создать `src/types/driver.ts`:
  - Интерфейс `Driver`
  - Интерфейс `CreateDriverForm`
  - Интерфейс `UpdateDriverForm`
  - Интерфейс `DriverFilters`
  - Тип `DriverStatus`

#### 4.2 API Services
- [x] Создать `src/services/api.ts`:
  - Базовый axios инстанс
  - Интерцепторы для запросов/ответов
  - Обработка ошибок

- [x] Создать `src/services/driverService.ts`:
  - `getDrivers(page, pageSize)` - GET /api/drivers
  - `getDriver(id)` - GET /api/drivers/{id}
  - `createDriver(data)` - POST /api/drivers
  - `updateDriver(id, data)` - PUT /api/drivers/{id}
  - `deleteDriver(id)` - DELETE /api/drivers/{id}
  - `searchDrivers(query, page, pageSize)` - GET /api/drivers/search
  - `uploadDriverPhoto(id, file)` - POST /api/drivers/{id}/photo
  - `uploadDriverLicense(id, file, type)` - POST /api/drivers/{id}/license
  - `uploadDriverPassport(id, file, type)` - POST /api/drivers/{id}/passport

- [x] Создать `src/services/websocket.ts`:
  - Класс `WebSocketClient`
  - Методы: `connect`, `disconnect`, `subscribe`, `unsubscribe`
  - Обработка событий: `driver.created`, `driver.updated`, `driver.deleted`

#### 4.3 State Management (Zustand)
- [x] Создать `src/store/driverStore.ts`:
  - Состояние: drivers, loading, error, pagination
  - Actions: fetchDrivers, addDriver, updateDriver, removeDriver
  - Реактивные селекторы

- [x] Создать `src/store/uiStore.ts`:
  - Состояние: notifications, modals, theme
  - Actions: showNotification, openModal, closeModal

#### 4.4 Custom Hooks
- [x] Создать `src/hooks/useDrivers.ts`:
  - `useDrivers()` - получение списка водителей
  - `useDriver(id)` - получение одного водителя
  - `useCreateDriver()` - создание водителя
  - `useUpdateDriver()` - обновление водителя
  - `useDeleteDriver()` - удаление водителя
  - `useSearchDrivers()` - поиск водителей

- [x] Создать `src/hooks/useWebSocket.ts`:
  - Подключение к WebSocket
  - Подписка на события
  - Автоматическое переподключение
  - Обновление данных при получении событий

#### 4.5 Components

##### 4.5.1 Layout Components
- [x] Создать `src/components/common/Layout.tsx`:
  - Основной layout с header и sidebar
  - Адаптивный дизайн

- [x] Создать `src/components/common/Header.tsx`:
  - Логотип
  - Навигация
  - Индикатор WebSocket соединения

- [x] Создать `src/components/common/Notification.tsx`:
  - Компонент для уведомлений
  - Использование Ant Design notification
##### 4.5.2 Driver Components
- [x] Создать `src/components/drivers/DriverList.tsx`:
  - Таблица водителей (Ant Design Table)
  - Пагинация
  - Сортировка
  - Фильтрация по статусу
  - Действия: просмотр, редактирование, удаление

- [x] Создать `src/components/drivers/DriverForm.tsx`:
  - Форма создания/редактирования водителя
  - Валидация полей
  - Загрузка файлов
  - Предпросмотр загруженных фото
  - Использование Ant Design Form

- [x] Создать `src/components/drivers/DriverCard.tsx`:
  - Карточка водителя для детального просмотра
  - Отображение фото и документов
  - Информация о водителе
- [x] Создать `src/components/drivers/DriverSearch.tsx`:
  - Поиск по водителям
  - Фильтры
  - Автодополнение

- [x] Создать `src/components/drivers/FileUpload.tsx`:
  - Компонент загрузки файлов
  - Drag & Drop
  - Предпросмотр изображений
  - Валидация типа и размера файла
  - Прогресс загрузки


##### 4.5.3 UI Components
- [x] Создать обертки над Ant Design компонентами для единообразия

#### 4.6 Pages
- [x] Создать `src/pages/DriversPage.tsx`:
  - Главная страница водителей
  - Интеграция всех компонентов
  - Обработка WebSocket событий

#### 4.7 App Configuration
- [x] Настроить `src/App.tsx`:
  - Роутинг с react-router-dom
  - Провайдеры (Zustand, Query)
  - Глобальные стили

- [x] Настроить `src/main.tsx`:
  - Подключение React
  - Подключение Ant Design
  - Инициализация WebSocket
---

### Этап 5: Интеграция и тестирование

#### 5.1 Backend Testing
- [x] Написать unit тесты для Repository слоя
- [x] Написать unit тесты для Service слоя
- [x] Написать интеграционные тесты для API endpoints
- [x] Тестирование WebSocket соединений

#### 5.2 Frontend Testing
- [ ] Написать unit тесты для компонентов
- [ ] Написать тесты для hooks
- [ ] E2E тесты с Playwright

#### 5.3 Integration Testing
- [ ] Тестирование полного цикла CRUD
- [ ] Тестирование real-time обновлений
- [ ] Тестирование загрузки файлов
- [ ] Тестирование валидации

---

### Этап 6: Документация и развертывание

#### 6.1 API Documentation
- [ ] Настроить Swagger для Go API
- [ ] Создать документацию для всех endpoints
- [ ] Добавить примеры запросов/ответов

#### 6.2 Project Documentation
- [x] Создать README.md с инструкциями:
  - Установка зависимостей
  - Настройка БД
  - Запуск backend
  - Запуск frontend
  - Структура проекта
#### 6.3 Docker Configuration
- [x] Создать `docker-compose.yml`:
  - PostgreSQL сервис
  - Drivers Service
  - Frontend сервис
  - Nginx (опционально)

- [x] Создать Dockerfile для backend
- [x] Создать Dockerfile для frontend
#### 6.4 Deployment Preparation
- [x] Настроить переменные окружения для production
- [ ] Настроить логирование
- [ ] Настроить мониторинг
- [ ] Подготовить скрипты для деплоя на reg.ru
---

## Технические решения (BEST PRACTICE)

### 1. Real-time Updates - WebSocket
**Почему WebSocket:**
- ✅ Двусторонняя связь
- ✅ Мгновенные обновления
- ✅ Меньше overhead чем polling
- ✅ Поддержка всеми современными браузерами
- ✅ Отлично работает с React

**Реализация:**
- Gorilla WebSocket для Go
- Native WebSocket API для React
- Hub pattern для управления подключениями
- Автоматическое переподключение на клиенте

### 2. File Storage - Локальное хранение с организованной структурой
**Почему локальное хранение:**
- ✅ Просто для разработки и тестирования
- ✅ Быстрый доступ к файлам
- ✅ Легкий бэкап
- ✅ Нет зависимости от внешних сервисов

**Структура:**
```
uploads/
├── drivers/
│   ├── {driver_id}/
│   │   ├── photo/
│   │   ├── license/
│   │   └── passport/
└── temp/
```

**Для production:**
- Переход на S3-совместимое хранилище
- Использование MinIO для локального S3

### 3. State Management - Zustand
**Почему Zustand:**
- ✅ Простой API
- ✅ Минимальный boilerplate
- ✅ Отличная производительность
- ✅ Поддержка TypeScript
- ✅ Совместим с React 19

### 4. API Client - Axios + React Query
**Почему:**
- ✅ Axios: надежный HTTP клиент с интерцепторами
- ✅ React Query: кэширование, автоматические refetch, optimistic updates
- ✅ Отличная интеграция с React 19

### 5. Form Handling - Ant Design Form
**Почему:**
- ✅ Встроенная валидация
- ✅ TypeScript поддержка
- ✅ Интеграция с Ant Design компонентами
- ✅ Простая работа с файлами

### 6. Database - PostgreSQL
**Почему:**
- ✅ ACID транзакции
- ✅ JSONB для гибких данных
- ✅ Отличная поддержка в Go
- ✅ Надежность и масштабируемость
- ✅ Поддержка на reg.ru

### 7. Go Architecture - Clean Architecture
**Почему:**
- ✅ Разделение ответственности
- ✅ Тестируемость
- ✅ Масштабируемость
- ✅ Поддержка микросервисов

**Слои:**
- Handler (Presentation)
- Service (Business Logic)
- Repository (Data Access)
- Models (Domain)

---

## API Endpoints

### Drivers API

```
POST   /api/drivers              - Создать водителя
GET    /api/drivers              - Получить список водителей (пагинация)
GET    /api/drivers/{id}         - Получить водителя по ID
PUT    /api/drivers/{id}         - Обновить водителя
DELETE /api/drivers/{id}         - Удалить водителя
GET    /api/drivers/search       - Поиск водителей
POST   /api/drivers/{id}/photo   - Загрузить фото водителя
POST   /api/drivers/{id}/license - Загрузить скан прав
POST   /api/drivers/{id}/passport - Загрузить скан паспорта
```

### WebSocket

```
WS /api/ws                       - WebSocket соединение
```

**WebSocket Events:**
```json
{
  "type": "driver.created",
  "data": { ...driver object... }
}

{
  "type": "driver.updated",
  "data": { ...driver object... }
}

{
  "type": "driver.deleted",
  "data": { "id": "uuid" }
}
```

---

## Валидация данных

### Обязательные поля при создании:
- fldFirstName
- fldLastName
- fldPhone (уникальный)
- fldDriverLicenseNumber (уникальный)
- fldDriverLicenseIssueDate
- fldDriverLicenseExpiryDate

### Форматы:
- Телефон: +7 (XXX) XXX-XX-XX
- Email: стандартный email формат
- Водительские права: 10 цифр (серия + номер)
- Паспорт: серия (4 цифры), номер (6 цифр)

### Ограничения:
- Возраст: минимум 18 лет
- Размер файла: максимум 10MB
- Типы файлов: jpg, jpeg, png, pdf

---

## Безопасность

### Backend:
- [ ] SQL Injection защита (prepared statements)
- [ ] XSS защита
- [ ] CORS настройки
- [ ] Rate limiting
- [ ] Валидация входных данных
- [ ] Санитизация загружаемых файлов

### Frontend:
- [ ] XSS защита (React по умолчанию)
- [ ] CSRF токены
- [ ] Валидация на клиенте и сервере
- [ ] Безопасное хранение токенов

---

## Мониторинг и логирование

### Backend:
- [ ] Структурированное логирование
- [ ] Логирование всех HTTP запросов
- [ ] Логирование ошибок
- [ ] Метрики производительности

### Frontend:
- [ ] Error boundaries
- [ ] Логирование ошибок в консоль
- [ ] Отправка ошибок на сервер (опционально)

---

## Следующие шаги после реализации Drivers Service

1. **Orders Service** - Микросервис заказов
2. **Vehicles Service** - Микросервис техники (автовышки)
3. **Users Service** - Микросервис пользователей и авторизации
4. **API Gateway** - Единая точка входа
5. **Notification Service** - Сервис уведомлений

---

## Вопросы для уточнения

1. **Аутентификация:** Нужна ли JWT авторизация для тестового сервиса?
2. **Роли пользователей:** Нужны ли разные роли (admin, manager, driver)?
3. **Логирование:** Какой уровень детализации логов нужен?
4. **Мониторинг:** Нужен ли Prometheus/Grafana для мониторинга?
5. **Тестирование:** Какой процент покрытия тестами нужен?
6. **Деплой:** Нужен ли CI/CD pipeline (GitHub Actions, GitLab CI)?

---

## Критерии успеха

- [ ] CRUD операции работают корректно
- [ ] Real-time обновления работают мгновенно
- [ ] Загрузка файлов работает без ошибок
- [ ] Валидация данных работает корректно
- [ ] UI/UX соответствует современным стандартам
- [ ] Код соответствует SOLID и KISS принципам
- [ ] Проект готов к развертыванию на reg.ru
- [ ] Документация полная и понятная
