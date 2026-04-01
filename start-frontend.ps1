# Скрипт для запуска frontend приложения

Write-Host "Запуск Frontend приложения..." -ForegroundColor Green

# Переходим в директорию frontend
$frontendDir = "frontend"
if (-not (Test-Path $frontendDir)) {
    Write-Host "Ошибка: Директория $frontendDir не найдена!" -ForegroundColor Red
    exit 1
}

Set-Location $frontendDir

# Проверяем, установлены ли зависимости
if (-not (Test-Path "node_modules")) {
    Write-Host "Установка зависимостей..." -ForegroundColor Yellow
    yarn install
}

# Проверяем порт 5173
$portInUse = netstat -ano | Select-String ":5173" | Select-String "LISTENING"
if ($portInUse) {
    Write-Host "Внимание: Порт 5173 уже используется. Попытка освободить..." -ForegroundColor Yellow
    $pid = ($portInUse -split '\s+')[-1]
    try {
        taskkill /F /PID $pid | Out-Null
        Write-Host "Порт 5173 освобожден." -ForegroundColor Green
    } catch {
        Write-Host "Не удалось освободить порт 5173. Пожалуйста, остановите процесс вручную." -ForegroundColor Red
        exit 1
    }
}

# Запускаем frontend
Write-Host "Запуск dev сервера на порту 5173..." -ForegroundColor Cyan
Write-Host "Frontend будет доступен по адресу: http://localhost:5173" -ForegroundColor Green
yarn dev
