# Быстрый старт - Режим разработки с Hot Reload

## 🚀 Запуск за 3 шага

### Шаг 1: Запуск PostgreSQL в Docker
```powershell
.\start-dev.ps1
```

### Шаг 2: Запуск Backend с Hot Reload (в новом терминале)
```powershell
.\start-backend.ps1
```

### Шаг 3: Запуск Frontend (в новом терминале)
```powershell
.\start-frontend.ps1
```

## ✅ Готово!

Откройте в браузере:
- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8082
- **PostgreSQL**: localhost:54320

## 🔄 Как работает Hot Reload

### Backend
- Измените любой `.go` файл
- Подождите 1-2 секунды
- Сервер автоматически перезапустится
- Изменения применены!

### Frontend
- Измените любой `.tsx` или `.ts` файл
- Браузер автоматически обновится
- Изменения видны мгновенно!

## 🛑 Остановка

```powershell
# Остановить PostgreSQL
.\stop-dev.ps1

# Остановить backend и frontend - нажмите Ctrl+C в соответствующих терминалах
```

## 📚 Дополнительная информация

Подробнее о режиме разработки см. [DEV_MODE.md](DEV_MODE.md)
