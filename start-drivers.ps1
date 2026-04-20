# Script to start drivers service with hot reload

# Set window title
$Host.UI.RawUI.WindowTitle = "Drivers Service"

Write-Host "Starting Drivers Service with hot reload..." -ForegroundColor Green
Write-Host "Using MySQL database on localhost" -ForegroundColor Cyan

# Change to drivers service directory
$backendDir = "backend/drivers-microservice"

if (-not (Test-Path $backendDir)) {    Write-Host "Error: Directory $backendDir not found!" -ForegroundColor Red
    exit 1
}
Set-Location $backendDir

# Check port 8082
$portInUse = netstat -ano | Select-String ":8082" | Select-String "LISTENING"
if ($portInUse) {
    Write-Host "Warning: Port 8082 is already in use. Attempting to free it..." -ForegroundColor Yellow
    $pid = ($portInUse -split '\s+')[-1]
    try {
        taskkill /F /PID $pid | Out-Null
        Write-Host "Port 8082 freed." -ForegroundColor Green
    } catch {
        Write-Host "Failed to free port 8082. Please stop the process manually." -ForegroundColor Red
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
Write-Host "Starting drivers service with hot reload on port 8082..." -ForegroundColor Cyan
Write-Host "Code changes will be applied automatically." -ForegroundColor Cyan
& $airPath
