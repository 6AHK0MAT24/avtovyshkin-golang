#!/bin/bash
set -e

echo "Applying MySQL database migrations..."

echo "This script is for reference only."
echo "The project now uses MySQL on localhost:3306."

echo "To apply migrations manually:"
echo "1. Use MySQL command line or phpMyAdmin"
echo "2. Connect to localhost:3306"
echo "3. Database: u3424187_avtovyshkin"

echo "Example MySQL command:"
echo "mysql -h localhost -u u3424187_root_avtovyshkin -p u3424187_avtovyshkin < migrations-mysql/001_create_drivers_table.sql"

echo "Migration setup complete!"
