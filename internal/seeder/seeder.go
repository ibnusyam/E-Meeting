package seeder

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Main function to run all seeders in the correct order
func Run(db *sql.DB) {
	// 1. Hash the password once
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password for seeder: %v", err)
	}
	passwordStr := string(hashedPassword)

	// 2. Run seeders for tables without dependencies first
	seedUsers(db, passwordStr)
	seedRooms(db)
	seedSnacks(db)

	// 3. Run seeders for tables that have dependencies
	seedReservations(db)
	seedReservationDetails(db)

	// 4. Reset sequences to avoid ID conflicts
	resetSequences(db)
}

// seedUsers seeds the users table
func seedUsers(db *sql.DB, passwordStr string) {
	_, err := db.Exec(`
		INSERT INTO public.users (id, role, username, email, password, phone_number, language, status_user) VALUES
		(1, 'admin', 'admin', 'admin@gmail.com', $1, '081234567890', 'English', 'active'),
		(2, 'customer', 'costumer', 'customer@example.com', $1, '085123456789', 'Indonesia', 'active'),
		(3, 'customer', 'ibnu', 'ibnu@example.com', $1, '081223344556', 'Indonesia', 'active'),
		(4, 'customer', 'alex', 'alex@example.com', $1, '081223341236', 'Indonesia', 'active'),
		(5, 'customer', 'novita', 'novita@example.com', $1, '08122335556', 'Indonesia', 'active')
		ON CONFLICT (id) DO NOTHING;
	`, passwordStr)
	logError("users", err)
}

// seedRooms seeds the rooms table
func seedRooms(db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO public.rooms (id, name, capacity, price, type) VALUES
		(1, 'Aster Room', 10, 100000.00, 'small'),
		(2, 'Bluebell Room', 25, 250000.00, 'medium'),
		(3, 'Camellia Room', 50, 500000.00, 'large')
		ON CONFLICT (id) DO NOTHING;
	`)
	logError("rooms", err)
}

// seedSnacks seeds the snacks table
func seedSnacks(db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO public.snacks (id, category, name, price) VALUES
		(1, 'Coffee Break', 'Coffee Break Package 1', 20000.00),
		(2, 'Coffee Break', 'Coffee Break Package 2', 50000.00),
		(3, 'Lunch', 'Lunch Package 1', 35000.00),
		(4, 'Lunch', 'Lunch Package 2', 60000.00)
		ON CONFLICT (id) DO NOTHING;
	`)
	logError("snacks", err)
}

// seedReservations seeds the reservations table
func seedReservations(db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO public.reservations (id, user_id, name, reserver_phone_number, company_name, note, status_reservation, sub_total_snack, sub_total_rooms, total) VALUES
		(1, 3, 'Ibnu', '08122344556', 'PT Maju Jaya', 'Lorem Ipsum has been the industry''s standard dummy text...', 'paid', 280000.00, 200000.00, 480000.00),
		(2, 3, 'Ibnu', '08122344556', 'Organisasi Muslim Pusat', 'Catatan untuk rapat organisasi.', 'booked', 0.00, 400000.00, 400000.00),
		(3, 5, 'Novita', '08122355556', 'Startup ABC', NULL, 'canceled', 0.00, 0.00, 200000.00)
		ON CONFLICT (id) DO NOTHING;
	`)
	logError("reservations", err)
}

// seedReservationDetails seeds the reservation_details table
func seedReservationDetails(db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO public.reservation_details (id, reservation_id, room_id, name, price, start_time, end_time, duration, total_participant, snack_id, snack_name, snack_price, total_snack, total_room) VALUES
		(1, 1, 1, 'Rapat Tim PT Maju Jaya', 100000.00, '2025-10-01 10:00:00+07', '2025-10-01 12:00:00+07', 2, 8, 3, 'Lunch Package 1', 35000.00, 280000.00, 200000.00),
		(2, 2, 2, 'Rapat Organisasi Muslim Pusat', 250000.00, '2025-10-30 13:00:00+07', '2025-10-30 15:00:00+07', 2, 15, NULL, NULL, NULL, NULL, 400000.00)
		ON CONFLICT (id) DO NOTHING;
	`)
	logError("reservation_details", err)
}

// resetSequences updates the auto-increment counters
func resetSequences(db *sql.DB) {
	db.Exec(`SELECT setval('public.users_id_seq', COALESCE((SELECT MAX(id) FROM public.users), 1));`)
	db.Exec(`SELECT setval('public.rooms_id_seq', COALESCE((SELECT MAX(id) FROM public.rooms), 1));`)
	db.Exec(`SELECT setval('public.snacks_id_seq', COALESCE((SELECT MAX(id) FROM public.snacks), 1));`)
	db.Exec(`SELECT setval('public.reservations_id_seq', COALESCE((SELECT MAX(id) FROM public.reservations), 1));`)
	db.Exec(`SELECT setval('public.reservation_details_id_seq', COALESCE((SELECT MAX(id) FROM public.reservation_details), 1));`)
	log.Println("Database sequences have been reset.")
}

// Helper function to log errors
func logError(tableName string, err error) {
	if err != nil {
		log.Printf("Failed to run seeder for %s: %v", tableName, err)
	} else {
		log.Printf("Seeder for %s ran successfully.", tableName)
	}
}
