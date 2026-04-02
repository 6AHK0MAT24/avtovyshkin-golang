# Script to start PostgreSQL in Docker for local development

Write-Host "Starting PostgreSQL in Docker..." -ForegroundColor Green

# Check if PostgreSQL is running
$postgresRunning = docker ps --filter "name=avtovyshkin-postgres" --format "{{.Names}}"
if ($postgresRunning) {
    Write-Host "PostgreSQL is already running." -ForegroundColor Yellow
} else {
    Write-Host "Starting PostgreSQL container..." -ForegroundColor Cyan
    docker-compose -f docker-compose.dev.yml up -d postgres

    # Wait for PostgreSQL to be ready
    Write-Host "Waiting for PostgreSQL to be ready..." -ForegroundColor Cyan
    Start-Sleep -Seconds 5

    # Check status
    $postgresStatus = docker ps --filter "name=avtovyshkin-postgres" --format "{{.Status}}"
    if ($postgresStatus) {
        Write-Host "PostgreSQL started successfully: $postgresStatus" -ForegroundColor Green
        Write-Host "Connection: localhost:54320" -ForegroundColor Cyan
        Write-Host "Database: avtovyshkin_db_go" -ForegroundColor Cyan
        Write-Host "User: postgres" -ForegroundColor Cyan
    } else {
        Write-Host "Error starting PostgreSQL. Check logs: docker-compose -f docker-compose.dev.yml logs postgres" -ForegroundColor Red
        exit 1
    }
}

Write-Host "`nTo start backend with hot reload run: .\start-backend.ps1" -ForegroundColor Yellow
Write-Host "To start frontend run: .\start-frontend.ps1" -ForegroundColor Yellow
