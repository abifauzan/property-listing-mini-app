package http

import (
	"encoding/json"
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
	search := r.URL.Query().Get("search")

	listings, err := h.usecase.GetListings(r.Context(), search)
	if err != nil {
		slog.Error("failed to get listings", "error", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	resp := domain.PropertyListingResponse{
		Data: domain.PropertyListingData{
			PropertyListings: listings,
		},
	}
	writeJSON(w, http.StatusOK, resp)
}

// GetDetail handles GET /api/v1/properties/{id}
func (h *PropertyHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing property id"})
		return
	}

	property, err := h.usecase.GetDetail(r.Context(), id)
	if err != nil {
		slog.Warn("property not found", "id", id, "error", err)
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "property not found"})
		return
	}

	resp := domain.PropertyDetailResponse{
		Data: domain.PropertyDetailData{
			PropertyListings: []domain.Property{*property},
		},
	}
	writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("failed to write JSON response", "error", err)
	}
}
