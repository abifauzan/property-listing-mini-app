-- Create properties table
CREATE TABLE IF NOT EXISTS properties (
    id SERIAL PRIMARY KEY,
    document_id VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    price VARCHAR(50) NOT NULL,
    description TEXT,
    terms TEXT,
    conditions TEXT,
    banner_url VARCHAR(500),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_properties_document_id ON properties(document_id);
CREATE INDEX IF NOT EXISTS idx_properties_title ON properties(title);
CREATE INDEX IF NOT EXISTS idx_properties_created_at ON properties(created_at);
