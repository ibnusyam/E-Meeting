package repository

import (
	"E-Meeting/model"
	"context"
	"database/sql"
	"time"
)

type ReservationCalculationRepository interface {
	GetRoomByID(ctx context.Context, roomID int) (*model.ReservationCalculationRooms, error)
	GetSnackByID(ctx context.Context, snackID int) (*model.ReservationCalculationSnack, error)
	CheckRoomAvailability(ctx context.Context, roomID int, startTime, endTime time.Time) (bool, error)
	CheckRoomCapacity(ctx context.Context, roomID int, participant int) (bool, error)
}

type reservationCalculationRepository struct {
	db *sql.DB
}

func NewReservationCalculationRepository(db *sql.DB) ReservationCalculationRepository {
	return &reservationCalculationRepository{db: db}
}

func (r *reservationCalculationRepository) GetRoomByID(ctx context.Context, roomID int) (*model.ReservationCalculationRooms, error) {
	query := `
		SELECT id, name, price, COALESCE(images_url, ''), capacity, type
		FROM rooms
		WHERE id = $1
		ORDER BY updated_at DESC
		LIMIT 1;
	`

	var room model.ReservationCalculationRooms
	var id int

	err := r.db.QueryRowContext(ctx, query, roomID).Scan(
		&id,
		&room.Name,
		&room.PricePerHour,
		&room.ImagesUrl,
		&room.Capacity,
		&room.Type,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &room, nil
}

func (r *reservationCalculationRepository) GetSnackByID(ctx context.Context, snackID int) (*model.ReservationCalculationSnack, error) {
	query := `
		SELECT id, name, price, category
		FROM snacks
		WHERE id = $1
		ORDER BY updated_at DESC
		LIMIT 1;
	`

	var snack model.ReservationCalculationSnack

	err := r.db.QueryRowContext(ctx, query, snackID).Scan(
		&snack.ID,
		&snack.Name,
		&snack.Price,
		&snack.Category,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &snack, nil
}

func (r *reservationCalculationRepository) CheckRoomAvailability(ctx context.Context, roomID int, startTime, endTime time.Time) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM reservation_details a
		LEFT JOIN reservations b on a.reservation_id = b.id
		WHERE a.room_id = $1
		AND b.status_reservation NOT IN ('canceled')
		AND (
			(a.start_time <= $2 AND a.end_time > $2) OR
			(a.start_time < $3 AND a.end_time >= $3) OR
			(a.start_time >= $2 AND a.end_time <= $3)
		)
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, roomID, startTime, endTime).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (r *reservationCalculationRepository) CheckRoomCapacity(ctx context.Context, roomID int, participant int) (bool, error) {
	query := `
		SELECT capacity
		FROM rooms
		WHERE id = $1
		ORDER BY updated_at DESC
		LIMIT 1;
	`

	var capacity int
	err := r.db.QueryRowContext(ctx, query, roomID).Scan(&capacity)
	if err != nil {
		return false, err
	}

	return participant <= capacity, nil
}
