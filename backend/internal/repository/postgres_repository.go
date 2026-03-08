package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresPropertyRepository implements domain.PropertyRepository using PostgreSQL.
type PostgresPropertyRepository struct {
	pool *pgxpool.Pool
}

// NewPostgresPropertyRepository creates a new PostgreSQL repository.
func NewPostgresPropertyRepository(dsn string) (*PostgresPropertyRepository, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return &PostgresPropertyRepository{
		pool: pool,
	}, nil
}

// Close closes the database connection pool.
func (r *PostgresPropertyRepository) Close() {
	r.pool.Close()
}

// GetAll returns all properties with their associated images and facilities.
func (r *PostgresPropertyRepository) GetAll(ctx context.Context) ([]domain.Property, error) {
	query := `
		SELECT 
			p.id, p.document_id, p.title, p.price, p.description, p.terms, p.conditions,
			p.banner_url, p.created_at,
			COALESCE(
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'url', pi.url,
						'order_index', pi.order_index
					) ORDER BY pi.order_index
				) FILTER (WHERE pi.url IS NOT NULL), 
				'[]'
			) as images,
			COALESCE(
				ARRAY_AGG(DISTINCT pf.facility_name) FILTER (WHERE pf.facility_name IS NOT NULL), 
				ARRAY[]::VARCHAR[]
			) as facilities
		FROM properties p
		LEFT JOIN property_images pi ON p.id = pi.property_id
		LEFT JOIN property_facilities pf ON p.id = pf.property_id
		GROUP BY p.id, p.document_id, p.title, p.price, p.description, p.terms, p.conditions, p.banner_url, p.created_at
		ORDER BY p.created_at DESC
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("querying properties: %w", err)
	}
	defer rows.Close()

	var properties []domain.Property
	for rows.Next() {
		var prop domain.Property
		var id int
		var bannerURL sql.NullString
		var description, terms, conditions sql.NullString
		var createdAt time.Time
		var imagesJSON []byte
		var facilities []string

		if err := rows.Scan(
			&id, &prop.DocumentID, &prop.Title, &prop.Price,
			&description, &terms, &conditions,
			&bannerURL, &createdAt,
			&imagesJSON, &facilities,
		); err != nil {
			return nil, fmt.Errorf("scanning property row: %w", err)
		}

		// Handle nullable fields
		if bannerURL.Valid {
			prop.Banner = domain.Banner{URL: bannerURL.String}
		}
		if description.Valid {
			prop.Description = description.String
		}
		if terms.Valid {
			prop.Terms = terms.String
		}
		if conditions.Valid {
			prop.Conditions = conditions.String
		}

		prop.CreatedAt = createdAt.Format(time.RFC3339)
		prop.Facilities = facilities

		// Parse images JSON
		if len(imagesJSON) > 0 && string(imagesJSON) != "null" {
			var images []struct {
				URL        string `json:"url"`
				OrderIndex int    `json:"order_index"`
			}
			if err := json.Unmarshal(imagesJSON, &images); err == nil {
				prop.Images = make([]domain.Image, len(images))
				for i, img := range images {
					prop.Images[i] = domain.Image{URL: img.URL}
				}
			}
		}

		properties = append(properties, prop)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating property rows: %w", err)
	}

	return properties, nil
}

// GetByID returns a single property matching the given documentId.
func (r *PostgresPropertyRepository) GetByID(ctx context.Context, id string) (*domain.Property, error) {
	query := `
		SELECT 
			p.id, p.document_id, p.title, p.price, p.description, p.terms, p.conditions,
			p.banner_url, p.created_at,
			COALESCE(
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'url', pi.url,
						'order_index', pi.order_index
					) ORDER BY pi.order_index
				) FILTER (WHERE pi.url IS NOT NULL), 
				'[]'
			) as images,
			COALESCE(
				ARRAY_AGG(DISTINCT pf.facility_name) FILTER (WHERE pf.facility_name IS NOT NULL), 
				ARRAY[]::VARCHAR[]
			) as facilities
		FROM properties p
		LEFT JOIN property_images pi ON p.id = pi.property_id
		LEFT JOIN property_facilities pf ON p.id = pf.property_id
		WHERE p.document_id = $1
		GROUP BY p.id, p.document_id, p.title, p.price, p.description, p.terms, p.conditions, p.banner_url, p.created_at
	`

	var prop domain.Property
	var dbID int
	var bannerURL sql.NullString
	var description, terms, conditions sql.NullString
	var createdAt time.Time
	var imagesJSON []byte
	var facilities []string

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&dbID, &prop.DocumentID, &prop.Title, &prop.Price,
		&description, &terms, &conditions,
		&bannerURL, &createdAt,
		&imagesJSON, &facilities,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("property with id %q not found", id)
		}
		return nil, fmt.Errorf("querying property by id: %w", err)
	}

	// Handle nullable fields
	if bannerURL.Valid {
		prop.Banner = domain.Banner{URL: bannerURL.String}
	}
	if description.Valid {
		prop.Description = description.String
	}
	if terms.Valid {
		prop.Terms = terms.String
	}
	if conditions.Valid {
		prop.Conditions = conditions.String
	}

	prop.CreatedAt = createdAt.Format(time.RFC3339)
	prop.Facilities = facilities

	// Parse images JSON
	if len(imagesJSON) > 0 && string(imagesJSON) != "null" {
		var images []struct {
			URL        string `json:"url"`
			OrderIndex int    `json:"order_index"`
		}
		if err := json.Unmarshal(imagesJSON, &images); err == nil {
			prop.Images = make([]domain.Image, len(images))
			for i, img := range images {
				prop.Images[i] = domain.Image{URL: img.URL}
			}
		}
	}

	return &prop, nil
}

// Search returns properties whose Title contains the query (case-insensitive).
func (r *PostgresPropertyRepository) Search(ctx context.Context, query string) ([]domain.Property, error) {
	sqlQuery := `
		SELECT 
			p.id, p.document_id, p.title, p.price, p.description, p.terms, p.conditions,
			p.banner_url, p.created_at,
			COALESCE(
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'url', pi.url,
						'order_index', pi.order_index
					) ORDER BY pi.order_index
				) FILTER (WHERE pi.url IS NOT NULL), 
				'[]'
			) as images,
			COALESCE(
				ARRAY_AGG(DISTINCT pf.facility_name) FILTER (WHERE pf.facility_name IS NOT NULL), 
				ARRAY[]::VARCHAR[]
			) as facilities
		FROM properties p
		LEFT JOIN property_images pi ON p.id = pi.property_id
		LEFT JOIN property_facilities pf ON p.id = pf.property_id
		WHERE LOWER(p.title) LIKE LOWER($1)
		GROUP BY p.id, p.document_id, p.title, p.price, p.description, p.terms, p.conditions, p.banner_url, p.created_at
		ORDER BY p.created_at DESC
	`

	searchPattern := "%" + strings.ToLower(query) + "%"

	rows, err := r.pool.Query(ctx, sqlQuery, searchPattern)
	if err != nil {
		return nil, fmt.Errorf("querying properties: %w", err)
	}
	defer rows.Close()

	var properties []domain.Property
	for rows.Next() {
		var prop domain.Property
		var id int
		var bannerURL sql.NullString
		var description, terms, conditions sql.NullString
		var createdAt time.Time
		var imagesJSON []byte
		var facilities []string

		if err := rows.Scan(
			&id, &prop.DocumentID, &prop.Title, &prop.Price,
			&description, &terms, &conditions,
			&bannerURL, &createdAt,
			&imagesJSON, &facilities,
		); err != nil {
			return nil, fmt.Errorf("scanning property row: %w", err)
		}

		// Handle nullable fields
		if bannerURL.Valid {
			prop.Banner = domain.Banner{URL: bannerURL.String}
		}
		if description.Valid {
			prop.Description = description.String
		}
		if terms.Valid {
			prop.Terms = terms.String
		}
		if conditions.Valid {
			prop.Conditions = conditions.String
		}

		prop.CreatedAt = createdAt.Format(time.RFC3339)
		prop.Facilities = facilities

		// Parse images JSON
		if len(imagesJSON) > 0 && string(imagesJSON) != "null" {
			var images []struct {
				URL        string `json:"url"`
				OrderIndex int    `json:"order_index"`
			}
			if err := json.Unmarshal(imagesJSON, &images); err == nil {
				prop.Images = make([]domain.Image, len(images))
				for i, img := range images {
					prop.Images[i] = domain.Image{URL: img.URL}
				}
			}
		}

		properties = append(properties, prop)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating property rows: %w", err)
	}

	return properties, nil
}
