# ИНСТРУКЦИЯ ПО МИГРАЦИИ НА REG.RU

## ✅ Что уже сделано:

1. ✅ Создан бекап PostgreSQL: `backup_postgres_20260420_170319.sql`
2. ✅ Экспортированы данные в CSV файлы:
   - `drivers_export.csv` (6 строк)
   - `vehicles_export.csv` (9 строк)
3. ✅ Созданы SQL скрипты для создания таблиц:
   - `sql/create_drivers_table_regru.sql`
   - `sql/create_vehicles_table_regru.sql`
4. ✅ Обновлены конфигурационные файлы для REG.RU:
   - `backend/drivers-microservice/.env.production`
   - `backend/vehicles-microservice/.env.production`
   - `frontend/.env.production`

## 📋 Пошаговая инструкция по миграции:

### Шаг 1: Подключение к REG.RU

1. Откройте phpMyAdmin: https://localhost:1500/
2. Войдите с учетными данными:
   - **Пользователь:** `u3424187_root_avtovyshkin`
   - **Пароль:** `JavaScript6315`
   - **База данных:** `u3424187_avtovyshkin`

### Шаг 2: Создание таблиц

1. В phpMyAdmin выберите базу данных `u3424187_avtovyshkin`
2. Нажмите вкладку "SQL"
3. Скопируйте и выполните содержимое файла `sql/create_drivers_table_regru.sql`
4. Скопируйте и выполните содержимое файла `sql/create_vehicles_table_regru.sql`

### Шаг 3: Импорт данных

**Через phpMyAdmin:**

1. Для таблицы `drivers`:
   - Выберите таблицу `drivers`
   - Нажмите "Import" (Импорт)
   - Выберите файл `drivers_export.csv`
   - Настройки импорта:
     - **Format:** CSV
     - **Columns separated by:** `,`
     - **Columns enclosed by:** `"`
     - **Columns escaped by:** `\`
     - **Lines terminated by:** `AUTO`
   - Нажмите "Go" (Выполнить)

2. Для таблицы `vehicles`:
   - Выберите таблицу `vehicles`
   - Нажмите "Import" (Импорт)
   - Выберите файл `vehicles_export.csv`
   - Те же настройки импорта
   - Нажмите "Go" (Выполнить)

### Шаг 4: Проверка данных

Выполните следующие SQL запросы для проверки:

```sql
-- Проверка количества записей
SELECT 'Drivers count:' as info, COUNT(*) as count FROM drivers
UNION ALL
SELECT 'Vehicles count:', COUNT(*) FROM vehicles;

-- Проверка данных drivers
SELECT * FROM drivers LIMIT 3;

-- Проверка данных vehicles
SELECT * FROM vehicles LIMIT 3;
```

## 🚀 Запуск приложения с REG.RU базой данных

### Вариант 1: Локальная разработка с REG.RU базой

```powershell
# Drivers microservice
cd backend/drivers-microservice
Copy-Item .env.production .env -Force
go run cmd/server/main.go

# Vehicles microservice (в новом терминале)
cd backend/vehicles-microservice
Copy-Item .env.production .env -Force
go run cmd/server/main.go

# Frontend (в новом терминале)
cd frontend
Copy-Item .env.production .env -Force
yarn dev
```

### Вариант 2: Production на REG.RU

1. Загрузите файлы проекта на сервер REG.RU
2. Установите зависимости и соберите проект:

```bash
# Drivers microservice
cd backend/drivers-microservice
go mod download
go build -o drivers-server cmd/server/main.go

# Vehicles microservice
cd ../vehicles-microservice
go mod download
go build -o vehicles-server cmd/server/main.go

# Frontend
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

## 🔧 Настройка веб-сервера (Nginx)

Создайте конфигурацию Nginx для проксирования:

```nginx
server {
    listen 80;
    server_name avtovyshkin.ru www.avtovyshkin.ru;

    # Frontend
    location / {
        root /path/to/frontend/dist;
        try_files $uri $uri/ /index.html;
    }

    # Drivers API
    location /api/drivers/ {
        proxy_pass http://localhost:8082/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Vehicles API
    location /api/vehicles/ {
        proxy_pass http://localhost:8081/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # WebSocket для drivers
    location /ws/drivers/ {
        proxy_pass http://localhost:8082/ws/;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
    }

    # WebSocket для vehicles
    location /ws/vehicles/ {
        proxy_pass http://localhost:8081/ws/;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
    }
}
```

## 📝 Проверка подключения

```powershell
# Тест подключения к MySQL на REG.RU
Test-NetConnection -ComputerName localhost -Port 3306
```

## ⚠️ Возможные проблемы и решения

### Проблема: Кодировка при импорте

**Решение:** Убедитесь, что при импорте в phpMyAdmin выбрана кодировка UTF-8.

### Проблема: Ошибка подключения к базе

**Решение:**
1. Проверьте правильность учетных данных
2. Убедитесь, что ваш IP адрес добавлен в список разрешенных в панели REG.RU
3. Проверьте, что порт 3306 открыт

### Проблема: JSON поля в vehicles

**Решение:** MySQL использует тип `JSON` вместо `JSONB`. Если возникнут проблемы:

```sql
-- Проверка JSON поля
SELECT fldImgArray FROM vehicles LIMIT 1;

-- Исправление пустых JSON полей
UPDATE vehicles SET fldImgArray = '[]' WHERE fldImgArray IS NULL OR fldImgArray = '';
```

## 🔄 Откат изменений

Если нужно вернуться на PostgreSQL:

1. Восстановите бекап: `backup_postgres_20260420_170319.sql`
2. Используйте локальные `.env` файлы вместо `.env.production`
3. Запустите локальный PostgreSQL контейнер

## ✅ Завершение

После успешного выполнения всех шагов:

- ✅ Данные перенесены в MySQL на REG.RU
- ✅ Таблицы созданы и проиндексированы
- ✅ Приложение настроено для работы с новой базой
- ✅ Бекап PostgreSQL сохранен для отката

**Проект готов к работе на REG.RU!** 🎉
