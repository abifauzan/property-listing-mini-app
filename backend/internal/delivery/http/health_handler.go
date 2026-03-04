package http

import (
	"net/http"
	"time"
)

var startTime = time.Now()

// HealthResponse represents the health check response.
type HealthResponse struct {
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
	Version string `json:"version,omitempty"`
}

// ReadinessResponse represents the readiness check response.
type ReadinessResponse struct {
	Status       string `json:"status"`
	DataLoaded   bool   `json:"dataLoaded"`
	PropertyCount int   `json:"propertyCount,omitempty"`
}

// HealthHandler holds health check handlers.
type HealthHandler struct {
	dataLoaded    bool
	propertyCount int
}

// NewHealthHandler creates a new health handler.
func NewHealthHandler(dataLoaded bool, propertyCount int) *HealthHandler {
	return &HealthHandler{
		dataLoaded:    dataLoaded,
		propertyCount: propertyCount,
	}
}

// RegisterRoutes registers health check routes.
func (h *HealthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Ready)
}

// Health returns basic health status.
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)
	resp := HealthResponse{
		Status:  "ok",
		Uptime:  uptime.String(),
		Version: "1.0.0",
	}
	writeJSON(w, http.StatusOK, resp)
}

// Ready returns readiness status including data load state.
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	resp := ReadinessResponse{
		Status:        "ready",
		DataLoaded:    h.dataLoaded,
		PropertyCount: h.propertyCount,
	}

	if !h.dataLoaded {
		resp.Status = "not_ready"
		writeJSON(w, http.StatusServiceUnavailable, resp)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
