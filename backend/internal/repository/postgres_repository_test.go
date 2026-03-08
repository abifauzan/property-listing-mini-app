package repository

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	testDSN = "postgres://postgres:password@localhost:5432/property_listing_test?sslmode=disable"
)

type PostgresRepositoryTestSuite struct {
	suite.Suite
	repo *PostgresPropertyRepository
	db   *sql.DB
}

func (suite *PostgresRepositoryTestSuite) SetupSuite() {
	// Skip tests if PostgreSQL is not available
	if _, err := sql.Open("pgx", testDSN); err != nil {
		suite.T().Skip("PostgreSQL not available for testing")
	}

	// Create test repository
	repo, err := NewPostgresPropertyRepository(testDSN)
	require.NoError(suite.T(), err)
	suite.repo = repo

	// Open database for setup/cleanup
	db, err := sql.Open("pgx", testDSN)
	require.NoError(suite.T(), err)
	suite.db = db

	// Clean up any existing test data
	suite.cleanupTestData()
}

func (suite *PostgresRepositoryTestSuite) TearDownSuite() {
	suite.cleanupTestData()
	if suite.repo != nil {
		suite.repo.Close()
	}
	if suite.db != nil {
		suite.db.Close()
	}
}

func (suite *PostgresRepositoryTestSuite) SetupTest() {
	suite.cleanupTestData()
	suite.seedTestData()
}

func (suite *PostgresRepositoryTestSuite) cleanupTestData() {
	if suite.db == nil {
		return
	}

	tables := []string{"property_facilities", "property_images", "properties"}
	for _, table := range tables {
		_, _ = suite.db.Exec("DELETE FROM " + table)
	}
}

func (suite *PostgresRepositoryTestSuite) seedTestData() {
	if suite.db == nil {
		return
	}

	// Insert test property
	_, err := suite.db.Exec(`
		INSERT INTO properties (document_id, title, price, description, banner_url, created_at)
		VALUES ('test-prop-1', 'Test Villa', '750000', 'A beautiful test villa', 'https://example.com/banner.jpg', NOW())
	`)
	require.NoError(suite.T(), err)

	var propertyID int
	err = suite.db.QueryRow("SELECT id FROM properties WHERE document_id = 'test-prop-1'").Scan(&propertyID)
	require.NoError(suite.T(), err)

	// Insert test images
	_, err = suite.db.Exec(`
		INSERT INTO property_images (property_id, url, order_index)
		VALUES ($1, 'https://example.com/image1.jpg', 0),
		       ($1, 'https://example.com/image2.jpg', 1)
	`, propertyID)
	require.NoError(suite.T(), err)

	// Insert test facilities
	_, err = suite.db.Exec(`
		INSERT INTO property_facilities (property_id, facility_name)
		VALUES ($1, 'Kitchen'), ($1, 'Pool')
	`, propertyID)
	require.NoError(suite.T(), err)
}

func (suite *PostgresRepositoryTestSuite) TestGetAll() {
	ctx := context.Background()

	properties, err := suite.repo.GetAll(ctx)
	require.NoError(suite.T(), err)
	assert.Len(suite.T(), properties, 1)

	property := properties[0]
	assert.Equal(suite.T(), "test-prop-1", property.DocumentID)
	assert.Equal(suite.T(), "Test Villa", property.Title)
	assert.Equal(suite.T(), "750000", property.Price)
	assert.Equal(suite.T(), "A beautiful test villa", property.Description)
	assert.Equal(suite.T(), "https://example.com/banner.jpg", property.Banner.URL)
	assert.Len(suite.T(), property.Images, 2)
	assert.Len(suite.T(), property.Facilities, 2)
	assert.Contains(suite.T(), property.Facilities, "Kitchen")
	assert.Contains(suite.T(), property.Facilities, "Pool")
}

func (suite *PostgresRepositoryTestSuite) TestGetByID() {
	ctx := context.Background()

	property, err := suite.repo.GetByID(ctx, "test-prop-1")
	require.NoError(suite.T(), err)

	assert.Equal(suite.T(), "test-prop-1", property.DocumentID)
	assert.Equal(suite.T(), "Test Villa", property.Title)
	assert.Equal(suite.T(), "750000", property.Price)
	assert.Len(suite.T(), property.Images, 2)
	assert.Len(suite.T(), property.Facilities, 2)
}

func (suite *PostgresRepositoryTestSuite) TestGetByIDNotFound() {
	ctx := context.Background()

	property, err := suite.repo.GetByID(ctx, "non-existent")
	require.Error(suite.T(), err)
	assert.Nil(suite.T(), property)
	assert.Contains(suite.T(), err.Error(), "not found")
}

func (suite *PostgresRepositoryTestSuite) TestSearch() {
	ctx := context.Background()

	// Test exact match
	properties, err := suite.repo.Search(ctx, "Test Villa")
	require.NoError(suite.T(), err)
	assert.Len(suite.T(), properties, 1)

	// Test partial match
	properties, err = suite.repo.Search(ctx, "villa")
	require.NoError(suite.T(), err)
	assert.Len(suite.T(), properties, 1)

	// Test case insensitive
	properties, err = suite.repo.Search(ctx, "VILLA")
	require.NoError(suite.T(), err)
	assert.Len(suite.T(), properties, 1)

	// Test no match
	properties, err = suite.repo.Search(ctx, "non-existent")
	require.NoError(suite.T(), err)
	assert.Len(suite.T(), properties, 0)
}

func TestPostgresPropertyRepository_New(t *testing.T) {
	// Test invalid DSN
	_, err := NewPostgresPropertyRepository("invalid://connection")
	assert.Error(t, err)

	// Skip if PostgreSQL not available
	if _, err := sql.Open("pgx", testDSN); err != nil {
		t.Skip("PostgreSQL not available for testing")
	}

	// Test valid DSN
	repo, err := NewPostgresPropertyRepository(testDSN)
	require.NoError(t, err)
	assert.NotNil(t, repo)
	repo.Close()
}

func TestPostgresRepositoryTestSuite(t *testing.T) {
	// Check if we should run integration tests
	if os.Getenv("SKIP_DB_TESTS") != "" {
		t.Skip("Skipping database tests")
	}

	suite.Run(t, new(PostgresRepositoryTestSuite))
}
