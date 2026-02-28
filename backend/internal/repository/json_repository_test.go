package repository

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func setupTestFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	fp := filepath.Join(dir, "test.json")
	if err := os.WriteFile(fp, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}
	return fp
}

const testData = `[
  {
    "documentId": "id-1",
    "Banner": {"url": "https://example.com/img1.jpg"},
    "Title": "My Villa Sample",
    "Price": "750000",
    "createdAt": "2025-01-31T08:18:56.778Z",
    "Description": "A nice villa",
    "Images": [{"url": "https://example.com/img1.jpg"}],
    "Facilities": ["Kitchen"],
    "Terms": "Some terms",
    "Conditions": "Some conditions"
  },
  {
    "documentId": "id-2",
    "Banner": {"url": "https://example.com/img2.jpg"},
    "Title": "Beach House",
    "Price": "500000",
    "createdAt": "2024-12-12T12:30:50.940Z",
    "Description": "A beach house",
    "Images": [],
    "Facilities": [],
    "Terms": "",
    "Conditions": ""
  }
]`

func TestNewJSONPropertyRepository_InvalidPath(t *testing.T) {
	_, err := NewJSONPropertyRepository("/nonexistent/path.json")
	if err == nil {
		t.Fatal("expected error for invalid file path, got nil")
	}
}

func TestNewJSONPropertyRepository_InvalidJSON(t *testing.T) {
	fp := setupTestFile(t, `not valid json`)
	_, err := NewJSONPropertyRepository(fp)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestGetAll(t *testing.T) {
	fp := setupTestFile(t, testData)
	repo, err := NewJSONPropertyRepository(fp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	properties, err := repo.GetAll(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(properties) != 2 {
		t.Fatalf("expected 2 properties, got %d", len(properties))
	}
}

func TestGetByID(t *testing.T) {
	fp := setupTestFile(t, testData)
	repo, err := NewJSONPropertyRepository(fp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{name: "existing id", id: "id-1", wantErr: false},
		{name: "another existing id", id: "id-2", wantErr: false},
		{name: "non-existing id", id: "id-999", wantErr: true},
		{name: "empty id", id: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := repo.GetByID(context.Background(), tt.id)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if p.DocumentID != tt.id {
				t.Fatalf("expected documentId %q, got %q", tt.id, p.DocumentID)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	fp := setupTestFile(t, testData)
	repo, err := NewJSONPropertyRepository(fp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name     string
		query    string
		wantLen  int
	}{
		{name: "match one - villa sample", query: "villa sample", wantLen: 1},
		{name: "match two - case insensitive", query: "my", wantLen: 1},
		{name: "match none", query: "castle", wantLen: 0},
		{name: "partial match", query: "beach", wantLen: 1},
		{name: "empty query returns all", query: "", wantLen: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := repo.Search(context.Background(), tt.query)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(results) != tt.wantLen {
				t.Fatalf("expected %d results, got %d", tt.wantLen, len(results))
			}
		})
	}
}
