package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
)

// mockUsecase implements domain.PropertyUsecase for testing.
type mockUsecase struct {
	listings []domain.PropertyListing
	detail   *domain.Property
	err      error
}

func (m *mockUsecase) GetListings(_ context.Context, _ string) ([]domain.PropertyListing, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.listings, nil
}

func (m *mockUsecase) GetDetail(_ context.Context, _ string) (*domain.Property, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.detail, nil
}

var testListings = []domain.PropertyListing{
	{
		DocumentID: "id-1",
		Banner:     domain.Banner{URL: "https://example.com/img1.jpg"},
		Title:      "My Villa Sample",
		Price:      "750000",
		CreatedAt:  "2025-01-31T08:18:56.778Z",
	},
}

var testDetail = &domain.Property{
	DocumentID:  "id-1",
	Banner:      domain.Banner{URL: "https://example.com/img1.jpg"},
	Title:       "My Villa Sample",
	Price:       "750000",
	CreatedAt:   "2025-01-31T08:18:56.778Z",
	Description: "A nice villa",
	Facilities:  []string{"Kitchen"},
}

func TestGetListings_Success(t *testing.T) {
	uc := &mockUsecase{listings: testListings}
	handler := NewPropertyHandler(uc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/properties", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp domain.PropertyListingResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(resp.Data.PropertyListings) != 1 {
		t.Fatalf("expected 1 listing, got %d", len(resp.Data.PropertyListings))
	}
	if resp.Data.PropertyListings[0].DocumentID != "id-1" {
		t.Fatalf("expected id 'id-1', got %q", resp.Data.PropertyListings[0].DocumentID)
	}
}

func TestGetListings_WithSearchParam(t *testing.T) {
	uc := &mockUsecase{listings: testListings}
	handler := NewPropertyHandler(uc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/properties?search=villa", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func TestGetListings_UsecaseError(t *testing.T) {
	uc := &mockUsecase{err: fmt.Errorf("internal error")}
	handler := NewPropertyHandler(uc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/properties", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rec.Code)
	}
}

func TestGetDetail_Success(t *testing.T) {
	uc := &mockUsecase{detail: testDetail}
	handler := NewPropertyHandler(uc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/properties/id-1", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp domain.PropertyDetailResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(resp.Data.PropertyListings) != 1 {
		t.Fatalf("expected 1 detail, got %d", len(resp.Data.PropertyListings))
	}
	if resp.Data.PropertyListings[0].DocumentID != "id-1" {
		t.Fatalf("expected id 'id-1', got %q", resp.Data.PropertyListings[0].DocumentID)
	}
}

func TestGetDetail_NotFound(t *testing.T) {
	uc := &mockUsecase{err: fmt.Errorf("not found")}
	handler := NewPropertyHandler(uc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/properties/id-999", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rec.Code)
	}
}

func TestGetListings_InvalidSearchQuery(t *testing.T) {
	uc := &mockUsecase{listings: testListings}
	handler := NewPropertyHandler(uc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	longQuery := string(make([]byte, 101))
	for i := range longQuery {
		longQuery = string(append([]byte(longQuery[:i]), 'a'))
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/properties?search="+longQuery, nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}

	var resp ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}
	if resp.Error.Code != ErrCodeInvalidInput {
		t.Fatalf("expected error code %s, got %s", ErrCodeInvalidInput, resp.Error.Code)
	}
}

func TestGetDetail_InvalidPropertyID(t *testing.T) {
	uc := &mockUsecase{detail: testDetail}
	handler := NewPropertyHandler(uc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/properties/invalid@id", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}

	var resp ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}
	if resp.Error.Code != ErrCodeInvalidInput {
		t.Fatalf("expected error code %s, got %s", ErrCodeInvalidInput, resp.Error.Code)
	}
}

func TestGetRequestID_WithContext(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	ctx := context.WithValue(req.Context(), RequestIDKey, "test-request-id")
	req = req.WithContext(ctx)

	id := getRequestID(req)
	if id != "test-request-id" {
		t.Fatalf("expected request ID 'test-request-id', got %q", id)
	}
}

func TestGetRequestID_NoContext(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	id := getRequestID(req)
	if id != "" {
		t.Fatalf("expected empty request ID, got %q", id)
	}
}
