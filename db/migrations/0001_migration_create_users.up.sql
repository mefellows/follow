CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    device_id UUID,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    
    -- Generic Model data
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
