// repository/reservation_repository.go
package repository

import (
	"E-Meeting/model"
	"database/sql"
	"errors"
)

type ReservationDetailRepository interface {
	GetReservationByID(id int) (*model.ReservationData, error)
	GetReservationOwnerID(id int) (int, error)
}

type reservationDetailRepository struct {
	db *sql.DB
}

func NewReservationDetailRepository(db *sql.DB) ReservationDetailRepository {
	return &reservationDetailRepository{db: db}
}

func (r *reservationDetailRepository) GetReservationOwnerID(id int) (int, error) {
	var userID int
	query := `SELECT user_id FROM reservations WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("reservation not found")
		}
		return 0, err
	}
	return userID, nil
}

func (r *reservationDetailRepository) GetReservationByID(id int) (*model.ReservationData, error) {
	// Get personal data
	var personalData model.PersonalData
	personalQuery := `
	SELECT name, reserver_phone_number, company_name
	FROM reservations 
	WHERE id = $1
	`

	err := r.db.QueryRow(personalQuery, id).Scan(
		&personalData.Name,
		&personalData.PhoneNumber,
		&personalData.Company,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("reservation not found")
		}
		return nil, err
	}

	// Get rooms data
	roomsQuery := `
		SELECT 
			r.name,
			r.price,
			COALESCE(r.images_url, ''),
			r.capacity,
			r.type,
			rr.total_snack,
			rr.total_room,
			rr.start_time,
			rr.end_time,
			rr.duration,
			rr.total_participant,
			s.id,
			s.name,
			rr.total_participant,
			s.price,
			s.category
		FROM reservation_details rr
		JOIN rooms r ON rr.room_id = r.id
		LEFT JOIN snacks s ON rr.snack_id = s.id
		WHERE rr.reservation_id = $1
	`
	rows, err := r.db.Query(roomsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []model.RoomReservationDetail
	var subTotalSnack, subTotalRoom float64

	for rows.Next() {
		var room model.RoomReservationDetail
		var totalSnack, totalRoom float64

		// Variables for nullable snack fields
		var snackID sql.NullInt64
		var snackName, snackCategory sql.NullString
		var snackUnit sql.NullInt64 // Changed to NullInt64 if Unit is int in model
		var snackPrice sql.NullFloat64

		err := rows.Scan(
			&room.Name,
			&room.PricePerHour,
			&room.ImageURL,
			&room.Capacity,
			&room.Type,
			&totalSnack,
			&totalRoom,
			&room.StartTime,
			&room.EndTime,
			&room.Duration,
			&room.Participant,
			&snackID,
			&snackName,
			&snackUnit,
			&snackPrice,
			&snackCategory,
		)
		if err != nil {
			return nil, err
		}

		// Set total for this room
		room.TotalSnack = totalSnack
		room.TotalRoom = totalRoom

		// If snack exists, add it to room
		if snackID.Valid {
			room.Snack = model.Snack{
				ID:       int(snackID.Int64),
				Name:     snackName.String,
				Unit:     int(snackUnit.Int64), // Convert to int
				Price:    snackPrice.Float64,
				Category: snackCategory.String,
			}
		}

		subTotalSnack += totalSnack
		subTotalRoom += totalRoom
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(rooms) == 0 {
		return nil, errors.New("no rooms found for this reservation")
	}

	reservationData := &model.ReservationData{
		Rooms:         rooms,
		PersonalData:  personalData,
		SubTotalSnack: subTotalSnack,
		SubTotalRoom:  subTotalRoom,
		Total:         subTotalSnack + subTotalRoom,
	}

	return reservationData, nil
}
