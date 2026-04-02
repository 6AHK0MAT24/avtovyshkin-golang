# Script to start frontend application

Write-Host "Starting Frontend application..." -ForegroundColor Green

# Change to frontend directory
$frontendDir = "frontend"
if (-not (Test-Path $frontendDir)) {
    Write-Host "Error: Directory $frontendDir not found!" -ForegroundColor Red
    exit 1
}

Set-Location $frontendDir

# Check if dependencies are installed
if (-not (Test-Path "node_modules")) {
    Write-Host "Installing dependencies..." -ForegroundColor Yellow
    yarn install
}

# Check port 5173
$portInUse = netstat -ano | Select-String ":5173" | Select-String "LISTENING"
if ($portInUse) {
    Write-Host "Warning: Port 5173 is already in use. Attempting to free it..." -ForegroundColor Yellow
    $pid = ($portInUse -split '\s+')[-1]
    try {
        taskkill /F /PID $pid | Out-Null
        Write-Host "Port 5173 freed." -ForegroundColor Green
    } catch {
        Write-Host "Failed to free port 5173. Please stop the process manually." -ForegroundColor Red
        exit 1
    }
}

# Start frontend
Write-Host "Starting dev server on port 5173..." -ForegroundColor Cyan
Write-Host "Frontend will be available at: http://localhost:5173" -ForegroundColor Green
yarn dev
