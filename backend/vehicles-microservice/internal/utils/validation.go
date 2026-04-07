package utils

import (
	"regexp"
	"strings"
)

// ValidateVIN validates VIN code (17 alphanumeric characters without I, O, Q)
func ValidateVIN(vin string) bool {
	// Remove spaces and convert to uppercase
	vin = strings.ToUpper(strings.ReplaceAll(vin, " ", ""))

	// Check length
	if len(vin) != 17 {
		return false
	}

	// Check if all characters are alphanumeric (excluding I, O, Q)
	vinRegex := regexp.MustCompile(`^[A-HJ-NPR-Z0-9]{17}$`)
	return vinRegex.MatchString(vin)
}

// ValidateGarageNumber validates garage number format
func ValidateGarageNumber(garageNumber string) bool {
	// Garage number should not be empty
	if strings.TrimSpace(garageNumber) == "" {
		return false
	}

	// Allow alphanumeric characters, spaces, and common separators
	garageRegex := regexp.MustCompile(`^[A-Za-zА-Яа-я0-9\s\-/]+$`)
	return garageRegex.MatchString(garageNumber)
}

// ValidateFileType validates file type for images
func ValidateFileType(filename string, allowedTypes []string) bool {
	// Get file extension
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return false
	}

	ext := strings.ToLower(parts[len(parts)-1])

	// Check if extension is in allowed types
	for _, allowedType := range allowedTypes {
		if ext == strings.ToLower(allowedType) {
			return true
		}
	}

	return false
}
