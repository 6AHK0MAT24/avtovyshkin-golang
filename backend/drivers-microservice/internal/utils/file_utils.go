package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

// SaveFile saves uploaded file to specified directory
func SaveFile(file multipart.File, filename, uploadDir string) (string, error) {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	uniqueFilename := GenerateUniqueFilename(filename)
	filePath := filepath.Join(uploadDir, uniqueFilename)

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Return relative path
	return strings.ReplaceAll(filePath, "\\", "/"), nil
}

// SaveFileFromBytes saves file from byte slice
func SaveFileFromBytes(data []byte, filename, uploadDir string) (string, error) {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	uniqueFilename := GenerateUniqueFilename(filename)
	filePath := filepath.Join(uploadDir, uniqueFilename)

	// Write file
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	// Return relative path
	return strings.ReplaceAll(filePath, "\\", "/"), nil
}

// DeleteFile deletes file at specified path
func DeleteFile(filePath string) error {
	if filePath == "" {
		return nil
	}

	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist, that's okay
		}
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// GenerateUniqueFilename generates unique filename using UUID
func GenerateUniqueFilename(originalName string) string {
	ext := filepath.Ext(originalName)
	nameWithoutExt := strings.TrimSuffix(originalName, ext)
	
	// Sanitize filename
	sanitizedName := SanitizeFilename(nameWithoutExt)
	
	// Generate UUID
	uniqueID := uuid.New().String()
	
	// Combine: sanitizedName_uuid.ext
	if sanitizedName == "" {
		return uniqueID + ext
	}
	
	return fmt.Sprintf("%s_%s%s", sanitizedName, uniqueID, ext)
}

// GetFileExtension returns file extension
func GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}

// GetFileSize returns file size in bytes
func GetFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

// FileExists checks if file exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// CreateDriverDirectories creates all necessary directories for driver files
func CreateDriverDirectories(baseDir, driverID string) (photoDir, licenseDir, passportDir string, err error) {
	photoDir = filepath.Join(baseDir, "drivers", driverID, "photo")
	licenseDir = filepath.Join(baseDir, "drivers", driverID, "license")
	passportDir = filepath.Join(baseDir, "drivers", driverID, "passport")

	dirs := []string{photoDir, licenseDir, passportDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", "", "", fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return photoDir, licenseDir, passportDir, nil
}

// DeleteDriverFiles deletes all files for a specific driver
func DeleteDriverFiles(baseDir, driverID string) error {
	driverDir := filepath.Join(baseDir, "drivers", driverID)
	
	if err := os.RemoveAll(driverDir); err != nil {
		return fmt.Errorf("failed to delete driver files: %w", err)
	}
	
	return nil
}

// GetRelativePath returns relative path from base directory
func GetRelativePath(basePath, fullPath string) string {
	relPath, err := filepath.Rel(basePath, fullPath)
	if err != nil {
		return fullPath
	}
	return strings.ReplaceAll(relPath, "\\", "/")
}

// ValidateFileSize validates file size against max size
func ValidateFileSize(file multipart.File, maxSize int64) bool {
	// Seek to end to get size
	size, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return false
	}
	// Seek back to beginning
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return false
	}

	return size <= maxSize
}

// GetMimeType returns MIME type of file
func GetMimeType(file multipart.File) string {	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return ""
	}
	
	file.Seek(0, 0)
	
	// Detect MIME type
	mimeType := http.DetectContentType(buffer)
	return mimeType
}

// GenerateBackupFilename generates backup filename with timestamp
func GenerateBackupFilename(originalPath string) string {
	ext := filepath.Ext(originalPath)
	nameWithoutExt := strings.TrimSuffix(originalPath, ext)
	timestamp := time.Now().Format("20060102_150405")
	
	return fmt.Sprintf("%s_backup_%s%s", nameWithoutExt, timestamp, ext)
}
