# Export MySQL data to CSV
Write-Host "=== Export MySQL Data ===" -ForegroundColor Green

Write-Host "`nThis script is for reference only." -ForegroundColor Yellow
Write-Host "The project now uses MySQL on localhost:3306." -ForegroundColor Cyan

Write-Host "`nTo export data:" -ForegroundColor Yellow
Write-Host "1. Use MySQL command line or phpMyAdmin" -ForegroundColor White
Write-Host "2. Connect to localhost:3306" -ForegroundColor White
Write-Host "3. Database: u3424187_avtovyshkin" -ForegroundColor White

Write-Host "`nExample MySQL command:" -ForegroundColor Yellow
Write-Host "mysql -h localhost -u u3424187_root_avtovyshkin -p u3424187_avtovyshkin -e 'SELECT * FROM drivers'" -ForegroundColor White

Write-Host "`n=== Export Setup Complete ===" -ForegroundColor Green
