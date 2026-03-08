-- Add a new test property to verify database connection
INSERT INTO properties (document_id, title, price, description, terms, conditions, banner_url, created_at)
VALUES (
    'test-property-2024',
    'Luxury Beachfront Villa',
    '1250000',
    'Stunning beachfront villa with panoramic ocean views. This modern architectural masterpiece features 5 bedrooms, 6 bathrooms, infinity pool, and direct beach access.',
    'Minimum 3 nights stay. No smoking indoors. Quiet hours after 10 PM.',
    'Valid credit card required for booking. Cancellation policy: 48 hours for full refund.',
    'https://images.unsplash.com/photo-1512917774080-9991f1c4c750?w=800',
    '2024-03-08T13:25:00.000Z'
) RETURNING id;

-- Add images for the new property
INSERT INTO property_images (property_id, url, order_index)
VALUES 
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'https://images.unsplash.com/photo-1512917774080-9991f1c4c750?w=800', 0),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'https://images.unsplash.com/photo-1448630360428-65456885c650?w=800', 1),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'https://images.unsplash.com/photo-1493801794477-d585b0c9bfd5?w=800', 2),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'https://images.unsplash.com/photo-1580587771525-78b9dba3b914?w=800', 3);

-- Add facilities for the new property
INSERT INTO property_facilities (property_id, facility_name)
VALUES 
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'Private Beach'),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'Infinity Pool'),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'Ocean View'),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'WiFi'),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'Air Conditioning'),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'Kitchen'),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'Parking'),
    ((SELECT id FROM properties WHERE document_id = 'test-property-2024'), 'Gym');
