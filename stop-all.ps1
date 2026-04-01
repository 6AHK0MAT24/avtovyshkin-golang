# Скрипт для остановки всех сервисов

Write-Host "Остановка всех сервисов..." -ForegroundColor Yellow

# Останавливаем процессы на портах 8080 и 5173
$ports = @("8080", "5173")

foreach ($port in $ports) {
    $portInUse = netstat -ano | Select-String ":$port" | Select-String "LISTENING"
    if ($portInUse) {
        Write-Host "Остановка процесса на порту $port..." -ForegroundColor Cyan
        $pid = ($portInUse -split '\s+')[-1]
        try {
            taskkill /F /PID $pid | Out-Null
            Write-Host "Процесс на порту $port остановлен." -ForegroundColor Green
        } catch {
            Write-Host "Не удалось остановить процесс на порту $port." -ForegroundColor Red
        }
    } else {
        Write-Host "Порт $port не используется." -ForegroundColor Gray
    }
}

# Спрашиваем, нужно ли останавливать PostgreSQL
$stopPostgres = Read-Host "Остановить PostgreSQL? (y/n)"
if ($stopPostgres -eq "y" -or $stopPostgres -eq "Y") {
    Write-Host "Остановка PostgreSQL..." -ForegroundColor Cyan
    docker-compose down
    Write-Host "PostgreSQL остановлен." -ForegroundColor Green
} else {
    Write-Host "PostgreSQL продолжает работать." -ForegroundColor Yellow
}

Write-Host ""
Write-Host "Все сервисы остановлены." -ForegroundColor Green
