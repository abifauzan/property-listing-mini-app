package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_Health(t *testing.T) {
	handler := NewHealthHandler(true, 3)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp HealthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status 'ok', got %q", resp.Status)
	}

	if resp.Uptime == "" {
		t.Error("expected uptime to be set")
	}

	if resp.Version != "1.0.0" {
		t.Errorf("expected version '1.0.0', got %q", resp.Version)
	}
}

func TestHealthHandler_Ready_DataLoaded(t *testing.T) {
	handler := NewHealthHandler(true, 5)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp ReadinessResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ready" {
		t.Errorf("expected status 'ready', got %q", resp.Status)
	}

	if !resp.DataLoaded {
		t.Error("expected dataLoaded to be true")
	}

	if resp.PropertyCount != 5 {
		t.Errorf("expected propertyCount 5, got %d", resp.PropertyCount)
	}
}

func TestHealthHandler_Ready_DataNotLoaded(t *testing.T) {
	handler := NewHealthHandler(false, 0)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", rec.Code)
	}

	var resp ReadinessResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "not_ready" {
		t.Errorf("expected status 'not_ready', got %q", resp.Status)
	}

	if resp.DataLoaded {
		t.Error("expected dataLoaded to be false")
	}
}
