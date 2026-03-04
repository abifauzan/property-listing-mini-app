package http

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
)

// contextKey is a custom type for context keys to avoid collisions.
type contextKey string

// RequestIDKey is the context key for request IDs.
const RequestIDKey contextKey = "requestID"

// ErrorCode represents standardized error codes.
type ErrorCode string

const (
	ErrCodeBadRequest     ErrorCode = "BAD_REQUEST"
	ErrCodeNotFound       ErrorCode = "NOT_FOUND"
	ErrCodeInternalServer ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrCodeInvalidInput   ErrorCode = "INVALID_INPUT"
)

// ErrorResponse represents a standardized error response.
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error information.
type ErrorDetail struct {
	Code      ErrorCode `json:"code"`
	Message   string    `json:"message"`
	RequestID string    `json:"requestId,omitempty"`
}

// writeJSON writes a JSON response with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("failed to write JSON response", "error", err)
	}
}

// writeError writes a standardized error response.
func writeError(w http.ResponseWriter, status int, code ErrorCode, message string, requestID string) {
	resp := ErrorResponse{
		Error: ErrorDetail{
			Code:      code,
			Message:   message,
			RequestID: requestID,
		},
	}
	writeJSON(w, status, resp)
}

// writeSuccess writes a successful property listing response.
func writeListingSuccess(w http.ResponseWriter, listings []domain.PropertyListing) {
	resp := domain.PropertyListingResponse{
		Data: domain.PropertyListingData{
			PropertyListings: listings,
		},
	}
	writeJSON(w, http.StatusOK, resp)
}

// writeDetailSuccess writes a successful property detail response.
func writeDetailSuccess(w http.ResponseWriter, property *domain.Property) {
	resp := domain.PropertyDetailResponse{
		Data: domain.PropertyDetailData{
			PropertyListings: []domain.Property{*property},
		},
	}
	writeJSON(w, http.StatusOK, resp)
}
