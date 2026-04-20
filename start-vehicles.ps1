# Script to start vehicles microservice with hot reload

# Set window title
$Host.UI.RawUI.WindowTitle = "Vehicles Microservice"

Write-Host "Starting Vehicles Microservice with hot reload..." -ForegroundColor Green
Write-Host "Using MySQL database on localhost:3306" -ForegroundColor Cyan
# Change to vehicles microservice directory
$vehiclesDir = "backend/vehicles-microservice"
if (-not (Test-Path $vehiclesDir)) {
    Write-Host "Error: Directory $vehiclesDir not found!" -ForegroundColor Red
    exit 1
}
Set-Location $vehiclesDir

# Check port 8081
$portInUse = netstat -ano | Select-String ":8081" | Select-String "LISTENING"
if ($portInUse) {
    Write-Host "Warning: Port 8081 is already in use. Attempting to free it..." -ForegroundColor Yellow
    $pid = ($portInUse -split '\s+')[-1]
    try {
        taskkill /F /PID $pid | Out-Null
        Write-Host "Port 8081 freed." -ForegroundColor Green
    } catch {
        Write-Host "Failed to free port 8081. Please stop the process manually." -ForegroundColor Red
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

# Start vehicles microservice with hot reload
Write-Host "Starting vehicles microservice with hot reload on port 8081..." -ForegroundColor Cyan
Write-Host "Code changes will be applied automatically." -ForegroundColor Cyan
& $airPath
