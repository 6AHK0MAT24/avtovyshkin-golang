package utils

import (
	"regexp"
	"strings"
	"time"
)

// ValidatePhone validates Russian phone number format
func ValidatePhone(phone string) bool {
	// Remove all non-digit characters
	digits := regexp.MustCompile(`[^\d]`).ReplaceAllString(phone, "")
	
	// Check if it's a valid Russian phone number (11 digits starting with 7 or 8)
	if len(digits) != 11 {
		return false
	}
	
	return digits[0] == '7' || digits[0] == '8'
}

// FormatPhone formats phone number to standard format +7 (XXX) XXX-XX-XX
func FormatPhone(phone string) string {
	// Remove all non-digit characters
	digits := regexp.MustCompile(`[^\d]`).ReplaceAllString(phone, "")
	
	if len(digits) != 11 {
		return phone
	}
	
	// Convert to +7 format
	if digits[0] == '8' {
		digits = "7" + digits[1:]
	}
	
	return "+" + digits[0:1] + " (" + digits[1:4] + ") " + digits[4:7] + "-" + digits[7:9] + "-" + digits[9:11]
}

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	if email == "" {
		return true // Email is optional
	}
	
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidateDriverLicense validates Russian driver license number (10 digits)
func ValidateDriverLicense(license string) bool {
	// Remove all non-digit characters
	digits := regexp.MustCompile(`[^\d]`).ReplaceAllString(license, "")
	
	// Russian driver license is 10 digits
	return len(digits) == 10
}

// FormatDriverLicense formats driver license number to XX XX XXXXXX
func FormatDriverLicense(license string) string {
	// Remove all non-digit characters
	digits := regexp.MustCompile(`[^\d]`).ReplaceAllString(license, "")
	
	if len(digits) != 10 {
		return license
	}
	
	return digits[0:2] + " " + digits[2:4] + " " + digits[4:10]
}

// ValidatePassport validates Russian passport (series: 4 digits, number: 6 digits)
func ValidatePassport(series, number string) bool {
	if series == "" && number == "" {
		return true // Passport is optional
	}
	
	// Validate series (4 digits)
	if series != "" {
		seriesDigits := regexp.MustCompile(`[^\d]`).ReplaceAllString(series, "")
		if len(seriesDigits) != 4 {
			return false
		}
	}
	
	// Validate number (6 digits)
	if number != "" {
		numberDigits := regexp.MustCompile(`[^\d]`).ReplaceAllString(number, "")
		if len(numberDigits) != 6 {
			return false
		}
	}
	
	return true
}

// FormatPassport formats passport to XXXX XXXXXX
func FormatPassport(series, number string) string {
	seriesDigits := regexp.MustCompile(`[^\d]`).ReplaceAllString(series, "")
	numberDigits := regexp.MustCompile(`[^\d]`).ReplaceAllString(number, "")
	
	if len(seriesDigits) == 4 && len(numberDigits) == 6 {
		return seriesDigits + " " + numberDigits
	}
	
	return strings.TrimSpace(series + " " + number)
}

// ValidateAge validates that person is at least minAge years old
func ValidateAge(birthDate *time.Time, minAge int) bool {
	if birthDate == nil {
		return true // Birth date is optional
	}
	
	age := time.Now().Year() - birthDate.Year()
	
	// Adjust if birthday hasn't occurred yet this year
	now := time.Now()
	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}
	
	return age >= minAge
}

// ValidateFileType validates file type
func ValidateFileType(filename string, allowedTypes []string) bool {
	if len(allowedTypes) == 0 {
		return true
	}
	
	// Get file extension
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return false
	}
	
	ext := strings.ToLower(parts[len(parts)-1])
	
	for _, allowedType := range allowedTypes {
		if ext == strings.ToLower(allowedType) {
			return true
		}
	}
	
	return false
}

// SanitizeFilename sanitizes filename by removing potentially dangerous characters
func SanitizeFilename(filename string) string {
	// Remove path separators and other dangerous characters
	dangerousChars := []string{"/", "\\", "..", "<", ">", ":", "\"", "|", "?", "*"}
	
	sanitized := filename
	for _, char := range dangerousChars {
		sanitized = strings.ReplaceAll(sanitized, char, "")
	}
	
	return strings.TrimSpace(sanitized)
}

// TruncateString truncates string to max length with ellipsis
func TruncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength-3] + "..."
}
