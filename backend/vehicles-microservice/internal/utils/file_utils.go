package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)
// SaveVehicleImages saves multiple vehicle images to uploads/cars/{brand}_{height}/ directory
func SaveVehicleImages(files []*multipart.FileHeader, brand string, height float64, baseDir string) ([]string, error) {
	if len(files) == 0 {
		return []string{}, nil
	}

	// Create directory path: uploads/cars/{brand}_{height}/
	dirName := fmt.Sprintf("%s_%.0f", brand, height)
	uploadDir := filepath.Join(baseDir, "cars", dirName)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	var savedPaths []string

	for _, fileHeader := range files {
		// Open uploaded file
		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %w", fileHeader.Filename, err)
		}
		defer file.Close()

		// Generate unique filename
		uniqueFilename := GenerateUniqueFilename(fileHeader.Filename)
		filePath := filepath.Join(uploadDir, uniqueFilename)

		// Create destination file
		dst, err := os.Create(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
		defer dst.Close()

		// Copy file content
		if _, err := io.Copy(dst, file); err != nil {
			return nil, fmt.Errorf("failed to save file: %w", err)
		}

		// Convert to relative path with forward slashes
		relativePath := strings.ReplaceAll(filePath, "\\", "/")
		// Remove duplicate slashes
		relativePath = strings.ReplaceAll(relativePath, "//", "/")
		// Ensure path starts with /
		if !strings.HasPrefix(relativePath, "/") {
			relativePath = "/" + relativePath
		}
		savedPaths = append(savedPaths, relativePath)
	}

	return savedPaths, nil
}

// CreateVehicleDirectory creates directory for vehicle images
func CreateVehicleDirectory(baseDir, brand string, height float64) (string, error) {
	dirName := fmt.Sprintf("%s_%.0f", brand, height)
	uploadDir := filepath.Join(baseDir, "cars", dirName)

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create vehicle directory: %w", err)
	}

	return uploadDir, nil
}

// DeleteVehicleImages deletes all images for a specific vehicle
func DeleteVehicleImages(imagePaths []string) error {
	for _, imagePath := range imagePaths {
		if err := DeleteFile(imagePath); err != nil {
			// Log error but continue deleting other files
			fmt.Printf("Warning: failed to delete file %s: %v\n", imagePath, err)
		}
	}
	return nil
}

// DeleteVehicleDirectory deletes entire directory for a vehicle
func DeleteVehicleDirectory(baseDir, brand string, height float64) error {
	dirName := fmt.Sprintf("%s_%.0f", brand, height)
	vehicleDir := filepath.Join(baseDir, "cars", dirName)

	if err := os.RemoveAll(vehicleDir); err != nil {
		return fmt.Errorf("failed to delete vehicle directory: %w", err)
	}

	return nil
}

// GenerateUniqueFilename generates unique filename using UUID
// Format: {uuid}_{original_filename}
func GenerateUniqueFilename(originalName string) string {
	// Generate UUID
	uniqueID := uuid.New().String()

	// Sanitize original filename
	sanitizedName := SanitizeFilename(originalName)

	// Combine: uuid_sanitizedName
	if sanitizedName == "" {
		return uniqueID
	}

	return fmt.Sprintf("%s_%s", uniqueID, sanitizedName)
}
// SanitizeFilename removes special characters from filename
func SanitizeFilename(filename string) string {
	// Remove or replace special characters
	sanitized := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '-' || r == '_' {
			return r
		}
		return '_'
	}, filename)

	// Remove leading/trailing underscores and dots
	sanitized = strings.Trim(sanitized, "_.")
	return sanitized
}

// DeleteFile deletes file at specified path
func DeleteFile(filePath string) error {
	if filePath == "" {
		return nil
	}

	// Convert path from /uploads/... to ./uploads/... if needed
	cleanPath := filePath
	if strings.HasPrefix(cleanPath, "/") {
		cleanPath = "." + cleanPath
	}

	if err := os.Remove(cleanPath); err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist, that's okay
		}
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// ValidateImageFile validates that the uploaded file is an image
func ValidateImageFile(file multipart.File) bool {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return false
	}

	// Seek back to beginning
	file.Seek(0, 0)

	// Check if it's an image by content type
	contentType := http.DetectContentType(buffer)
	return strings.HasPrefix(contentType, "image/")
}

// GetFileExtension returns file extension
func GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
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
