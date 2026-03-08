package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations runs all database migrations from the embedded filesystem.
func RunMigrations(dsn string) error {
	m, err := migrate.New(
		"file://migrations",
		"postgres://"+dsn,
	)
	if err != nil {
		return fmt.Errorf("creating migrate instance: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("running migrations: %w", err)
	}

	return nil
}

// RunMigrationsFromFilesystem runs migrations from a given filesystem.
func RunMigrationsFromFilesystem(dsn string, migrationFS fs.FS) error {
	// Create source driver from embedded filesystem
	d, err := iofs.New(migrationFS, "migrations")
	if err != nil {
		return fmt.Errorf("creating migration source: %w", err)
	}

	// Create database driver
	db, err := sql.Open("postgres", "postgres://"+dsn)
	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("creating database driver: %w", err)
	}

	// Create migrate instance
	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		return fmt.Errorf("creating migrate instance: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("running migrations: %w", err)
	}

	return nil
}

// GetMigrationFiles returns a sorted list of migration files from the filesystem.
func GetMigrationFiles(migrationFS fs.FS) ([]string, error) {
	var files []string

	err := fs.WalkDir(migrationFS, "migrations", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(path, ".up.sql") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("walking migration files: %w", err)
	}

	sort.Strings(files)
	return files, nil
}

// ExecuteMigrationFile executes a single migration file directly.
func ExecuteMigrationFile(ctx context.Context, db *sql.DB, filePath string, content string) error {
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
	migrationName := filepath.Base(filePath)
	if executed, err := isMigrationExecuted(ctx, tx, migrationName); err != nil {
		return fmt.Errorf("checking migration status: %w", err)
	} else if executed {
		return nil // Already executed
	}

	// Execute migration
	if _, err := tx.ExecContext(ctx, content); err != nil {
		return fmt.Errorf("executing migration %s: %w", filePath, err)
	}

	// Record migration as executed
	if err := recordMigration(ctx, tx, migrationName); err != nil {
		return fmt.Errorf("recording migration: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing migration: %w", err)
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

func isMigrationExecuted(ctx context.Context, tx *sql.Tx, migrationName string) (bool, error) {
	var count int
	err := tx.QueryRowContext(ctx, 
		"SELECT COUNT(*) FROM schema_migrations WHERE version = $1", 
		migrationName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func recordMigration(ctx context.Context, tx *sql.Tx, migrationName string) error {
	_, err := tx.ExecContext(ctx, 
		"INSERT INTO schema_migrations (version) VALUES ($1)", 
		migrationName)
	return err
}
