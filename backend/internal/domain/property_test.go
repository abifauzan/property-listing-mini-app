package domain

import "testing"

func TestPropertyIsValid(t *testing.T) {
	tests := []struct {
		name     string
		property Property
		want     bool
	}{
		{
			name: "valid property",
			property: Property{
				DocumentID: "id-1",
				Title:      "Sample Property",
				Price:      "Rp 1.000.000",
				Banner:     Banner{URL: "https://example.com/banner.jpg"},
			},
			want: true,
		},
		{
			name: "missing document ID",
			property: Property{
				DocumentID: "",
				Title:      "Sample Property",
				Price:      "Rp 1.000.000",
				Banner:     Banner{URL: "https://example.com/banner.jpg"},
			},
			want: false,
		},
		{
			name: "missing title",
			property: Property{
				DocumentID: "id-1",
				Title:      "",
				Price:      "Rp 1.000.000",
				Banner:     Banner{URL: "https://example.com/banner.jpg"},
			},
			want: false,
		},
		{
			name: "missing price",
			property: Property{
				DocumentID: "id-1",
				Title:      "Sample Property",
				Price:      "",
				Banner:     Banner{URL: "https://example.com/banner.jpg"},
			},
			want: false,
		},
		{
			name: "missing banner URL",
			property: Property{
				DocumentID: "id-1",
				Title:      "Sample Property",
				Price:      "Rp 1.000.000",
				Banner:     Banner{URL: ""},
			},
			want: false,
		},
		{
			name: "all fields empty",
			property: Property{
				DocumentID: "",
				Title:      "",
				Price:      "",
				Banner:     Banner{URL: ""},
			},
			want: false,
		},
		{
			name: "valid with optional fields",
			property: Property{
				DocumentID:  "id-1",
				Title:       "Sample Property",
				Price:       "Rp 1.000.000",
				Banner:      Banner{URL: "https://example.com/banner.jpg"},
				Description: "A nice property",
				Images: []Image{
					{URL: "https://example.com/img1.jpg"},
					{URL: "https://example.com/img2.jpg"},
				},
				Facilities: []string{"WiFi", "Parking"},
				Terms:      "No smoking",
				Conditions: "Good condition",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.property.IsValid()
			if got != tt.want {
				t.Errorf("Property.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
