package domain

import "context"

// Banner represents a property's banner image.
type Banner struct {
	URL string `json:"url"`
}

// Image represents a single image in the property gallery.
type Image struct {
	URL string `json:"url"`
}

// Property represents the full property entity used across layers.
type Property struct {
	DocumentID  string   `json:"documentId"`
	Banner      Banner   `json:"Banner"`
	Title       string   `json:"Title"`
	Price       string   `json:"Price"`
	CreatedAt   string   `json:"createdAt"`
	Description string   `json:"Description,omitempty"`
	Images      []Image  `json:"Images,omitempty"`
	Facilities  []string `json:"Facilities,omitempty"`
	Terms       string   `json:"Terms,omitempty"`
	Conditions  string   `json:"Conditions,omitempty"`
}

// PropertyListingResponse is the top-level JSON envelope for listing responses.
type PropertyListingResponse struct {
	Data PropertyListingData `json:"data"`
}

// PropertyListingData wraps the property listings array.
type PropertyListingData struct {
	PropertyListings []PropertyListing `json:"propertyListings"`
}

// PropertyListing is the subset of fields returned in the listing endpoint.
type PropertyListing struct {
	DocumentID string `json:"documentId"`
	Banner     Banner `json:"Banner"`
	Title      string `json:"Title"`
	Price      string `json:"Price"`
	CreatedAt  string `json:"createdAt"`
}

// PropertyDetailResponse is the top-level JSON envelope for detail responses.
type PropertyDetailResponse struct {
	Data PropertyDetailData `json:"data"`
}

// PropertyDetailData wraps a single property detail inside the listings array
// to stay consistent with the provided schema contract.
type PropertyDetailData struct {
	PropertyListings []Property `json:"propertyListings"`
}

// PropertyRepository defines the data-access contract.
type PropertyRepository interface {
	GetAll(ctx context.Context) ([]Property, error)
	GetByID(ctx context.Context, id string) (*Property, error)
	Search(ctx context.Context, query string) ([]Property, error)
}

// PropertyUsecase defines the business-logic contract.
type PropertyUsecase interface {
	GetListings(ctx context.Context, search string) ([]PropertyListing, error)
	GetDetail(ctx context.Context, id string) (*Property, error)
}

// IsValid checks if a property has all required fields.
func (p *Property) IsValid() bool {
	return p.DocumentID != "" &&
		p.Title != "" &&
		p.Price != "" &&
		p.Banner.URL != ""
}
