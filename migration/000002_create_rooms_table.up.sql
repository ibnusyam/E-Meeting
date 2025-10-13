CREATE TABLE public.rooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    capacity INTEGER NOT NULL,
    price NUMERIC(12,2) NOT NULL,
    type VARCHAR(20) CHECK (type IN ('small', 'medium', 'large')),
    images_url VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);