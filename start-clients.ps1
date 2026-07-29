# Script to start the Clients Microservice

# Set window title
$Host.UI.RawUI.WindowTitle = "Clients Microservice"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Starting Clients Microservice" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Navigate to the clients microservice directory
$clientsDir = Join-Path $PSScriptRoot "backend\clients-microservice"
Set-Location $clientsDir

Write-Host "Working directory: $clientsDir" -ForegroundColor Yellow
Write-Host ""

# Check if go.mod exists
if (-not (Test-Path "go.mod")) {
    Write-Host "Error: go.mod not found in $clientsDir" -ForegroundColor Red
    Write-Host "Please run 'go mod init clients-service' first" -ForegroundColor Yellow
    exit 1
}

# Run the service
Write-Host "Starting Clients Microservice on port 8003..." -ForegroundColor Green
Write-Host ""
Write-Host "API Endpoints:" -ForegroundColor Cyan
Write-Host "  GET    /api/clients              - Get all clients" -ForegroundColor White
Write-Host "  GET    /api/clients/type/{type}   - Get clients by type (individual/legal)" -ForegroundColor White
Write-Host "  GET    /api/clients/{id}          - Get client by ID" -ForegroundColor White
Write-Host "  POST   /api/clients              - Create new client" -ForegroundColor White
Write-Host "  PUT    /api/clients/{id}          - Update client" -ForegroundColor White
Write-Host "  DELETE /api/clients/{id}          - Delete client" -ForegroundColor White
Write-Host "  GET    /api/clients/search        - Search clients" -ForegroundColor White
Write-Host "  GET    /health                    - Health check" -ForegroundColor White
Write-Host ""
Write-Host "Press Ctrl+C to stop the service" -ForegroundColor Yellow
Write-Host ""

go run cmd/server/main.go
