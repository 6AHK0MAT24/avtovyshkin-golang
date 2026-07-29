package utils

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

// Error implements the error interface
func (ve ValidationErrors) Error() string {
	var messages []string
	for _, err := range ve {
		messages = append(messages, fmt.Sprintf("%s: %s", err.Field, err.Message))
	}
	return strings.Join(messages, "; ")
}

// Validator provides validation utilities
type Validator struct {
	errors ValidationErrors
}

// NewValidator creates a new Validator
func NewValidator() *Validator {
	return &Validator{
		errors: make(ValidationErrors, 0),
	}
}

// AddError adds a validation error
func (v *Validator) AddError(field, message string) {
	v.errors = append(v.errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

// HasErrors returns true if there are validation errors
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// GetErrors returns all validation errors
func (v *Validator) GetErrors() ValidationErrors {
	return v.errors
}

// ToError converts validation errors to a single error
func (v *Validator) ToError() error {
	if !v.HasErrors() {
		return nil
	}
	return v.errors
}

// Required checks if a field is not empty
func (v *Validator) Required(field, value string) {
	if strings.TrimSpace(value) == "" {
		v.AddError(field, "is required")
	}
}

// MinLength checks if a field meets minimum length
func (v *Validator) MinLength(field, value string, min int) {
	if len(value) < min {
		v.AddError(field, fmt.Sprintf("must be at least %d characters", min))
	}
}

// MaxLength checks if a field doesn't exceed maximum length
func (v *Validator) MaxLength(field, value string, max int) {
	if len(value) > max {
		v.AddError(field, fmt.Sprintf("must not exceed %d characters", max))
	}
}

// Email checks if a field is a valid email
func (v *Validator) Email(field, value string) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(value) {
		v.AddError(field, "must be a valid email")
	}
}

// Phone checks if a field is a valid phone number
func (v *Validator) Phone(field, value string) {
	// Remove all non-digit characters
	phoneRegex := regexp.MustCompile(`^\+?[\d\s\-\(\)]+$`)
	if !phoneRegex.MatchString(value) {
		v.AddError(field, "must be a valid phone number")
	}
}

// URL checks if a field is a valid URL
func (v *Validator) URL(field, value string) {
	urlRegex := regexp.MustCompile(`^https?:\/\/.+$`)
	if !urlRegex.MatchString(value) {
		v.AddError(field, "must be a valid URL")
	}
}

// Match checks if a field matches a regular expression
func (v *Validator) Match(field, value, pattern string) {
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		v.AddError(field, "invalid pattern")
		return
	}
	if !matched {
		v.AddError(field, "format is invalid")
	}
}

// In checks if a field value is in the allowed values
func (v *Validator) In(field, value string, allowed []string) {
	for _, allowedValue := range allowed {
		if value == allowedValue {
			return
		}
	}
	v.AddError(field, fmt.Sprintf("must be one of: %s", strings.Join(allowed, ", ")))
}

// Min checks if a numeric field meets minimum value
func (v *Validator) Min(field string, value, min int) {
	if value < min {
		v.AddError(field, fmt.Sprintf("must be at least %d", min))
	}
}

// Max checks if a numeric field doesn't exceed maximum value
func (v *Validator) Max(field string, value, max int) {
	if value > max {
		v.AddError(field, fmt.Sprintf("must not exceed %d", max))
	}
}

// Range checks if a numeric field is within a range
func (v *Validator) Range(field string, value, min, max int) {
	if value < min || value > max {
		v.AddError(field, fmt.Sprintf("must be between %d and %d", min, max))
	}
}

// PasswordStrength checks if a password meets strength requirements
func (v *Validator) PasswordStrength(field, password string) {
	var (
		hasMinLen  = len(password) >= 8
		hasNumber  = false
		hasUpper   = false
		hasLower   = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		v.AddError(field, "must be at least 8 characters")
	}
	if !hasNumber {
		v.AddError(field, "must contain at least one number")
	}
	if !hasUpper {
		v.AddError(field, "must contain at least one uppercase letter")
	}
	if !hasLower {
		v.AddError(field, "must contain at least one lowercase letter")
	}
	if !hasSpecial {
		v.AddError(field, "must contain at least one special character")
	}
}

// ValidateStruct validates a struct based on struct tags
// This is a placeholder for more advanced struct validation
func ValidateStruct(s interface{}) error {
	// Implement struct validation based on tags
	// For now, return nil
	return nil
}

// IsEmpty checks if a value is empty
func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v) == ""
	case []interface{}:
		return len(v) == 0
	case map[string]interface{}:
		return len(v) == 0
	default:
		return false
	}
}

// SanitizeString removes potentially dangerous characters from a string
func SanitizeString(s string) string {
	// Remove HTML tags and other potentially dangerous content
	// This is a basic implementation - consider using a proper HTML sanitizer
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(s, "")
}
