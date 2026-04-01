# Скрипт для запуска всего проекта (backend + frontend)

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Запуск проекта Автовышкин" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Проверяем, запущен ли PostgreSQL
$postgresRunning = docker ps --filter "name=avtovyshkin-postgres" --format "{{.Names}}"
if (-not $postgresRunning) {
    Write-Host "[1/3] Запуск PostgreSQL..." -ForegroundColor Yellow
    docker-compose up -d postgres
    Start-Sleep -Seconds 3
    Write-Host "PostgreSQL запущен." -ForegroundColor Green
} else {
    Write-Host "[1/3] PostgreSQL уже запущен." -ForegroundColor Green
}

# Запускаем backend в новом окне
Write-Host "[2/3] Запуск Backend..." -ForegroundColor Yellow
$backendScript = Join-Path $PSScriptRoot "start-backend.ps1"
Start-Process powershell -ArgumentList "-NoExit", "-File", $backendScript

# Ждем немного, чтобы backend успел запуститься
Start-Sleep -Seconds 3

# Запускаем frontend в новом окне
Write-Host "[3/3] Запуск Frontend..." -ForegroundColor Yellow
$frontendScript = Join-Path $PSScriptRoot "start-frontend.ps1"
Start-Process powershell -ArgumentList "-NoExit", "-File", $frontendScript

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Проект запущен!" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Backend:  http://localhost:8080" -ForegroundColor White
Write-Host "Frontend: http://localhost:5173" -ForegroundColor White
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Для остановки закройте окна с backend и frontend." -ForegroundColor Yellow
Write-Host "Для остановки PostgreSQL выполните: docker-compose down" -ForegroundColor Yellow
