# Script to start drivers service with hot reload

Write-Host "Starting Drivers Service with hot reload..." -ForegroundColor Green
# Check if PostgreSQL is running
$postgresRunning = docker ps --filter "name=avtovyshkin-postgres" --format "{{.Names}}"
if (-not $postgresRunning) {
    Write-Host "PostgreSQL is not running. Starting..." -ForegroundColor Yellow
    docker-compose -f docker-compose.dev.yml up -d postgres
    Start-Sleep -Seconds 3
}

# Change to drivers service directory
$backendDir = "backend/drivers-microservice"
if (-not (Test-Path $backendDir)) {
    Write-Host "Error: Directory $backendDir not found!" -ForegroundColor Red
    exit 1
}
Set-Location $backendDir

# Check port 8080
$portInUse = netstat -ano | Select-String ":8080" | Select-String "LISTENING"
if ($portInUse) {
    Write-Host "Warning: Port 8080 is already in use. Attempting to free it..." -ForegroundColor Yellow
    $pid = ($portInUse -split '\s+')[-1]
    try {
        taskkill /F /PID $pid | Out-Null
        Write-Host "Port 8080 freed." -ForegroundColor Green
    } catch {
        Write-Host "Failed to free port 8080. Please stop the process manually." -ForegroundColor Red
        exit 1
    }
}

# Check if air is installed
$airPath = Join-Path $env:USERPROFILE "go\bin\air.exe"
if (-not (Test-Path $airPath)) {
    Write-Host "air is not installed. Installing..." -ForegroundColor Yellow
    go install github.com/air-verse/air@latest
    Write-Host "air installed successfully." -ForegroundColor Green
}

# Start drivers service with hot reload
Write-Host "Starting drivers service with hot reload on port 8080..." -ForegroundColor Cyan
Write-Host "Code changes will be applied automatically." -ForegroundColor Cyan
& $airPath
