package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
)

// mockRepository implements domain.PropertyRepository for testing.
type mockRepository struct {
	properties []domain.Property
	err        error
}

func (m *mockRepository) GetAll(_ context.Context) ([]domain.Property, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.properties, nil
}

func (m *mockRepository) GetByID(_ context.Context, id string) (*domain.Property, error) {
	if m.err != nil {
		return nil, m.err
	}
	for i := range m.properties {
		if m.properties[i].DocumentID == id {
			return &m.properties[i], nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (m *mockRepository) Search(_ context.Context, _ string) ([]domain.Property, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.properties, nil
}

var testProperties = []domain.Property{
	{
		DocumentID: "id-1",
		Banner:     domain.Banner{URL: "https://example.com/img1.jpg"},
		Title:      "My Villa Sample",
		Price:      "750000",
		CreatedAt:  "2025-01-31T08:18:56.778Z",
	},
	{
		DocumentID: "id-2",
		Banner:     domain.Banner{URL: "https://example.com/img2.jpg"},
		Title:      "Beach House",
		Price:      "500000",
		CreatedAt:  "2024-12-12T12:30:50.940Z",
	},
}

func TestGetListings_NoSearch(t *testing.T) {
	repo := &mockRepository{properties: testProperties}
	uc := NewPropertyUsecase(repo)

	listings, err := uc.GetListings(context.Background(), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(listings) != 2 {
		t.Fatalf("expected 2 listings, got %d", len(listings))
	}
	if listings[0].DocumentID != "id-1" {
		t.Fatalf("expected first listing id 'id-1', got %q", listings[0].DocumentID)
	}
}

func TestGetListings_WithSearch(t *testing.T) {
	repo := &mockRepository{properties: testProperties[:1]}
	uc := NewPropertyUsecase(repo)

	listings, err := uc.GetListings(context.Background(), "villa")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(listings) != 1 {
		t.Fatalf("expected 1 listing, got %d", len(listings))
	}
}

func TestGetListings_RepoError(t *testing.T) {
	repo := &mockRepository{err: fmt.Errorf("db failure")}
	uc := NewPropertyUsecase(repo)

	_, err := uc.GetListings(context.Background(), "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetDetail_Found(t *testing.T) {
	repo := &mockRepository{properties: testProperties}
	uc := NewPropertyUsecase(repo)

	p, err := uc.GetDetail(context.Background(), "id-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.DocumentID != "id-1" {
		t.Fatalf("expected id 'id-1', got %q", p.DocumentID)
	}
}

func TestGetDetail_NotFound(t *testing.T) {
	repo := &mockRepository{properties: testProperties}
	uc := NewPropertyUsecase(repo)

	_, err := uc.GetDetail(context.Background(), "id-999")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetDetail_RepoError(t *testing.T) {
	repo := &mockRepository{err: fmt.Errorf("db failure")}
	uc := NewPropertyUsecase(repo)

	_, err := uc.GetDetail(context.Background(), "id-1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
