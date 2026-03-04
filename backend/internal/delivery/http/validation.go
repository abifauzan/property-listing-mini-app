package http

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	maxSearchQueryLength = 100
	maxPropertyIDLength  = 50
	minPropertyIDLength  = 1
)

var (
	propertyIDPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
)

// ValidationError represents an input validation error.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateSearchQuery validates and sanitizes a search query string.
// Returns the sanitized query and an error if validation fails.
func ValidateSearchQuery(query string) (string, error) {
	if query == "" {
		return "", nil
	}

	trimmed := strings.TrimSpace(query)

	if utf8.RuneCountInString(trimmed) > maxSearchQueryLength {
		return "", &ValidationError{
			Field:   "search",
			Message: fmt.Sprintf("query exceeds maximum length of %d characters", maxSearchQueryLength),
		}
	}

	return trimmed, nil
}

// ValidatePropertyID validates a property ID.
// Returns an error if the ID is invalid.
func ValidatePropertyID(id string) error {
	if id == "" {
		return &ValidationError{
			Field:   "id",
			Message: "property id cannot be empty",
		}
	}

	length := len(id)
	if length < minPropertyIDLength || length > maxPropertyIDLength {
		return &ValidationError{
			Field:   "id",
			Message: fmt.Sprintf("property id must be between %d and %d characters", minPropertyIDLength, maxPropertyIDLength),
		}
	}

	if !propertyIDPattern.MatchString(id) {
		return &ValidationError{
			Field:   "id",
			Message: "property id can only contain alphanumeric characters, hyphens, and underscores",
		}
	}

	return nil
}
