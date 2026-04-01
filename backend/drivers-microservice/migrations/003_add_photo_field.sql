-- Migration: 003_add_photo_field.sql
-- Description: Add photo field to drivers table

-- Add photo column to drivers table
ALTER TABLE drivers ADD COLUMN IF NOT EXISTS fldPhoto TEXT;

-- Add comment for the new column
COMMENT ON COLUMN drivers.fldPhoto IS 'Path to driver profile photo file (base64 encoded)';
