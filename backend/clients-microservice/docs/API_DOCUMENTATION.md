# Clients Service API Documentation

## Обзор

Clients Service - это микросервис для управления клиентами в системе Автовышкин. Сервис предоставляет RESTful API для выполнения CRUD операций с клиентами.

## Базовая информация

- **Версия API**: 1.0
- **Базовый URL**: `http://localhost:8080/api`
- **Формат данных**: JSON
- **Кодировка**: UTF-8

## Аутентификация

В текущей версии API аутентификация не требуется. Все endpoints доступны без авторизации.

## Структура ответов

### Успешный ответ

```json
{
  "success": true,
  "data": { ... },
  "pagination": {
    "page": 1,
    "perPage": 10,
    "total": 100,
    "totalPages": 10
  }
}
```

### Ответ с ошибкой

```json
{
  "success": false,
  "error": "Error type",
  "message": "Detailed error message"
}
```

## API Endpoints

### Клиенты

#### 1. Получить список всех клиентов

**Endpoint**: `GET /api/clients`

**Описание**: Получить список всех клиентов с пагинацией.

**Параметры запроса**:
- `page` (optional, integer): Номер страницы. По умолчанию: 1
- `perPage` (optional, integer): Количество элементов на странице. По умолчанию: 10

**Пример запроса**:
```bash
curl -X GET "http://localhost:8080/api/clients?page=1&perPage=10"
```

**Пример ответа**:
```json
{
  "success": true,
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "Иван Иванов",
      "type": "individual",
      "status": "active",
      "phone": "+79001234567",
      "email": "ivan@example.com",
      "address": "г. Москва, ул. Примерная, д. 1",
      "createdAt": "2024-01-15T10:30:00Z",
      "updatedAt": "2024-01-15T10:30:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "perPage": 10,
    "total": 1,
    "totalPages": 1
  }
}
```

#### 2. Получить клиента по ID

**Endpoint**: `GET /api/clients/{id}`

**Описание**: Получить информацию о конкретном клиенте по его ID.

**Параметры URL**:
- `id` (required, string): UUID клиента

**Пример запроса**:
```bash
curl -X GET "http://localhost:8080/api/clients/123e4567-e89b-12d3-a456-426614174000"
```

**Пример ответа**:
```json
{
  "success": true,
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Иван Иванов",
    "type": "individual",
    "status": "active",
    "phone": "+79001234567",
    "email": "ivan@example.com",
    "address": "г. Москва, ул. Примерная, д. 1",
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-01-15T10:30:00Z"
  }
}
```

#### 3. Получить клиентов по типу

**Endpoint**: `GET /api/clients/type/{type}`

**Описание**: Получить список клиентов определенного типа.

**Параметры URL**:
- `type` (required, string): Тип клиента (`individual` или `legal`)

**Параметры запроса**:
- `page` (optional, integer): Номер страницы. По умолчанию: 1
- `perPage` (optional, integer): Количество элементов на странице. По умолчанию: 10

**Пример запроса**:
```bash
curl -X GET "http://localhost:8080/api/clients/type/individual?page=1&perPage=10"
```

#### 4. Поиск клиентов

**Endpoint**: `GET /api/clients/search`

**Описание**: Поиск клиентов по имени, телефону или email.

**Параметры запроса**:
- `query` (required, string): Строка поиска
- `page` (optional, integer): Номер страницы. По умолчанию: 1
- `perPage` (optional, integer): Количество элементов на странице. По умолчанию: 10

**Пример запроса**:
```bash
curl -X GET "http://localhost:8080/api/clients/search?query=Иван&page=1&perPage=10"
```

#### 5. Создать нового клиента

**Endpoint**: `POST /api/clients`

**Описание**: Создать нового клиента.

**Тело запроса**:
```json
{
  "name": "Иван Иванов",
  "type": "individual",
  "status": "active",
  "phone": "+79001234567",
  "email": "ivan@example.com",
  "address": "г. Москва, ул. Примерная, д. 1"
}
```

**Пример запроса**:
```bash
curl -X POST "http://localhost:8080/api/clients" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Иван Иванов",
    "type": "individual",
    "status": "active",
    "phone": "+79001234567",
    "email": "ivan@example.com",
    "address": "г. Москва, ул. Примерная, д. 1"
  }'
```

**Пример ответа**:
```json
{
  "success": true,
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Иван Иванов",
    "type": "individual",
    "status": "active",
    "phone": "+79001234567",
    "email": "ivan@example.com",
    "address": "г. Москва, ул. Примерная, д. 1",
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-01-15T10:30:00Z"
  }
}
```

#### 6. Обновить клиента

**Endpoint**: `PUT /api/clients/{id}`

**Описание**: Обновить информацию о существующем клиенте.

**Параметры URL**:
- `id` (required, string): UUID клиента

**Тело запроса** (все поля необязательные):
```json
{
  "name": "Иван Иванович",
  "type": "individual",
  "status": "inactive",
  "phone": "+79001234568",
  "email": "ivan.new@example.com",
  "address": "г. Москва, ул. Новая, д. 2"
}
```

**Пример запроса**:
```bash
curl -X PUT "http://localhost:8080/api/clients/123e4567-e89b-12d3-a456-426614174000" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Иван Иванович",
    "status": "inactive"
  }'
```

**Пример ответа**:
```json
{
  "success": true,
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Иван Иванович",
    "type": "individual",
    "status": "inactive",
    "phone": "+79001234568",
    "email": "ivan.new@example.com",
    "address": "г. Москва, ул. Новая, д. 2",
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-01-15T11:00:00Z"
  }
}
```

#### 7. Удалить клиента

**Endpoint**: `DELETE /api/clients/{id}`

**Описание**: Удалить клиента по ID.

**Параметры URL**:
- `id` (required, string): UUID клиента

**Пример запроса**:
```bash
curl -X DELETE "http://localhost:8080/api/clients/123e4567-e89b-12d3-a456-426614174000"
```

**Пример ответа**:
```json
{
  "success": true,
  "message": "Client deleted successfully"
}
```

### Системные endpoints

#### 8. Проверка здоровья сервиса

**Endpoint**: `GET /health`

**Описание**: Проверить состояние сервиса.

**Пример запроса**:
```bash
curl -X GET "http://localhost:8080/health"
```

**Пример ответа**:
```
OK
```

#### 9. Документация API

**Endpoint**: `GET /docs`

**Описание**: Получить HTML документацию API.

**Пример запроса**:
```bash
curl -X GET "http://localhost:8080/docs"
```

## Типы данных

### ClientType

- `individual` - Физическое лицо
- `legal` - Юридическое лицо

### ClientStatus

- `active` - Активный
- `inactive` - Неактивный
- `blocked` - Заблокированный

## Обработка ошибок

### Коды ошибок HTTP

- `200 OK` - Успешный запрос
- `400 Bad Request` - Неверный формат запроса
- `404 Not Found` - Ресурс не найден
- `405 Method Not Allowed` - Метод не поддерживается
- `500 Internal Server Error` - Внутренняя ошибка сервера

### Примеры ошибок

#### Клиент не найден

```json
{
  "success": false,
  "error": "Client not found",
  "message": "Клиент с указанным ID не найден"
}
```

#### Неверный формат запроса

```json
{
  "success": false,
  "error": "Invalid request",
  "message": "Неверный формат данных"
}
```

## Ограничения

- Максимальное количество элементов на странице: 100
- Минимальное количество элементов на странице: 1
- Длина строки поиска: минимум 2 символа

## CORS

API поддерживает CORS для следующих доменов:
- `http://localhost:3000`
- `http://localhost:5173`

Поддерживаемые методы:
- GET
- POST
- PUT
- DELETE
- OPTIONS

## Логирование

Все запросы логируются с следующей информацией:
- Время запроса
- Метод HTTP
- Путь
- IP адрес клиента
- Код статуса ответа
- Время выполнения

## Примечания

- Все даты и времени возвращаются в формате ISO 8601 (UTC)
- Все ID возвращаются в формате UUID v4
- API поддерживает пагинацию для всех списочных endpoints
- Поля в запросах на обновление являются необязательными
- Сервис автоматически устанавливает `createdAt` и `updatedAt` timestamps