# Script for local development
# Using MySQL on localhost (OpenServer)

# Set window title
$Host.UI.RawUI.WindowTitle = "Development Environment"

Write-Host "Development environment setup" -ForegroundColor Green
Write-Host "Backend services use MySQL on localhost:3306." -ForegroundColor Cyan
Write-Host "`nTo start backend services run: .\start-backend.ps1" -ForegroundColor Yellow
Write-Host "To start frontend run: .\start-frontend.ps1" -ForegroundColor Yellow
