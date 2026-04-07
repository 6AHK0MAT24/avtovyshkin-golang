# Script to start the entire project (backend + frontend)

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Starting Avtovyshkin Project" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Check if PostgreSQL is running
$postgresRunning = docker ps --filter "name=avtovyshkin-postgres" --format "{{.Names}}"
if (-not $postgresRunning) {
    Write-Host "[1/3] Starting PostgreSQL..." -ForegroundColor Yellow
    docker-compose up -d postgres
    Start-Sleep -Seconds 3
    Write-Host "PostgreSQL started." -ForegroundColor Green
} else {
    Write-Host "[1/3] PostgreSQL is already running." -ForegroundColor Green
}

# Start backend in new window
Write-Host "[2/4] Starting Backend..." -ForegroundColor Yellow
$backendScript = Join-Path $PSScriptRoot "start-backend.ps1"
Start-Process powershell -ArgumentList "-NoExit", "-File", $backendScript

# Wait a bit for backend to start
Start-Sleep -Seconds 3

# Start vehicles microservice in new window
Write-Host "[3/4] Starting Vehicles Microservice..." -ForegroundColor Yellow
$vehiclesScript = Join-Path $PSScriptRoot "start-vehicles.ps1"
Start-Process powershell -ArgumentList "-NoExit", "-File", $vehiclesScript

# Wait a bit for vehicles microservice to start
Start-Sleep -Seconds 3

# Start frontend in new window
Write-Host "[4/4] Starting Frontend..." -ForegroundColor Yellow
$frontendScript = Join-Path $PSScriptRoot "start-frontend.ps1"
Start-Process powershell -ArgumentList "-NoExit", "-File", $frontendScript

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Project started!" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Backend:           http://localhost:8080" -ForegroundColor White
Write-Host "Vehicles Service:  http://localhost:8081" -ForegroundColor White
Write-Host "Frontend:          http://localhost:5173" -ForegroundColor White
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "To stop, close the backend and frontend windows." -ForegroundColor Yellow
Write-Host "To stop PostgreSQL run: docker-compose down" -ForegroundColor Yellow
