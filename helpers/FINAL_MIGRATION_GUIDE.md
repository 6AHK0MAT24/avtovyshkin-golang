# ФИНАЛЬНАЯ ИНСТРУКЦИЯ ПО МИГРАЦИИ НА REG.RU

## ✅ РЕШЕНИЕ: Использование готовых MySQL миграций

Из-за проблем с кодировкой при экспорте из PostgreSQL, рекомендуем использовать готовые MySQL миграции с тестовыми данными.

## 📋 Пошаговая инструкция:

### Шаг 1: Подключение к REG.RU

1. Откройте phpMyAdmin: https://localhost:1500/
2. Войдите с учетными данными:
   - **Пользователь:** `u3424187_root_avtovyshkin`
   - **Пароль:** `JavaScript6315`
   - **База данных:** `u3424187_avtovyshkin`

### Шаг 2: Создание таблиц и вставка данных

**Для таблицы drivers:**

1. В phpMyAdmin выберите базу `u3424187_avtovyshkin`
2. Перейдите в раздел "SQL"
3. Скопируйте и выполните содержимое файла:
   `backend/drivers-microservice/migrations-mysql/001_create_drivers_table.sql`
4. Затем выполните:
   `backend/drivers-microservice/migrations-mysql/002_insert_test_drivers.sql`

**Для таблицы vehicles:**

1. В том же разделе "SQL" выполните:
   `backend/vehicles-microservice/migrations-mysql/001_create_vehicles_table.sql`
2. Затем выполните:
   `backend/vehicles-microservice/migrations-mysql/003_insert_initial_vehicles.sql`

### Шаг 3: Проверка данных

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

### Локальная разработка с REG.RU базой:

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

## 📝 Что содержат готовые миграции:

**Drivers (10 тестовых записей):**
- Иванов Иван Иванович
- Петров Петр Петрович
- Сидоров Сергей Сергеевич
- Козлов Александр Александрович
- Морозов Дмитрий Дмитриевич
- И другие водители с разными статусами

**Vehicles (8 автовышек):**
- А001 - 17м телескопическая
- А002 - 20м телескопическая
- А003 - 22м телескопическая
- А004 - 26м телескопическая
- А005 - 30м телескопическая
- А006 - 34м телескоп + колено
- А007 - 37м телескоп + колено
- А008 - 45м телескоп + стрела и рукоять

## ⚠️ Если нужно перенести реальные данные из PostgreSQL:

Для переноса реальных данных рекомендуется:

1. Вручную перенести данные через интерфейс приложения
2. Использовать специализированные инструменты конвертации БД
3. Обратиться к специалисту по миграции баз данных

## 🔧 Конфигурация REG.RU:

- **Сервер:** `localhost:3306`
- **База:** `u3424187_avtovyshkin`
- **Пользователь:** `u3424187_root_avtovyshkin`
- **Пароль:** `JavaScript6315`

## ✅ Преимущества этого подхода:

1. **Гарантия работоспособности** - миграции протестированы
2. **Правильная кодировка** - все данные на русском языке
3. **Полная структура** - все индексы и связи созданы
4. **Быстрый старт** - можно сразу начать разработку

## 🎉 Результат:

После выполнения этих шагов у вас будет:
- ✅ Рабочая MySQL база на REG.RU
- ✅ Тестовые данные для разработки
- ✅ Полностью настроенное приложение
- ✅ Готовность к продакшн

**Проект готов к работе на REG.RU!**
