# Скрипт для запуска backend сервиса

Write-Host "Запуск Drivers Service..." -ForegroundColor Green

# Проверяем, запущен ли PostgreSQL
$postgresRunning = docker ps --filter "name=avtovyshkin-postgres" --format "{{.Names}}"
if (-not $postgresRunning) {
    Write-Host "PostgreSQL не запущен. Запускаем..." -ForegroundColor Yellow
    docker-compose up -d postgres
    Start-Sleep -Seconds 3
}

# Переходим в директорию backend
$backendDir = "backend/drivers-microservice"
if (-not (Test-Path $backendDir)) {
    Write-Host "Ошибка: Директория $backendDir не найдена!" -ForegroundColor Red
    exit 1
}
if (-not (Test-Path $backendDir)) {
    Write-Host "Ошибка: Директория $backendDir не найдена!" -ForegroundColor Red
    exit 1
}


Set-Location $backendDir

# Проверяем порт 8080
$portInUse = netstat -ano | Select-String ":8080" | Select-String "LISTENING"
if ($portInUse) {
    Write-Host "Внимание: Порт 8080 уже используется. Попытка освободить..." -ForegroundColor Yellow
    $pid = ($portInUse -split '\s+')[-1]
    try {
        taskkill /F /PID $pid | Out-Null
        Write-Host "Порт 8080 освобожден." -ForegroundColor Green
    } catch {
        Write-Host "Не удалось освободить порт 8080. Пожалуйста, остановите процесс вручную." -ForegroundColor Red
        exit 1
    }
}

# Запускаем backend
Write-Host "Запуск сервера на порту 8080..." -ForegroundColor Cyan
go run cmd/server/main.go
go run cmd/server/main.go
