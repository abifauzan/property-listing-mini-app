-- Create property_facilities table for facility tags
CREATE TABLE IF NOT EXISTS property_facilities (
    id SERIAL PRIMARY KEY,
    property_id INTEGER NOT NULL REFERENCES properties(id) ON DELETE CASCADE,
    facility_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(property_id, facility_name)
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_property_facilities_property_id ON property_facilities(property_id);
CREATE INDEX IF NOT EXISTS idx_property_facilities_name ON property_facilities(facility_name);
