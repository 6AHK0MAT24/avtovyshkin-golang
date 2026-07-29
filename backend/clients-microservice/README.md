# Клиенты (Clients Microservice)

Микросервис для управления клиентами системы, поддерживающий работу с физическими и юридическими лицами.

## Архитектура

### Бэкенд (Go)
- **Порт**: 8003
- **База данных**: MySQL
- **Структура**:
  - `cmd/server/main.go` - точка входа приложения
  - `internal/config/` - конфигурация
  - `internal/models/` - модели данных
  - `internal/repository/` - работа с БД
  - `internal/service/` - бизнес-логика
  - `internal/handler/` - HTTP обработчики
  - `internal/middleware/` - middleware (CORS, логирование, recovery)
  - `migrations-mysql/` - миграции БД

### Фронтенд (React + TypeScript)
- **Компоненты**:
  - `ClientCard.tsx` - карточка клиента
  - `ClientForm.tsx` - форма создания/редактирования
  - `ClientList.tsx` - список клиентов
- **Страница**: `ClientsPage.tsx` - страница с табами для физических и юридических лиц
- **Сервис**: `clientService.ts` - API клиент

## База данных

### Таблица `clients`

Общая таблица для хранения данных о клиентах обоих типов.

#### Общие поля:
- `fldId` - UUID идентификатор
- `fldClientType` - тип клиента (individual/legal_entity)
- `fldPhone` - телефон
- `fldEmail` - email
- `fldAddress` - адрес
- `fldStatus` - статус (active/inactive/blocked)
- `fldNotes` - заметки
- `fldCreatedAt` - дата создания
- `fldUpdatedAt` - дата обновления

#### Поля для физических лиц:
- `fldFirstName` - имя
- `fldLastName` - фамилия
- `fldMiddleName` - отчество
- `fldBirthDate` - дата рождения
- `fldPassportSeries` - серия паспорта
- `fldPassportNumber` - номер паспорта
- `fldPassportIssueDate` - дата выдачи паспорта
- `fldPassportIssuedBy` - кем выдан паспорт
- `fldINN` - ИНН (12 цифр)

#### Поля для юридических лиц:
- `fldCompanyName` - название компании
- `fldCompanyLegalName` - полное юридическое название
- `fldOGRN` - ОГРН (15 цифр)
- `fldINNLegal` - ИНН (10 цифр)
- `fldKPP` - КПП (9 цифр)
- `fldLegalAddress` - юридический адрес
- `fldActualAddress` - фактический адрес
- `fldBankName` - название банка
- `fldBIC` - БИК (9 цифр)
- `fldAccountNumber` - расчетный счет (20 цифр)
- `fldCorrespondentAccount` - корреспондентский счет (20 цифр)
- `fldDirectorName` - ФИО директора
- `fldDirectorPosition` - должность директора
- `fldContactPerson` - контактное лицо
- `fldContactPersonPhone` - телефон контактного лица
- `fldContactPersonEmail` - email контактного лица

## API Эндпоинты

### Клиенты
- `GET /api/clients` - получить всех клиентов (с пагинацией)
- `GET /api/clients/type/{type}` - получить клиентов по типу (individual/legal)
- `GET /api/clients/{id}` - получить клиента по ID
- `POST /api/clients` - создать нового клиента
- `PUT /api/clients/{id}` - обновить клиента
- `DELETE /api/clients/{id}` - удалить клиента
- `GET /api/clients/search` - поиск клиентов (с фильтрами)

### Health Check
- `GET /health` - проверка работоспособности сервиса

## Запуск

### Запуск микросервиса клиентов
```powershell
.\start-clients.ps1
```

### Запуск всех сервисов
```powershell
.\start-all.ps1
```

## Применение миграций

Перед запуском сервиса необходимо применить миграции к базе данных:

```bash
psql -U postgres -d avtovyshkin -f backend/clients-microservice/migrations/000_create_update_timestamp_function.sql
psql -U postgres -d avtovyshkin -f backend/clients-microservice/migrations/001_create_clients_table.sql
```

## Функционал фронтенда

### Страница клиентов (/clients)
- **Табы**: переключение между физическими и юридическими лицами
- **Создание**: кнопка добавления нового клиента
- **Редактирование**: редактирование существующего клиента
- **Удаление**: удаление клиента с подтверждением
- **Пагинация**: постраничная навигация
- **Отображение**: карточки с информацией о клиентах

### Форма клиента
- Динамические поля в зависимости от типа клиента
- Валидация обязательных полей
- Маски для полей (ИНН, паспорт, банковские реквизиты)
