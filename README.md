# Автовышкин - Система управления водителями и автовышками

Микросервисная система для управления водителями и автовышками с использованием Go, React и MySQL.

**⚠️ Важно:** Проект использует MySQL базу данных. Локальная разработка использует MySQL через OpenServer на localhost:3306.## Структура проекта
```
avtovyshkin-golang/
├── backend/
│   ├── drivers-microservice/ # Drivers Service (Go)
│   │   ├── cmd/              # Точка входа приложения
│   │   ├── internal/         # Внутренние пакеты
│   │   │   ├── config/      # Конфигурация
│   │   │   ├── handler/     # HTTP handlers
│   │   │   ├── middleware/  # Middleware
│   │   │   ├── models/      # Модели данных
│   │   │   ├── repository/  # Data Access Layer
│   │   │   ├── service/     # Бизнес-логика
│   │   │   ├── utils/       # Утилиты
│   │   │   └── websocket/   # WebSocket для real-time
│   │   ├── migrations/      # SQL миграции
│   │   ├── uploads/         # Загруженные файлы
│   │   ├── go.mod
│   │   └── go.sum
│   └── vehicles-microservice/ # Vehicles Service (Go)
│       ├── cmd/              # Точка входа приложения
│       ├── internal/         # Внутренние пакеты
│       │   ├── config/      # Конфигурация
│       │   ├── handler/     # HTTP handlers
│       │   ├── middleware/  # Middleware
│       │   ├── models/      # Модели данных
│       │   ├── repository/  # Data Access Layer
│       │   ├── service/     # Бизнес-логика
│       │   ├── utils/       # Утилиты
│       │   └── websocket/   # WebSocket для real-time
│       ├── migrations/      # SQL миграции
│       ├── uploads/         # Загруженные файлы
│       ├── go.mod
│       └── go.sum
├── frontend/                # Frontend (React + TypeScript)
│   ├── src/
│   │   ├── components/      # React компоненты
│   │   ├── services/        # API сервисы
│   │   ├── hooks/           # Custom hooks
│   │   ├── types/           # TypeScript типы
│   │   └── utils/           # Утилиты
│   ├── package.json
│   └── vite.config.ts
└── README.md```

## Технологический стек

### Drivers Service (Backend)
- **Go 1.23** - Язык программирования
- **MySQL** - База данных
- **Gorilla Mux** - HTTP роутер
- **Gorilla WebSocket** - WebSocket для real-time
- **go-sql-driver/mysql** - MySQL драйвер

### Vehicles Service (Backend)
- **Go 1.23** - Язык программирования
- **MySQL** - База данных
- **Gorilla Mux** - HTTP роутер
- **Gorilla WebSocket** - WebSocket для real-time
- **sqlx** - MySQL драйвер
### Frontend
- **React 19** - UI библиотека
- **TypeScript** - Типизация
- **Vite** - Сборщик
- **Ant Design 6** - UI компоненты
- **React Router DOM** - Роутинг
- **TanStack Query** - Server state management
- **Zustand** - Client state management
- **Axios** - HTTP клиент

## Быстрый старт

### Требования
- Go 1.23+ (для локальной разработки)
- Node.js 18+ (для локальной разработки)
- Yarn (для локальной разработки)
- MySQL (OpenServer или другой MySQL сервер)

### Локальная разработка

#### Быстрый запуск с PowerShell скриптами

Для удобного запуска проекта используйте готовые скрипты:

```powershell
# Запуск всего проекта (Backend + Frontend)
.\start-all.ps1

# Запуск только Backend
.\start-backend.ps1

# Запуск только Vehicles Service
.\start-vehicles.ps1

# Запуск только Frontend
.\start-frontend.ps1

# Остановка всех сервисов
.\stop-all.ps1
```

### Локальная разработка

#### Рекомендуемый режим разработки (PostgreSQL в Docker + Hot Reload)

Для локальной разработки рекомендуется использовать PostgreSQL в Docker с hot reload для микросервисов:

```powershell
# 1. Запуск PostgreSQL в Docker
.\start-dev.ps1

# 2. Запуск Drivers Service с hot reload (в новом терминале)
.\start-backend.ps1

# 3. Запуск Vehicles Service с hot reload (в новом терминале)
.\start-vehicles.ps1

# 4. Запуск фронтенда (в новом терминале)
.\start-frontend.ps1

# Остановка PostgreSQL
.\stop-dev.ps1
```

**Преимущества этого режима:**
- ✅ PostgreSQL работает в Docker (изолирован и легко управляем)
- ✅ Микросервисы запускаются локально с hot reload (изменения применяются автоматически)
- ✅ Быстрый цикл разработки без пересборки Docker контейнеров
- ✅ Полный доступ к отладке и логам

#### Быстрый запуск с PowerShell скриптами

Для удобного запуска проекта используйте готовые скрипты:

```powershell
# Запуск всего проекта (PostgreSQL + Backend + Frontend)
.\start-all.ps1

# Запуск только Backend
.\start-backend.ps1

# Запуск только Frontend
.\start-frontend.ps1

# Остановка всех сервисов
.\stop-all.ps1
```

Подробнее о скриптах см. [START.md](START.md)

#### Ручной запуск

**Backend**

```bash
cd backend/drivers-microservice

# Установка зависимостей
go mod download

# Установка air для hot reload (опционально)
go install github.com/cosmtrek/air@latest

# Запуск с hot reload
air

# Или обычный запуск
go run cmd/server/main.go
```
```

**Frontend**

```bash
cd frontend

# Установка зависимостей
yarn install

# Запуск dev сервера
yarn dev
```
```

## API Эндпоинты

### Drivers

- `GET /api/drivers` - Получить список водителей
- `GET /api/drivers/:id` - Получить водителя по ID
- `POST /api/drivers` - Создать водителя
- `PUT /api/drivers/:id` - Обновить водителя
- `DELETE /api/drivers/:id` - Удалить водителя
- `GET /api/drivers/search` - Поиск водителей

### File Upload

- `POST /api/drivers/:id/photo` - Загрузить фото водителя
- `POST /api/drivers/:id/license` - Загрузить скан водительских прав
- `POST /api/drivers/:id/passport` - Загрузить скан паспорта

### WebSocket

- `WS /ws` - WebSocket соединение для real-time обновлений

## Переменные окружения

### Backend (.env)

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=u3424187_root_avtovyshkin
DB_PASSWORD=JavaScript6315
DB_NAME=u3424187_avtovyshkin
SERVER_PORT=8082
UPLOAD_DIR=./uploads
MAX_UPLOAD_SIZE=10485760
```
```

### Frontend (.env)

```
VITE_API_BASE_URL=http://localhost:8082
VITE_WS_BASE_URL=ws://localhost:8082
```
```

## Миграции базы данных

```bash
cd backend/drivers-microservice

# Применить миграции MySQL
mysql -h localhost -u u3424187_root_avtovyshkin -p u3424187_avtovyshkin < migrations-mysql/001_create_drivers_table.sql
```
```

## Разработка

### Добавление новых фич

1. Создайте модель в `internal/models/`
2. Создайте репозиторий в `internal/repository/`
3. Создайте сервис в `internal/service/`
4. Создайте handler в `internal/handler/`
5. Добавьте роут в `cmd/server/main.go`
6. Создайте типы в `frontend/src/types/`
7. Создайте API сервис в `frontend/src/services/`
8. Создайте компоненты в `frontend/src/components/`

### Код стиль

- Go: следуйте [Effective Go](https://golang.org/doc/effective_go)
- TypeScript: используйте строгий режим
- React: используйте функциональные компоненты и hooks

## Лицензия

MIT
