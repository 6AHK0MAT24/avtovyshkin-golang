# Скрипты для запуска проекта

Для удобного запуска проекта локально созданы следующие PowerShell скрипты.

## 🚀 Режимы запуска

### Рекомендуемый режим разработки (Hot Reload)
PostgreSQL в Docker + Backend локально с hot reload + Frontend локально.

### Режим разработки (Development)
Frontend запускается локально через `yarn dev` с HMR и полной отладкой.

### Режим продакшена (Production)
Frontend запускается в Docker с nginx и оптимизированной статикой.

## 📋 Доступные скрипты

### `start-dev.ps1` - Запуск PostgreSQL в Docker (Рекомендуемый режим)
Запускает только PostgreSQL в Docker для локальной разработки.

```powershell
.\start-dev.ps1
```

**После запуска:**
- PostgreSQL: localhost:54320
- База данных: avtovyshkin_db_go

**Далее запустите в отдельных терминалах:**
```powershell
# Backend с hot reload
.\start-backend.ps1

# Frontend
.\start-frontend.ps1
```

### `start-backend.ps1` - Запуск Backend с Hot Reload
Запускает backend сервис локально с автоматическим перезапуском при изменениях.

```powershell
.\start-backend.ps1
```

**Особенности:**
- ✅ Автоматическая перекомпиляция при изменении файлов
- ✅ Быстрый цикл разработки
- ✅ Полный доступ к отладке

### `stop-dev.ps1` - Остановка PostgreSQL
Останавливает PostgreSQL контейнер.

```powershell
.\stop-dev.ps1
```

### `start-all.ps1` - Запуск всего проекта (Development)
Запускает PostgreSQL, Backend (Docker) и Frontend (локально) в отдельных окнах.

```powershell
.\start-all.ps1
```

**После запуска:**
- Backend: http://localhost:8080
- Frontend: http://localhost:5173

### `start-frontend.ps1` - Запуск только Frontend
Запускает frontend приложение локально через `yarn dev` на порту 5173.

```powershell
.\start-frontend.ps1
```

### `stop-all.ps1` - Остановка всех сервисов
Останавливает процессы на портах 8080 и 5173, опционально останавливает PostgreSQL.

```powershell
.\stop-all.ps1
```

## 🐳 Docker Compose команды

### Development (только postgres)
```powershell
# Запуск только postgres
docker-compose -f docker-compose.dev.yml up -d

# Остановка
docker-compose -f docker-compose.dev.yml down
```

### Development (backend + postgres)
```powershell
# Запуск backend и postgres
docker-compose up -d

# Остановка
docker-compose down
```

### Production (включая frontend)
```powershell
# Запуск всех сервисов включая frontend
docker-compose --profile production up -d

# Остановка
docker-compose --profile production down
```

## 🔧 Особенности скриптов

### Автоматическая проверка PostgreSQL
Скрипты автоматически проверяют, запущен ли PostgreSQL, и запускают его при необходимости.

### Hot Reload для Backend
Скрипт `start-backend.ps1` использует `air` для автоматического перезапуска при изменениях:
- Изменения в Go файлах → автоматическая перекомпиляция
- Изменения применяются без ручного перезапуска
- Быстрый цикл разработки

### Освобождение портов
Если порты 8080 или 5173 уже заняты, скрипты попытаются освободить их автоматически.

### Установка зависимостей
Скрипт `start-frontend.ps1` автоматически установит зависимости, если `node_modules` не существует.

## 📝 Примеры использования

### Рекомендуемый цикл работы (Hot Reload)

1. **Запуск PostgreSQL:**
   ```powershell
   .\start-dev.ps1
   ```

2. **Запуск Backend с hot reload (в новом терминале):**
   ```powershell
   .\start-backend.ps1
   ```

3. **Запуск Frontend (в новом терминале):**
   ```powershell
   .\start-frontend.ps1
   ```

4. **Разработка:**
   - Вносите изменения в код бэкенда → автоматический перезапуск
   - Вносите изменения в код фронтенда → мгновенное обновление в браузере

5. **Остановка PostgreSQL:**
   ```powershell
   .\stop-dev.ps1
   ```

### Полный цикл работы (Development)

1. **Запуск всего проекта:**
   ```powershell
   .\start-all.ps1
   ```

2. **Работа с проектом** (откроются два окна PowerShell)
   - Frontend с HMR - изменения видны мгновенно
   - Backend в Docker

3. **Остановка всех сервисов:**
   ```powershell
   .\stop-all.ps1
   ```

### Раздельный запуск

Если нужно запустить только часть проекта:

```powershell
# Только backend
.\start-backend.ps1

# Только frontend (в другом окне)
.\start-frontend.ps1
```

### Production режим

Для развертывания в production:

```powershell
# Запуск всех сервисов в Docker
docker-compose --profile production up -d

# Проверка статуса
docker-compose --profile production ps
```

## ⚠️ Важно

- **Рекомендуемый режим**: PostgreSQL в Docker + Backend локально с hot reload
- **Development**: Frontend запускается локально через `yarn dev` с HMR
- **Production**: Frontend запускается в Docker с nginx
- Для работы скриптов необходим PowerShell
- Убедитесь, что Docker и Docker Compose установлены
- При первом запуске frontend скрипт установит зависимости (может занять время)
- Для остановки PostgreSQL используйте `.\stop-dev.ps1` или `docker-compose -f docker-compose.dev.yml down`

## 🐛 Устранение проблем

### Порт уже используется
Если скрипт не может освободить порт, остановите процесс вручную:
```powershell
# Для порта 8080
netstat -ano | findstr :8080
taskkill /F /PID <PID>

# Для порта 5173
netstat -ano | findstr :5173
taskkill /F /PID <PID>
```

### PostgreSQL не запускается
```powershell
docker-compose -f docker-compose.dev.yml logs postgres
docker-compose -f docker-compose.dev.yml down
docker-compose -f docker-compose.dev.yml up -d postgres
```

### Ошибка выполнения скрипта
Если PowerShell блокирует выполнение скриптов:
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

### Air не установлен
Если `air` не установлен, скрипт `start-backend.ps1` установит его автоматически. Или установите вручную:
```powershell
go install github.com/cosmtrek/air@latest
```

### Frontend в Docker не запускается
Frontend по умолчанию не запускается в Docker для разработки. Используйте:
```powershell
# Для development
cd frontend
yarn dev

# Для production
docker-compose --profile production up -d frontend
```

**После запуска:**
- Backend: http://localhost:8080
- Frontend: http://localhost:5173


**После запуска:**
- Backend: http://localhost:8080
- Frontend: http://localhost:5173

### `start-backend.ps1` - Запуск только Backend
Запускает backend сервис в Docker на порту 8080.

```powershell
.\start-backend.ps1
```

### `start-frontend.ps1` - Запуск только Frontend
Запускает frontend приложение локально через `yarn dev` на порту 5173.

```powershell
.\start-frontend.ps1
```

### `stop-all.ps1` - Остановка всех сервисов
Останавливает процессы на портах 8080 и 5173, опционально останавливает PostgreSQL.

```powershell
.\stop-all.ps1
```

## 🐳 Docker Compose команды

### Development (только backend + postgres)
```powershell
# Запуск backend и postgres
docker-compose up -d

# Остановка
docker-compose down
```

### Production (включая frontend)
```powershell
# Запуск всех сервисов включая frontend
docker-compose --profile production up -d

# Остановка
docker-compose --profile production down
```

## 🔧 Особенности скриптов

### Автоматическая проверка PostgreSQL
Скрипты автоматически проверяют, запущен ли PostgreSQL, и запускают его при необходимости.

### Освобождение портов
Если порты 8080 или 5173 уже заняты, скрипты попытаются освободить их автоматически.

### Установка зависимостей
Скрипт `start-frontend.ps1` автоматически установит зависимости, если `node_modules` не существует.

## 📝 Примеры использования

### Полный цикл работы (Development)

1. **Запуск всего проекта:**
   ```powershell
   .\start-all.ps1
   ```

2. **Работа с проектом** (откроются два окна PowerShell)
   - Frontend с HMR - изменения видны мгновенно
   - Backend в Docker

3. **Остановка всех сервисов:**
   ```powershell
   .\stop-all.ps1
   ```

### Раздельный запуск

Если нужно запустить только часть проекта:

```powershell
# Только backend
.\start-backend.ps1

# Только frontend (в другом окне)
.\start-frontend.ps1
```

### Production режим

Для развертывания в production:

```powershell
# Запуск всех сервисов в Docker
docker-compose --profile production up -d

# Проверка статуса
docker-compose --profile production ps
```

## ⚠️ Важно

- **Development**: Frontend запускается локально через `yarn dev` с HMR
- **Production**: Frontend запускается в Docker с nginx
- Для работы скриптов необходим PowerShell
- Убедитесь, что Docker и Docker Compose установлены
- При первом запуске frontend скрипт установит зависимости (может занять время)
- Для остановки PostgreSQL используйте `docker-compose down`

## 🐛 Устранение проблем

### Порт уже используется
Если скрипт не может освободить порт, остановите процесс вручную:
```powershell
# Для порта 8080
netstat -ano | findstr :8080
taskkill /F /PID <PID>

# Для порта 5173
netstat -ano | findstr :5173
taskkill /F /PID <PID>
```

### PostgreSQL не запускается
```powershell
docker-compose logs postgres
docker-compose down
docker-compose up -d postgres
```

### Ошибка выполнения скрипта
Если PowerShell блокирует выполнение скриптов:
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

### Frontend в Docker не запускается
Frontend по умолчанию не запускается в Docker для разработки. Используйте:
```powershell
# Для development
cd frontend
yarn dev

# Для production
docker-compose --profile production up -d frontend
```

**После запуска:**
- Backend: http://localhost:8080
- Frontend: http://localhost:5173

### `start-backend.ps1` - Запуск только Backend
Запускает backend сервис на порту 8080.

```powershell
.\start-backend.ps1
```

### `start-frontend.ps1` - Запуск только Frontend
Запускает frontend приложение на порту 5173.

```powershell
.\start-frontend.ps1
```

### `stop-all.ps1` - Остановка всех сервисов
Останавливает процессы на портах 8080 и 5173, опционально останавливает PostgreSQL.

```powershell
.\stop-all.ps1
```

## 🔧 Особенности скриптов

### Автоматическая проверка PostgreSQL
Скрипты автоматически проверяют, запущен ли PostgreSQL, и запускают его при необходимости.

### Освобождение портов
Если порты 8080 или 5173 уже заняты, скрипты попытаются освободить их автоматически.

### Установка зависимостей
Скрипт `start-frontend.ps1` автоматически установит зависимости, если `node_modules` не существует.

## 📝 Примеры использования

### Полный цикл работы

1. **Запуск всего проекта:**
   ```powershell
   .\start-all.ps1
   ```

2. **Работа с проектом** (откроются два окна PowerShell)

3. **Остановка всех сервисов:**
   ```powershell
   .\stop-all.ps1
   ```

### Раздельный запуск

Если нужно запустить только часть проекта:

```powershell
# Только backend
.\start-backend.ps1

# Только frontend (в другом окне)
.\start-frontend.ps1
```

## ⚠️ Важно

- Для работы скриптов необходим PowerShell
- Убедитесь, что Docker и Docker Compose установлены
- При первом запуске frontend скрипт установит зависимости (может занять время)
- Для остановки PostgreSQL используйте `docker-compose down`

## 🐛 Устранение проблем

### Порт уже используется
Если скрипт не может освободить порт, остановите процесс вручную:
```powershell
# Для порта 8080
netstat -ano | findstr :8080
taskkill /F /PID <PID>

# Для порта 5173
netstat -ano | findstr :5173
taskkill /F /PID <PID>
```

### PostgreSQL не запускается
```powershell
docker-compose logs postgres
docker-compose down
docker-compose up -d postgres
```

### Ошибка выполнения скрипта
Если PowerShell блокирует выполнение скриптов:
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```
