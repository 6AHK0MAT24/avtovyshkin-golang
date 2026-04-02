# Script to stop all services

Write-Host "Stopping all services..." -ForegroundColor Yellow

# Stop processes on ports 8080 and 5173
$ports = @("8080", "5173")

foreach ($port in $ports) {
    $portInUse = netstat -ano | Select-String ":$port" | Select-String "LISTENING"
    if ($portInUse) {
        Write-Host "Stopping process on port $port..." -ForegroundColor Cyan
        $pid = ($portInUse -split '\s+')[-1]
        try {
            taskkill /F /PID $pid | Out-Null
            Write-Host "Process on port $port stopped." -ForegroundColor Green
        } catch {
            Write-Host "Failed to stop process on port $port." -ForegroundColor Red
        }
    } else {
        Write-Host "Port $port is not in use." -ForegroundColor Gray
    }
}

# Ask if PostgreSQL should be stopped
$stopPostgres = Read-Host "Stop PostgreSQL? (y/n)"
if ($stopPostgres -eq "y" -or $stopPostgres -eq "Y") {
    Write-Host "Stopping PostgreSQL..." -ForegroundColor Cyan
    docker-compose down
    Write-Host "PostgreSQL stopped." -ForegroundColor Green
} else {
    Write-Host "PostgreSQL continues running." -ForegroundColor Yellow
}

Write-Host ""
Write-Host "All services stopped." -ForegroundColor Green
