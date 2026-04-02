# Script to stop PostgreSQL in Docker

Write-Host "Stopping PostgreSQL..." -ForegroundColor Yellow

# Check if PostgreSQL is running
$postgresRunning = docker ps --filter "name=avtovyshkin-postgres" --format "{{.Names}}"
if ($postgresRunning) {
    docker-compose -f docker-compose.dev.yml down
    Write-Host "PostgreSQL stopped." -ForegroundColor Green
} else {
    Write-Host "PostgreSQL is not running." -ForegroundColor Yellow
}
