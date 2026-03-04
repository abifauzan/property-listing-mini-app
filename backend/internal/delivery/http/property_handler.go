package http

import (
	"log/slog"
	"net/http"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
)

// PropertyHandler holds HTTP handlers for property endpoints.
type PropertyHandler struct {
	usecase domain.PropertyUsecase
}

// NewPropertyHandler returns a new PropertyHandler.
func NewPropertyHandler(uc domain.PropertyUsecase) *PropertyHandler {
	return &PropertyHandler{usecase: uc}
}

// RegisterRoutes registers the property routes on the given mux.
func (h *PropertyHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/properties", h.GetListings)
	mux.HandleFunc("GET /api/v1/properties/{id}", h.GetDetail)
}

// GetListings handles GET /api/v1/properties?search={title}
func (h *PropertyHandler) GetListings(w http.ResponseWriter, r *http.Request) {
	requestID := getRequestID(r)
	searchQuery := r.URL.Query().Get("search")

	sanitizedQuery, err := ValidateSearchQuery(searchQuery)
	if err != nil {
		slog.Warn("invalid search query", "error", err, "requestId", requestID)
		writeError(w, http.StatusBadRequest, ErrCodeInvalidInput, err.Error(), requestID)
		return
	}

	listings, err := h.usecase.GetListings(r.Context(), sanitizedQuery)
	if err != nil {
		slog.Error("failed to get listings", "error", err, "requestId", requestID)
		writeError(w, http.StatusInternalServerError, ErrCodeInternalServer, "failed to retrieve property listings", requestID)
		return
	}

	writeListingSuccess(w, listings)
}

// GetDetail handles GET /api/v1/properties/{id}
func (h *PropertyHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	requestID := getRequestID(r)
	id := r.PathValue("id")

	if err := ValidatePropertyID(id); err != nil {
		slog.Warn("invalid property id", "id", id, "error", err, "requestId", requestID)
		writeError(w, http.StatusBadRequest, ErrCodeInvalidInput, err.Error(), requestID)
		return
	}

	property, err := h.usecase.GetDetail(r.Context(), id)
	if err != nil {
		slog.Warn("property not found", "id", id, "error", err, "requestId", requestID)
		writeError(w, http.StatusNotFound, ErrCodeNotFound, "property not found", requestID)
		return
	}

	writeDetailSuccess(w, property)
}

// getRequestID extracts the request ID from the request context.
func getRequestID(r *http.Request) string {
	if id := r.Context().Value(RequestIDKey); id != nil {
		if idStr, ok := id.(string); ok {
			return idStr
		}
	}
	return ""
}
