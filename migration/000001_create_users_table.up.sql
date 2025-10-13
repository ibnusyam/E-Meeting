CREATE TABLE public.users (
    id SERIAL PRIMARY KEY,
    role VARCHAR(20) CHECK (role IN ('admin', 'customer')),
    username VARCHAR(255) NOT NULL,
    email VARCHAR(30) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(30),
    language VARCHAR(20) CHECK (language IN ('Indonesia', 'English')),
    status_user VARCHAR(20) CHECK (status_user IN ('active', 'inactive')),
    profile_picture VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON public.users USING btree (email);