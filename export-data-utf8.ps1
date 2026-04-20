# Export MySQL data to CSV with UTF-8 encoding
Write-Host "=== Export MySQL Data (UTF-8) ===" -ForegroundColor Green

Write-Host "`nThis script is for reference only." -ForegroundColor Yellow
Write-Host "The project now uses MySQL on localhost:3306." -ForegroundColor Cyan

Write-Host "`nTo export data with UTF-8 encoding:" -ForegroundColor Yellow
Write-Host "1. Use MySQL command line or phpMyAdmin" -ForegroundColor White
Write-Host "2. Connect to localhost:3306" -ForegroundColor White
Write-Host "3. Database: u3424187_avtovyshkin" -ForegroundColor White

Write-Host "`nExample MySQL command:" -ForegroundColor Yellow
Write-Host "mysql -h localhost -u u3424187_root_avtovyshkin -p u3424187_avtovyshkin --default-character-set=utf8mb4 -e 'SELECT * FROM drivers'" -ForegroundColor White

Write-Host "`n=== Export Setup Complete ===" -ForegroundColor Green
