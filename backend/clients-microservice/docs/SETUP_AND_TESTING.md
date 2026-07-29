# Руководство по запуску и тестированию Clients Service

## Предварительные требования

### Необходимое программное обеспечение

- **Go** версии 1.21 или выше
- **MySQL** версии 8.0 или выше
- **PowerShell** (для Windows) или **Bash** (для Linux/Mac)
- **Git** (для клонирования репозитория)

### Проверка требований

```powershell
# Проверка версии Go
go version

# Проверка версии MySQL
mysql --version

# Проверка PowerShell
$PSVersionTable.PSVersion
```

## Установка и настройка

### 1. Клонирование репозитория

```powershell
git clone <repository-url>
cd avtovyshkin-golang/backend/clients-microservice
```

### 2. Настройка базы данных MySQL

#### Создание базы данных

```sql
CREATE DATABASE clients_service CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'clients_user'@'localhost' IDENTIFIED BY 'secure_password';
GRANT ALL PRIVILEGES ON clients_service.* TO 'clients_user'@'localhost';
FLUSH PRIVILEGES;
```

#### Запуск миграций

```powershell
# Применение миграций
mysql -u clients_user -p clients_service < migrations/000_create_update_timestamp_function.sql
mysql -u clients_user -p clients_service < migrations/001_create_clients_table.sql
mysql -u clients_user -p clients_service < migrations/002_insert_test_clients.sql
```

### 3. Настройка переменных окружения

Создайте файл `.env` в корневой директории проекта:

```env
# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8080
SERVER_READ_TIMEOUT=30s
SERVER_WRITE_TIMEOUT=30s
SERVER_IDLE_TIMEOUT=120s

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=clients_user
DB_PASSWORD=secure_password
DB_NAME=clients_service

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173,http://localhost:5174

# Environment
ENV=development
```

Или используйте готовый файл `.env.production`:

```powershell
cp .env.production .env
# Отредактируйте .env файл с вашими настройками
```

### 4. Установка зависимостей Go

```powershell
go mod download
go mod tidy
```

## Запуск сервиса

### Способ 1: Прямой запуск

```powershell
go run cmd/server/main.go
```

### Способ 2: Сборка и запуск

```powershell
# Сборка проекта
go build -o bin/server ./cmd/server

# Запуск собранного исполняемого файла
./bin/server
```

### Способ 3: Использование Docker

```powershell
# Сборка Docker образа
docker build -t clients-service .

# Запуск контейнера
docker run -p 8080:8080 --env-file .env clients-service
```

### Способ 4: Использование Docker Compose

```powershell
# Запуск всех сервисов
docker-compose up -d

# Просмотр логов
docker-compose logs -f clients-service
```

## Проверка работоспособности

### 1. Проверка здоровья сервиса

```powershell
curl http://localhost:8080/health
```

Ожидаемый ответ:
```
OK
```

### 2. Проверка документации API

Откройте в браузере:
```
http://localhost:8080/docs
```

### 3. Проверка списка клиентов

```powershell
curl http://localhost:8080/api/clients
```

## Тестирование API

### Автоматическое тестирование

Запустите автоматический тестовый скрипт:

```powershell
# Убедитесь, что сервис запущен
# Запустите тестовый скрипт
.\test-api.ps1
```

Тестовый скрипт проверит:
- ✓ Health check
- ✓ Доступность документации
- ✓ Создание клиента (физическое лицо)
- ✓ Создание клиента (юридическое лицо)
- ✓ Получение клиента по ID
- ✓ Получение клиентов по типу
- ✓ Поиск клиентов
- ✓ Обновление клиента
- ✓ Удаление клиента
- ✓ Пагинацию
- ✓ Фильтрацию по типу

### Ручное тестирование

#### 1. Создание клиента

```powershell
curl -X POST "http://localhost:8080/api/clients" `
  -H "Content-Type: application/json" `
  -d '{
    "name": "Иван Иванов",
    "type": "individual",
    "status": "active",
    "phone": "+79001234567",
    "email": "ivan@example.com",
    "address": "г. Москва, ул. Примерная, д. 1"
  }'
```

#### 2. Получение списка клиентов

```powershell
curl "http://localhost:8080/api/clients?page=1&perPage=10"
```

#### 3. Получение клиента по ID

```powershell
curl "http://localhost:8080/api/clients/{id}"
```

#### 4. Обновление клиента

```powershell
curl -X PUT "http://localhost:8080/api/clients/{id}" `
  -H "Content-Type: application/json" `
  -d '{
    "name": "Иван Иванович",
    "status": "inactive"
  }'
```

#### 5. Удаление клиента

```powershell
curl -X DELETE "http://localhost:8080/api/clients/{id}"
```

#### 6. Поиск клиентов

```powershell
curl "http://localhost:8080/api/clients/search?query=Иван&page=1&perPage=10"
```

#### 7. Получение клиентов по типу

```powershell
curl "http://localhost:8080/api/clients/type/individual?page=1&perPage=10"
```

## Мониторинг и логирование

### Просмотр логов

Логи сервиса выводятся в консоль и содержат:
- Время запроса
- Метод HTTP
- Путь запроса
- IP адрес клиента
- Код статуса ответа
- Время выполнения

Пример лога:
```
2024/01/15 10:30:45 GET /api/clients - 200 OK - 127.0.0.1 - 45ms
```

### Проверка статуса сервиса

```powershell
# Health check
curl http://localhost:8080/health

# Проверка документации
curl http://localhost:8080/docs
```

## Устранение неполадок

### Проблема: Сервис не запускается

**Возможные причины:**
1. База данных недоступна
2. Неверные параметры подключения к БД
3. Порт уже занят

**Решения:**
```powershell
# Проверка подключения к MySQL
mysql -u clients_user -p -h localhost clients_service

# Проверка занятости порта
netstat -ano | findstr :8080

# Изменение порта в .env файле
SERVER_PORT=8081
```

### Проблема: Ошибки при выполнении запросов

**Возможные причины:**
1. Неверный формат JSON
2. Отсутствие обязательных полей
3. Неверный тип данных

**Решения:**
```powershell
# Проверка формата JSON
# Используйте валидаторы JSON или Postman для проверки запросов

# Проверка логов сервиса для детальной информации об ошибках
```

### Проблема: CORS ошибки

**Возможные причины:**
1. Фронтенд не в списке разрешенных источников
2. Неверная конфигурация CORS

**Решения:**
```powershell
# Проверка и обновление CORS_ALLOWED_ORIGINS в .env
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173,http://your-frontend-domain.com
```

### Проблема: Ошибки миграции базы данных

**Возможные причины:**
1. Таблицы уже существуют
2. Неверные права доступа
3. Неверная кодировка

**Решения:**
```powershell
# Проверка существования таблиц
mysql -u clients_user -p clients_service -e "SHOW TABLES;"

# Удаление существующих таблиц (внимательно!)
mysql -u clients_user -p clients_service -e "DROP TABLE IF EXISTS clients;"

# Повторное применение миграций
mysql -u clients_user -p clients_service < migrations/001_create_clients_table.sql
```

## Производительность и оптимизация

### Настройка пула соединений

В файле `internal/repository/client_repository.go` можно настроить параметры пула соединений:

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(25)
db.SetConnMaxLifetime(5 * time.Minute)
```

### Мониторинг производительности

Используйте следующие метрики для мониторинга:
- Время ответа API
- Количество активных соединений с БД
- Использование памяти
- CPU нагрузка

## Безопасность

### Рекомендации по безопасности

1. **Используйте сильные пароли** для базы данных
2. **Ограничьте доступ** к базе данных по IP
3. **Используйте HTTPS** в production
4. **Настройте firewall** для ограничения доступа
5. **Регулярно обновляйте** зависимости Go
6. **Внедрите аутентификацию** для API endpoints

### Настройка HTTPS

Для production среды используйте обратный прокси (nginx, Apache) с SSL:

```nginx
server {
    listen 443 ssl;
    server_name your-domain.com;

    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## Развертывание в production

### Подготовка к production

1. **Измените переменные окружения** в `.env.production`
2. **Настройте HTTPS** с помощью SSL сертификатов
3. **Настройте мониторинг** и логирование
4. **Настройте резервное копирование** базы данных
5. **Протестируйте** все функции в staging среде

### Использование systemd (Linux)

Создайте файл `/etc/systemd/system/clients-service.service`:

```ini
[Unit]
Description=Clients Service
After=network.target mysql.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/clients-service
ExecStart=/opt/clients-service/bin/server
Restart=always
RestartSec=10
Environment=ENV=production

[Install]
WantedBy=multi-user.target
```

Запуск сервиса:
```bash
sudo systemctl daemon-reload
sudo systemctl enable clients-service
sudo systemctl start clients-service
sudo systemctl status clients-service
```

## Дополнительные ресурсы

- [API Documentation](./API_DOCUMENTATION.md)
- [Project README](../README.md)
- [Go Documentation](https://golang.org/doc/)
- [MySQL Documentation](https://dev.mysql.com/doc/)

## Поддержка

При возникновении проблем:
1. Проверьте логи сервиса
2. Ознакомьтесь с разделом "Устранение неполадок"
3. Проверьте документацию API
4. Создайте issue в репозитории проекта

## Чек-лист перед запуском в production

- [ ] Настроена база данных MySQL
- [ ] Применены все миграции
- [ ] Настроены переменные окружения
- [ ] Проверено подключение к базе данных
- [ ] Настроен HTTPS
- [ ] Настроен firewall
- [ ] Настроено резервное копирование
- [ ] Протестированы все API endpoints
- [ ] Настроено логирование и мониторинг
- [ ] Проверена производительность
- [ ] Настроены оповещения об ошибках