#!/bin/bash
set -e

echo "Applying database migrations..."

# Apply migrations in order
for migration in /docker-entrypoint-initdb.d/*.sql; do
    if [ -f "$migration" ]; then
        echo "Applying migration: $migration"
        psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f "$migration"
    fi
done

echo "All migrations applied successfully!"
