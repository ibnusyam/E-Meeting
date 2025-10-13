CREATE TABLE public.reservation_details (
    id SERIAL PRIMARY KEY,
    reservation_id INTEGER,
    room_id INTEGER,
    name VARCHAR(30),
    price NUMERIC(12,2),
    start_time TIMESTAMPTZ,
    end_time TIMESTAMPTZ,
    duration INTEGER,
    total_participant INTEGER,
    snack_id INTEGER,
    snack_name VARCHAR(30),
    snack_price NUMERIC(12,2),
    total_snack NUMERIC(12,2),
    total_room NUMERIC(12,2),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (reservation_id) REFERENCES public.reservations(id) ON DELETE CASCADE,
    FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE SET NULL,
    FOREIGN KEY (snack_id) REFERENCES public.snacks(id) ON DELETE SET NULL
);

CREATE INDEX idx_reservation_details_reservation_id ON public.reservation_details (reservation_id);
CREATE INDEX idx_reservation_details_room_id ON public.reservation_details (room_id);
CREATE INDEX idx_reservation_details_snack_id ON public.reservation_details (snack_id);