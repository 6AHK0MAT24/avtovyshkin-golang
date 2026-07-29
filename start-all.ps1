# Script to start the entire project (backend + frontend)
# Using MySQL on REG.RU instead of local PostgreSQL

# Set window title
$Host.UI.RawUI.WindowTitle = "Avtovyshkin Project - All Services"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Starting Avtovyshkin Project" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Using MySQL (localhost:3306)" -ForegroundColor Yellow
Write-Host ""
# Start drivers service in new window
Write-Host "[1/4] Starting Drivers Service..." -ForegroundColor Yellow
$backendScript = Join-Path $PSScriptRoot "start-drivers.ps1"
Start-Process powershell -ArgumentList "-NoExit", "-File", $backendScript
# Wait a bit for backend to start
Start-Sleep -Seconds 3

# Start vehicles microservice in new window
Write-Host "[2/4] Starting Vehicles Microservice..." -ForegroundColor Yellow
$vehiclesScript = Join-Path $PSScriptRoot "start-vehicles.ps1"
Start-Process powershell -ArgumentList "-NoExit", "-File", $vehiclesScript

# Wait a bit for vehicles microservice to start
Start-Sleep -Seconds 3

# Start clients microservice in new window
Write-Host "[3/4] Starting Clients Microservice..." -ForegroundColor Yellow
$clientsScript = Join-Path $PSScriptRoot "start-clients.ps1"
Start-Process powershell -ArgumentList "-NoExit", "-File", $clientsScript

# Wait a bit for clients microservice to start
Start-Sleep -Seconds 3

# Start frontend in new window
Write-Host "[4/4] Starting Frontend..." -ForegroundColor Yellow
$frontendScript = Join-Path $PSScriptRoot "start-frontend.ps1"
Start-Process powershell -ArgumentList "-NoExit", "-File", $frontendScript

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Project started!" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Drivers Service:   http://localhost:8082" -ForegroundColor White
Write-Host "Vehicles Service:  http://localhost:8081" -ForegroundColor White
Write-Host "Clients Service:   http://localhost:8003" -ForegroundColor White
Write-Host "Frontend:          http://localhost:5173" -ForegroundColor White
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "To stop, close the backend and frontend windows." -ForegroundColor Yellow
