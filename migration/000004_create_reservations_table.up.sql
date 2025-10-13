CREATE TABLE public.reservations (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    name VARCHAR(255) NOT NULL,
    reserver_phone_number VARCHAR(30),
    company_name VARCHAR(255),
    note TEXT,
    status_reservation VARCHAR(20) CHECK (status_reservation IN ('booked', 'paid', 'canceled')),
    sub_total_snack NUMERIC(12,2) DEFAULT 0,
    sub_total_rooms NUMERIC(12,2) DEFAULT 0,
    total NUMERIC(12,2) DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);

CREATE INDEX idx_reservations_user_id ON public.reservations USING btree (user_id);