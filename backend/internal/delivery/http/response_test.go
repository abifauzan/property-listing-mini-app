package http

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
)

func TestWriteJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	data := map[string]string{"message": "test"}

	writeJSON(rec, 200, data)

	if rec.Code != 200 {
		t.Errorf("status = %d, want 200", rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
	}

	var result map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result["message"] != "test" {
		t.Errorf("message = %q, want %q", result["message"], "test")
	}
}

func TestWriteError(t *testing.T) {
	rec := httptest.NewRecorder()

	writeError(rec, 400, ErrCodeBadRequest, "invalid input", "req-123")

	if rec.Code != 400 {
		t.Errorf("status = %d, want 400", rec.Code)
	}

	var resp ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Error.Code != ErrCodeBadRequest {
		t.Errorf("error code = %s, want %s", resp.Error.Code, ErrCodeBadRequest)
	}

	if resp.Error.Message != "invalid input" {
		t.Errorf("error message = %q, want %q", resp.Error.Message, "invalid input")
	}

	if resp.Error.RequestID != "req-123" {
		t.Errorf("request ID = %q, want %q", resp.Error.RequestID, "req-123")
	}
}

func TestWriteListingSuccess(t *testing.T) {
	rec := httptest.NewRecorder()

	listings := []domain.PropertyListing{
		{
			DocumentID: "id-1",
			Title:      "Test Property",
			Price:      "1000000",
			Banner:     domain.Banner{URL: "https://example.com/img.jpg"},
		},
	}

	writeListingSuccess(rec, listings)

	if rec.Code != 200 {
		t.Errorf("status = %d, want 200", rec.Code)
	}

	var resp domain.PropertyListingResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(resp.Data.PropertyListings) != 1 {
		t.Errorf("listings count = %d, want 1", len(resp.Data.PropertyListings))
	}

	if resp.Data.PropertyListings[0].DocumentID != "id-1" {
		t.Errorf("document ID = %q, want %q", resp.Data.PropertyListings[0].DocumentID, "id-1")
	}
}

func TestWriteDetailSuccess(t *testing.T) {
	rec := httptest.NewRecorder()

	property := &domain.Property{
		DocumentID:  "id-1",
		Title:       "Test Property",
		Price:       "1000000",
		Banner:      domain.Banner{URL: "https://example.com/img.jpg"},
		Description: "A nice property",
	}

	writeDetailSuccess(rec, property)

	if rec.Code != 200 {
		t.Errorf("status = %d, want 200", rec.Code)
	}

	var resp domain.PropertyDetailResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(resp.Data.PropertyListings) != 1 {
		t.Errorf("property count = %d, want 1", len(resp.Data.PropertyListings))
	}

	if resp.Data.PropertyListings[0].DocumentID != "id-1" {
		t.Errorf("document ID = %q, want %q", resp.Data.PropertyListings[0].DocumentID, "id-1")
	}

	if resp.Data.PropertyListings[0].Description != "A nice property" {
		t.Errorf("description = %q, want %q", resp.Data.PropertyListings[0].Description, "A nice property")
	}
}
