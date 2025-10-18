// repository/reservation_repository.go
package repository

import (
	"E-Meeting/model"
	"database/sql"
	"log"
	"time"
)

type ReservationRepository struct {
	DB *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{DB: db}
}

func (r *ReservationRepository) GetRoomDetail(roomID string) (model.Room, error) {
	var room model.Room
	query := `SELECT id, name, capacity, price
	          FROM rooms 
	          WHERE id = $1 
	          ORDER BY updated_at DESC 
	          LIMIT 1`
	err := r.DB.QueryRow(query, roomID).Scan(
		&room.ID, &room.Name, &room.Capacity, &room.Price,
	)
	return room, err
}

func (r *ReservationRepository) GetSnackDetail(snackID string) (model.Snack, error) {
	var snack model.Snack
	query := `SELECT id, name, price 
	          FROM snacks 
	          WHERE id = $1 
	          ORDER BY updated_at DESC 
	          LIMIT 1`
	err := r.DB.QueryRow(query, snackID).Scan(&snack.ID, &snack.Name, &snack.Price)
	return snack, err
}

func (r *ReservationRepository) CheckRoomExists(roomID int) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM rooms WHERE id = $1)"
	err := r.DB.QueryRow(query, roomID).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *ReservationRepository) CheckRoomBooked(roomID int, start, end time.Time) bool {
	var count int
	query := `
		SELECT COUNT(*) FROM reservation_details
		WHERE room_id = $1
		AND (
			($2 < end_time)
			AND ($3 > start_time)
		)`

	err := r.DB.QueryRow(query, roomID, start, end).Scan(&count)
	if err != nil {
		log.Printf("CheckRoomBooked error: %v", err)
		return false
	}
	return count > 0
}

func (r *ReservationRepository) InsertReservation(reservation model.Reservation) error {
	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Begin transaction error: %v", err)
		return err
	}
	defer tx.Rollback()

	now := time.Now()

	// Insert into reservations table
	queryReservation := `
		INSERT INTO reservations (
			user_id, name, reserver_phone_number, company_name, note, 
			status_reservation, sub_total_snack, sub_total_rooms, total, 
			created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`

	var reservationID int
	err = tx.QueryRow(
		queryReservation,
		reservation.UserID,
		reservation.Name,
		reservation.ReserverPhoneNumber,
		reservation.CompanyName,
		reservation.Note,
		reservation.StatusReservation,
		reservation.SubTotalSnack,
		reservation.SubTotalRooms,
		reservation.Total,
		now,
		now,
	).Scan(&reservationID)

	if err != nil {
		log.Printf("Insert reservation error: %v", err)
		return err
	}

	log.Printf("Reservation inserted with ID: %d", reservationID)

	// Insert into reservation_details table
	queryDetail := `
		INSERT INTO reservation_details (
			reservation_id, room_id, name, price, start_time, end_time, 
			duration, total_participant, snack_id, snack_name, snack_price, 
			total_snack, total_room, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`

	for _, detail := range reservation.ReservationDetails {
		log.Printf("Inserting reservation detail: %+v", detail)

		// Handle NULL values for snack fields
		var snackID interface{}
		var snackName interface{}
		var snackPrice interface{}

		if detail.SnackID > 0 {
			snackID = detail.SnackID
			snackName = detail.SnackName
			snackPrice = detail.SnackPrice
		} else {
			snackID = nil
			snackName = nil
			snackPrice = nil
		}

		_, err = tx.Exec(
			queryDetail,
			reservationID,
			detail.RoomID,
			detail.Name,
			detail.Price,
			detail.StartTime,
			detail.EndTime,
			detail.Duration,
			detail.TotalParticipant,
			snackID,
			snackName,
			snackPrice,
			detail.TotalSnack,
			detail.TotalRoom,
			now,
			now,
		)
		if err != nil {
			log.Printf("Insert reservation detail error: %v", err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Commit transaction error: %v", err)
		return err
	}

	log.Printf("Reservation and details successfully inserted")
	return nil
}
