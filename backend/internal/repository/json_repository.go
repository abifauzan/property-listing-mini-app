package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
)

// JSONPropertyRepository implements domain.PropertyRepository using an in-memory
// store loaded from a JSON file. It is safe for concurrent reads.
type JSONPropertyRepository struct {
	mu         sync.RWMutex
	properties []domain.Property
}

// NewJSONPropertyRepository creates a new repository and loads data from the
// given JSON file path. Returns an error if the file cannot be read or parsed.
func NewJSONPropertyRepository(filePath string) (*JSONPropertyRepository, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("reading properties file: %w", err)
	}

	var properties []domain.Property
	if err := json.Unmarshal(data, &properties); err != nil {
		return nil, fmt.Errorf("parsing properties JSON: %w", err)
	}

	return &JSONPropertyRepository{
		properties: properties,
	}, nil
}

// GetAll returns all properties.
func (r *JSONPropertyRepository) GetAll(_ context.Context) ([]domain.Property, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.Property, len(r.properties))
	copy(result, r.properties)
	return result, nil
}

// GetByID returns a single property matching the given documentId.
func (r *JSONPropertyRepository) GetByID(_ context.Context, id string) (*domain.Property, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i := range r.properties {
		if r.properties[i].DocumentID == id {
			p := r.properties[i]
			return &p, nil
		}
	}
	return nil, fmt.Errorf("property with id %q not found", id)
}

// Search returns properties whose Title contains the query (case-insensitive).
func (r *JSONPropertyRepository) Search(_ context.Context, query string) ([]domain.Property, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	lowerQuery := strings.ToLower(query)
	var result []domain.Property
	for _, p := range r.properties {
		if strings.Contains(strings.ToLower(p.Title), lowerQuery) {
			result = append(result, p)
		}
	}
	return result, nil
}
