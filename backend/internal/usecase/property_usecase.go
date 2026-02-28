package usecase

import (
	"context"
	"fmt"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
)

// propertyUsecase implements domain.PropertyUsecase.
type propertyUsecase struct {
	repo domain.PropertyRepository
}

// NewPropertyUsecase returns a new PropertyUsecase backed by the given repository.
func NewPropertyUsecase(repo domain.PropertyRepository) domain.PropertyUsecase {
	return &propertyUsecase{repo: repo}
}

// GetListings returns a slice of PropertyListing. If search is non-empty,
// results are filtered by title (case-insensitive).
func (u *propertyUsecase) GetListings(ctx context.Context, search string) ([]domain.PropertyListing, error) {
	var properties []domain.Property
	var err error

	if search != "" {
		properties, err = u.repo.Search(ctx, search)
	} else {
		properties, err = u.repo.GetAll(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("fetching listings: %w", err)
	}

	listings := make([]domain.PropertyListing, 0, len(properties))
	for _, p := range properties {
		listings = append(listings, domain.PropertyListing{
			DocumentID: p.DocumentID,
			Banner:     p.Banner,
			Title:      p.Title,
			Price:      p.Price,
			CreatedAt:  p.CreatedAt,
		})
	}
	return listings, nil
}

// GetDetail returns the full property detail for the given id.
func (u *propertyUsecase) GetDetail(ctx context.Context, id string) (*domain.Property, error) {
	property, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetching property detail: %w", err)
	}
	return property, nil
}
