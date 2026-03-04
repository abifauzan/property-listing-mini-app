package http

import (
	"strings"
	"testing"
)

func TestValidateSearchQuery(t *testing.T) {
	tests := []struct {
		name      string
		query     string
		wantQuery string
		wantErr   bool
	}{
		{
			name:      "empty query",
			query:     "",
			wantQuery: "",
			wantErr:   false,
		},
		{
			name:      "valid query",
			query:     "villa",
			wantQuery: "villa",
			wantErr:   false,
		},
		{
			name:      "query with whitespace",
			query:     "  beach house  ",
			wantQuery: "beach house",
			wantErr:   false,
		},
		{
			name:      "query at max length",
			query:     strings.Repeat("a", 100),
			wantQuery: strings.Repeat("a", 100),
			wantErr:   false,
		},
		{
			name:    "query exceeds max length",
			query:   strings.Repeat("a", 101),
			wantErr: true,
		},
		{
			name:      "query with special characters",
			query:     "villa-sample_123",
			wantQuery: "villa-sample_123",
			wantErr:   false,
		},
		{
			name:      "unicode query",
			query:     "别墅样本",
			wantQuery: "别墅样本",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateSearchQuery(tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSearchQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.wantQuery {
				t.Errorf("ValidateSearchQuery() = %q, want %q", got, tt.wantQuery)
			}
		})
	}
}

func TestValidatePropertyID(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "empty id",
			id:      "",
			wantErr: true,
		},
		{
			name:    "valid alphanumeric id",
			id:      "x0zzhpyrox8ixdy75e2zpw1m",
			wantErr: false,
		},
		{
			name:    "valid id with hyphens",
			id:      "property-123-abc",
			wantErr: false,
		},
		{
			name:    "valid id with underscores",
			id:      "property_123_abc",
			wantErr: false,
		},
		{
			name:    "valid id mixed",
			id:      "prop-123_abc",
			wantErr: false,
		},
		{
			name:    "id with spaces",
			id:      "prop 123",
			wantErr: true,
		},
		{
			name:    "id with special characters",
			id:      "prop@123",
			wantErr: true,
		},
		{
			name:    "id with slashes",
			id:      "prop/123",
			wantErr: true,
		},
		{
			name:    "id exceeds max length",
			id:      strings.Repeat("a", 51),
			wantErr: true,
		},
		{
			name:    "id at max length",
			id:      strings.Repeat("a", 50),
			wantErr: false,
		},
		{
			name:    "single character id",
			id:      "a",
			wantErr: false,
		},
		{
			name:    "numeric only id",
			id:      "12345",
			wantErr: false,
		},
		{
			name:    "id with dots",
			id:      "prop.123",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePropertyID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePropertyID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
