# Скрипт экспорта данных из MySQL
Write-Host "=== Экспорт данных из MySQL ===" -ForegroundColor Green

Write-Host "`nЭтот скрипт для справки." -ForegroundColor Yellow
Write-Host "Проект теперь использует MySQL на localhost:3306." -ForegroundColor Cyan

Write-Host "`nДля экспорта данных:" -ForegroundColor Yellow
Write-Host "1. Используйте MySQL командную строку или phpMyAdmin" -ForegroundColor White
Write-Host "2. Подключитесь к localhost:3306" -ForegroundColor White
Write-Host "3. База данных: u3424187_avtovyshkin" -ForegroundColor White

Write-Host "`nПример команды MySQL:" -ForegroundColor Yellow
Write-Host "mysql -h localhost -u u3424187_root_avtovyshkin -p u3424187_avtovyshkin -e 'SELECT * FROM drivers'" -ForegroundColor White

Write-Host "`n=== Настройка экспорта завершена ===" -ForegroundColor Green
