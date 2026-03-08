package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/config"
	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: migrate <command> [args]")
		log.Fatal("Commands:")
		log.Fatal("  db-up - Run database migrations")
		log.Fatal("  db-down - Rollback database migrations")
		log.Fatal("  seed - Load JSON data into database")
	}

	cfg := config.Load()
	dsn := fmt.Sprintf("%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User, cfg.Database.Password,
		cfg.Database.Host, cfg.Database.Port,
		cfg.Database.DBName, cfg.Database.SSLMode)

	command := os.Args[1]

	switch command {
	case "db-up":
		if err := runMigrations(dsn); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Migrations completed successfully")

	case "db-down":
		if err := rollbackMigrations(dsn); err != nil {
			log.Fatalf("Failed to rollback migrations: %v", err)
		}
		log.Println("Rollback completed successfully")

	case "seed":
		jsonPath := cfg.DataPath
		if len(os.Args) > 2 {
			jsonPath = os.Args[2]
		}
		if err := seedDatabase(dsn, jsonPath); err != nil {
			log.Fatalf("Failed to seed database: %v", err)
		}
		log.Println("Database seeded successfully")

	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

func runMigrations(dsn string) error {
	db, err := sql.Open("postgres", "postgres://"+dsn)
	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}
	defer db.Close()

	migrationFiles := []string{
		"migrations/000001_create_properties_table.up.sql",
		"migrations/000002_create_images_table.up.sql",
		"migrations/000003_create_facilities_table.up.sql",
	}

	ctx := context.Background()
	for _, file := range migrationFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("reading migration file %s: %w", file, err)
		}

		if err := executeMigration(ctx, db, filepath.Base(file), string(content)); err != nil {
			return fmt.Errorf("executing migration %s: %w", file, err)
		}

		fmt.Printf("Applied migration: %s\n", file)
	}

	return nil
}

func rollbackMigrations(dsn string) error {
	db, err := sql.Open("postgres", "postgres://"+dsn)
	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}
	defer db.Close()

	migrationFiles := []string{
		"migrations/000003_create_facilities_table.down.sql",
		"migrations/000002_create_images_table.down.sql",
		"migrations/000001_create_properties_table.down.sql",
	}

	ctx := context.Background()
	for _, file := range migrationFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("reading migration file %s: %w", file, err)
		}

		if _, err := db.ExecContext(ctx, string(content)); err != nil {
			return fmt.Errorf("executing rollback %s: %w", file, err)
		}

		fmt.Printf("Rolled back migration: %s\n", file)
	}

	return nil
}

func executeMigration(ctx context.Context, db *sql.DB, name, content string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}
	defer tx.Rollback()

	// Create migrations table if it doesn't exist
	if err := createMigrationsTable(ctx, tx); err != nil {
		return fmt.Errorf("creating migrations table: %w", err)
	}

	// Check if migration already executed
	if executed, err := isMigrationExecuted(ctx, tx, name); err != nil {
		return fmt.Errorf("checking migration status: %w", err)
	} else if executed {
		return nil // Already executed
	}

	// Execute migration
	if _, err := tx.ExecContext(ctx, content); err != nil {
		return fmt.Errorf("executing migration: %w", err)
	}

	// Record migration as executed
	if err := recordMigration(ctx, tx, name); err != nil {
		return fmt.Errorf("recording migration: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing migration: %w", err)
	}

	return nil
}

func seedDatabase(dsn, jsonPath string) error {
	db, err := sql.Open("postgres", "postgres://"+dsn)
	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}
	defer db.Close()

	// Read JSON data
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("reading JSON file: %w", err)
	}

	var properties []domain.Property
	if err := json.Unmarshal(data, &properties); err != nil {
		return fmt.Errorf("parsing JSON: %w", err)
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}
	defer tx.Rollback()

	for _, prop := range properties {
		// Insert property
		var propertyID int
		err := tx.QueryRowContext(ctx, `
			INSERT INTO properties (document_id, title, price, description, terms, conditions, banner_url, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			ON CONFLICT (document_id) DO UPDATE SET
				title = EXCLUDED.title,
				price = EXCLUDED.price,
				description = EXCLUDED.description,
				terms = EXCLUDED.terms,
				conditions = EXCLUDED.conditions,
				banner_url = EXCLUDED.banner_url
			RETURNING id
		`, prop.DocumentID, prop.Title, prop.Price, prop.Description, prop.Terms, prop.Conditions, prop.Banner.URL, prop.CreatedAt).Scan(&propertyID)

		if err != nil {
			return fmt.Errorf("inserting property %s: %w", prop.DocumentID, err)
		}

		// Insert images
		for i, img := range prop.Images {
			_, err := tx.ExecContext(ctx, `
				INSERT INTO property_images (property_id, url, order_index)
				VALUES ($1, $2, $3)
				ON CONFLICT DO NOTHING
			`, propertyID, img.URL, i)
			if err != nil {
				return fmt.Errorf("inserting image for property %s: %w", prop.DocumentID, err)
			}
		}

		// Insert facilities
		for _, facility := range prop.Facilities {
			_, err := tx.ExecContext(ctx, `
				INSERT INTO property_facilities (property_id, facility_name)
				VALUES ($1, $2)
				ON CONFLICT DO NOTHING
			`, propertyID, facility)
			if err != nil {
				return fmt.Errorf("inserting facility for property %s: %w", prop.DocumentID, err)
			}
		}

		fmt.Printf("Seeded property: %s\n", prop.Title)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

func createMigrationsTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		)
	`)
	return err
}

func isMigrationExecuted(ctx context.Context, tx *sql.Tx, version string) (bool, error) {
	var count int
	err := tx.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM schema_migrations WHERE version = $1",
		version).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func recordMigration(ctx context.Context, tx *sql.Tx, version string) error {
	_, err := tx.ExecContext(ctx,
		"INSERT INTO schema_migrations (version) VALUES ($1)",
		version)
	return err
}
