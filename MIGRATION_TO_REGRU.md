# Инструкция по миграции на reg.ru

## Подготовка к миграции

### 1. Создание базы данных на reg.ru

1. Зайдите в панель управления reg.ru
2. Перейдите в раздел "Базы данных" → "PostgreSQL"
3. Создайте новую базу данных с параметрами:
   - Версия: PostgreSQL 13.x или выше
   - Кодировка: UTF-8
   - Локаль: ru_RU.UTF-8 или en_US.UTF-8

Сохраните данные подключения:
- Хост (например: postgresql.reg.ru)
- Порт (обычно 5432)
- Имя базы данных
- Имя пользователя
- Пароль

### 2. Настройка переменных окружения

Создайте файл `.env.production` на основе `.env.example`:

```bash
# PostgreSQL Configuration (reg.ru)
POSTGRES_USER=your-username
POSTGRES_PASSWORD=your-password
POSTGRES_DB=your-database-name
POSTGRES_PORT=5432
POSTGRES_HOST=your-reg.ru-host

# Backend Configuration
SERVER_PORT=8080

# Frontend Configuration
FRONTEND_PORT=5173
```

### 3. Экспорт данных из локальной базы

Выполните экспорт данных из локального контейнера:

```powershell
# Остановите контейнеры
docker-compose down

# Запустите только postgres
docker-compose up -d postgres

# Подождите пока база станет доступной
Start-Sleep -Seconds 5

# Экспорт схемы и данных
docker exec avtovyshkin-postgres pg_dump -U postgres avtovyshkin_db_go > backup.sql

# Или экспорт только данных (без схемы)
docker exec avtovyshkin-postgres pg_dump -U postgres -a avtovyshkin_db_go > data_only.sql
```

### 4. Импорт данных на reg.ru

Способ 1: Через pgAdmin или DBeaver
1. Подключитесь к базе данных на reg.ru
2. Откройте SQL редактор
3. Выполните миграции по порядку:
   - `000_create_update_timestamp_function.sql`
   - `001_create_drivers_table.sql`
   - `002_insert_test_drivers.sql`

Способ 2: Через командную строку
```bash
psql -h your-reg.ru-host -p 5432 -U your-username -d your-database -f backup.sql
```

### 5. Обновление docker-compose.yml для production

Для работы с внешней базой данных на reg.ru:

```yaml
version: '3.8'

services:
  drivers-microservice:
    build:
      context: ./backend/drivers-microservice
      dockerfile: Dockerfile
    container_name: avtovyshkin-drivers-microservice
    environment:
      DB_HOST: ${POSTGRES_HOST}
      DB_PORT: ${POSTGRES_PORT}
      DB_USER: ${POSTGRES_USER}
      DB_PASSWORD: ${POSTGRES_PASSWORD}
      DB_NAME: ${POSTGRES_DB}
      SERVER_PORT: ${SERVER_PORT}
    ports:
      - "${SERVER_PORT}:8080"
    networks:
      - avtovyshkin-network
    volumes:
      - ./backend/drivers-microservice/uploads:/app/uploads

  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    container_name: avtovyshkin-frontend
    ports:
      - "${FRONTEND_PORT}:5173"
    depends_on:
      - drivers-microservice
    networks:
      - avtovyshkin-network

networks:
  avtovyshkin-network:
    driver: bridge
```

### 6. Запуск на reg.ru

```powershell
# Загрузите файлы на сервер reg.ru
# Установите Docker и Docker Compose на сервер

# Запустите с production конфигурацией
docker-compose --env-file .env.production up -d
```

## Проверка работоспособности

### Проверка подключения к базе данных

```powershell
# Локальная проверка
docker exec -it avtovyshkin-postgres psql -U postgres -d avtovyshkin_db_go -c "SELECT COUNT(*) FROM drivers;"

# Проверка на reg.ru
psql -h your-reg.ru-host -p 5432 -U your-username -d your-database -c "SELECT COUNT(*) FROM drivers;"
```

### Проверка API

```powershell
# Проверка health endpoint
curl http://localhost:8080/health

# Получение списка водителей
curl http://localhost:8080/api/drivers
```

## Резервное копирование

### Регулярное резервное копирование

Создайте скрипт `backup.sh`:

```bash
#!/bin/bash
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/backups"
mkdir -p $BACKUP_DIR

# Локальный бэкап
docker exec avtovyshkin-postgres pg_dump -U postgres avtovyshkin_db_go > $BACKUP_DIR/backup_$DATE.sql

# Бэкап с reg.ru
pg_dump -h your-reg.ru-host -p 5432 -U your-username your-database > $BACKUP_DIR/production_$DATE.sql

# Хранить последние 7 бэкапов
find $BACKUP_DIR -name "backup_*.sql" -mtime +7 -delete
find $BACKUP_DIR -name "production_*.sql" -mtime +7 -delete
```

## Восстановление из бэкапа

```bash
# Восстановление на локальной базе
docker exec -i avtovyshkin-postgres psql -U postgres avtovyshkin_db_go < backup.sql

# Восстановление на reg.ru
psql -h your-reg.ru-host -p 5432 -U your-username your-database < backup.sql
```

## Полезные команды

### Работа с Docker

```powershell
# Просмотр логов
docker-compose logs -f postgres
docker-compose logs -f drivers-microservice

# Перезапуск сервисов
docker-compose restart

# Остановка всех сервисов
docker-compose down

# Полная очистка (внимание: удаляет данные!)
docker-compose down -v
```

### Работа с PostgreSQL

```powershell
# Подключение к базе
docker exec -it avtovyshkin-postgres psql -U postgres -d avtovyshkin_db_go

# Просмотр таблиц
\dt

# Просмотр структуры таблицы
\d drivers

# Выход из psql
\q
```

## Безопасность

1. **Никогда не коммитите `.env` файлы в Git**
2. Используйте сложные пароли для базы данных
3. Ограничьте доступ к базе данных по IP
4. Используйте SSL соединения для production
5. Регулярно обновляйте пароли
6. Настройте firewall на сервере reg.ru

## Мониторинг

Рекомендуемые инструменты для мониторинга:
- Prometheus + Grafana для метрик
- pgAdmin для управления базой данных
- Sentry для логирования ошибок

## Поддержка

При возникновении проблем:
1. Проверьте логи: `docker-compose logs`
2. Проверьте подключение к базе данных
3. Убедитесь, что все миграции применены
4. Проверьте переменные окружения
