# Автовышкин - Система управления водителями

Микросервисная система для управления водителями с использованием Go, React и PostgreSQL.

## Структура проекта

```
avtovyshkin-golang/
├── drivers-service/          # Backend (Go)
│   ├── cmd/                  # Точка входа приложения
│   ├── internal/             # Внутренние пакеты
│   │   ├── config/          # Конфигурация
│   │   ├── handler/         # HTTP handlers
│   │   ├── middleware/      # Middleware
│   │   ├── models/          # Модели данных
│   │   ├── repository/      # Data Access Layer
│   │   ├── service/         # Бизнес-логика
│   │   ├── utils/           # Утилиты
│   │   └── websocket/       # WebSocket для real-time
│   ├── migrations/          # SQL миграции
│   ├── uploads/             # Загруженные файлы
│   ├── go.mod
│   └── go.sum
├── frontend/                # Frontend (React + TypeScript)
│   ├── src/
│   │   ├── components/      # React компоненты
│   │   ├── services/        # API сервисы
│   │   ├── hooks/           # Custom hooks
│   │   ├── types/           # TypeScript типы
│   │   └── utils/           # Утилиты
│   ├── package.json
│   └── vite.config.ts
├── docker-compose.yml        # Docker Compose конфигурация
└── README.md
```

## Технологический стек

### Backend
- **Go 1.21** - Язык программирования
- **PostgreSQL 13.3** - База данных
- **Gorilla Mux** - HTTP роутер
- **Gorilla WebSocket** - WebSocket для real-time
- **lib/pq** - PostgreSQL драйвер

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
- Docker и Docker Compose
- Go 1.21+ (для локальной разработки)
- Node.js 18+ (для локальной разработки)
- Yarn (для локальной разработки)

### Запуск с Docker Compose

```bash
# Запуск всех сервисов
docker-compose up -d

# Просмотр логов
docker-compose logs -f

# Остановка сервисов
docker-compose down
```

### Локальная разработка

#### Рекомендуемый режим разработки (PostgreSQL в Docker + Hot Reload)

Для локальной разработки рекомендуется использовать PostgreSQL в Docker с hot reload для бэкенда:

```powershell
# 1. Запуск PostgreSQL в Docker
.\start-dev.ps1

# 2. Запуск бэкенда с hot reload (в новом терминале)
.\start-backend.ps1

# 3. Запуск фронтенда (в новом терминале)
.\start-frontend.ps1

# Остановка PostgreSQL
.\stop-dev.ps1
```

**Преимущества этого режима:**
- ✅ PostgreSQL работает в Docker (изолирован и легко управляем)
- ✅ Бэкенд запускается локально с hot reload (изменения применяются автоматически)
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
DB_PORT=54320
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=avtovyshkin_db_go
SERVER_PORT=8080
UPLOAD_DIR=./uploads
MAX_UPLOAD_SIZE=10485760
```

### Frontend (.env)

```
VITE_API_BASE_URL=http://localhost:8080
VITE_WS_BASE_URL=ws://localhost:8080
```

## Миграции базы данных

```bash
cd drivers-service

# Применить миграции
psql -h localhost -p 54320 -U postgres -d avtovyshkin_db_go -f migrations/000_create_update_timestamp_function.sql
psql -h localhost -p 54320 -U postgres -d avtovyshkin_db_go -f migrations/001_create_drivers_table.sql
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
