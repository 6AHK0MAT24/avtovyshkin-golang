# Режим разработки с Hot Reload

## 🚀 Быстрый старт

### 1. Запуск PostgreSQL в Docker
```powershell
.\start-dev.ps1
```

### 2. Запуск Backend с Hot Reload (в новом терминале)
```powershell
.\start-backend.ps1
```

### 3. Запуск Frontend (в новом терминале)
```powershell
.\start-frontend.ps1
```

## ✨ Преимущества

- ✅ **PostgreSQL в Docker** - изолированная база данных, легко управлять
- ✅ **Hot Reload для Backend** - изменения в Go коде применяются автоматически
- ✅ **Hot Reload для Frontend** - изменения в React коде видны мгновенно
- ✅ **Быстрый цикл разработки** - без пересборки Docker контейнеров
- ✅ **Полная отладка** - доступ к логам и отладчику

## 📝 Как это работает

### Backend Hot Reload
Используется инструмент **air** для автоматического перезапуска:
- При изменении `.go` файлов → автоматическая перекомпиляция
- Сервер перезапускается с новыми изменениями
- Занимает ~1-2 секунды

### Frontend Hot Reload
Используется **Vite HMR**:
- При изменении `.tsx/.ts` файлов → мгновенное обновление в браузере
- Сохраняется состояние приложения
- Занимает < 1 секунды

## 🔧 Конфигурация

### Backend (.air.toml)
Конфигурация hot reload находится в `backend/drivers-microservice/.air.toml`:
- `bin` - путь к скомпилированному исполняемому файлу
- `cmd` - команда для сборки
- `include_ext` - расширения файлов для отслеживания
- `exclude_dir` - директории для исключения

### Environment Variables
- `.env` - корневой файл с настройками PostgreSQL
- `backend/drivers-microservice/.env` - настройки бэкенда

## 🛠️ Полезные команды

### PostgreSQL
```powershell
# Запуск
.\start-dev.ps1

# Остановка
.\stop-dev.ps1

# Просмотр логов
docker-compose -f docker-compose.dev.yml logs -f postgres

# Подключение к базе
docker exec -it avtovyshkin-postgres psql -U postgres -d avtovyshkin_db_go
```

### Backend
```powershell
# Запуск с hot reload
.\start-backend.ps1

# Ручной запуск (без hot reload)
cd backend/drivers-microservice
go run cmd/server/main.go

# Сборка
cd backend/drivers-microservice
go build -o tmp/main.exe ./cmd/server/main.go
```

### Frontend
```powershell
# Запуск
.\start-frontend.ps1

# Ручной запуск
cd frontend
yarn dev

# Сборка для production
cd frontend
yarn build
```

## 🐛 Устранение проблем

### Backend не перезапускается
```powershell
# Проверьте, что air установлен
go install github.com/cosmtrek/air@latest

# Проверьте конфигурацию
cd backend/drivers-microservice
air --version
```

### PostgreSQL не подключается
```powershell
# Проверьте статус контейнера
docker ps | findstr postgres

# Проверьте логи
docker-compose -f docker-compose.dev.yml logs postgres

# Перезапустите PostgreSQL
.\stop-dev.ps1
.\start-dev.ps1
```

### Порт уже занят
```powershell
# Для порта 8080
netstat -ano | findstr :8080
taskkill /F /PID <PID>

# Для порта 54320
netstat -ano | findstr :54320
taskkill /F /PID <PID>
```

## 📊 Структура проекта

```
avtovyshkin-golang/
├── docker-compose.dev.yml      # PostgreSQL для разработки
├── docker-compose.yml          # Полная конфигурация (production)
├── start-dev.ps1              # Запуск PostgreSQL
├── stop-dev.ps1               # Остановка PostgreSQL
├── start-backend.ps1          # Запуск backend с hot reload
├── start-frontend.ps1         # Запуск frontend
├── .env                       # Настройки PostgreSQL
├── backend/
│   └── drivers-microservice/
│       ├── .air.toml          # Конфигурация hot reload
│       ├── .env               # Настройки backend
│       └── ...
└── frontend/
    └── ...
```

## 🎯 Советы по разработке

1. **Используйте 3 терминала**:
   - Терминал 1: PostgreSQL (можно закрыть после запуска)
   - Терминал 2: Backend с hot reload
   - Терминал 3: Frontend

2. **Следите за логами**:
   - Backend: логи в терминале с air
   - Frontend: логи в терминале с yarn dev
   - PostgreSQL: `docker-compose -f docker-compose.dev.yml logs -f postgres`

3. **Тестируйте изменения**:
   - Backend: сохраните файл → подождите 1-2 сек → проверьте API
   - Frontend: сохраните файл → обновление в браузере автоматически

4. **Используйте .gitignore**:
   - `backend/drivers-microservice/tmp/` - временные файлы air
   - `backend/drivers-microservice/build-errors.log` - логи сборки

## 📚 Дополнительные ресурсы

- [Air Documentation](https://github.com/cosmtrek/air)
- [Vite HMR](https://vitejs.dev/guide/features.html#hot-module-replacement)
- [Docker Compose](https://docs.docker.com/compose/)
