# Generate driver photos script
$ErrorActionPreference = "Stop"

$photoDir = "uploads/photo"
$dbName = "u3424187_avtovyshkin"
$dbUser = "u3424187_root_avtovyshkin"
$dbPassword = "JavaScript6315"

Write-Host "=== Generating driver photos ===" -ForegroundColor Green

if (-not (Test-Path $photoDir)) {
    New-Item -ItemType Directory -Path $photoDir -Force | Out-Null
    Write-Host "Created directory: $photoDir" -ForegroundColor Yellow
}

Write-Host "`nThis script is for reference only." -ForegroundColor Yellow
Write-Host "The project now uses MySQL on localhost:3306." -ForegroundColor Cyan

Write-Host "`nTo generate driver photos:" -ForegroundColor Yellow
Write-Host "1. Use MySQL command line or phpMyAdmin" -ForegroundColor White
Write-Host "2. Connect to localhost:3306" -ForegroundColor White
Write-Host "3. Database: $dbName" -ForegroundColor White

Write-Host "`nExample MySQL command:" -ForegroundColor Yellow
Write-Host "mysql -h localhost -u $dbUser -p $dbName -e 'SELECT fldid, fldfirstname, fldlastname FROM drivers'" -ForegroundColor White

Write-Host "`n=== Photo generation setup complete ===" -ForegroundColor Green
