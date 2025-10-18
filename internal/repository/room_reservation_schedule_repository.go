package repository

import (
	"E-Meeting/model"
	"database/sql"
)

type RoomReservationScheduleRepository struct {
	DB *sql.DB
}

func NewRoomReservationScheduleRepository(db *sql.DB) *RoomReservationScheduleRepository {
	return &RoomReservationScheduleRepository{DB: db}
}

func (r *RoomReservationScheduleRepository) GetRoomReservationSchedules(roomID int, startDate string) (*model.RoomReservationResponse, error) {
	query := `SELECT rd.start_time, rd.end_time, r.name AS room_name, res.status_reservation
		FROM reservation_details rd
		JOIN reservations res ON rd.reservation_id = res.id
		JOIN rooms r ON rd.room_id = r.id
		WHERE rd.room_id = $1 
		  AND DATE(rd.start_time) = $2
		ORDER BY rd.start_time ASC`

	rows, err := r.DB.Query(query, roomID, startDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response model.RoomReservationResponse
	response.RoomName = ""
	response.TotalBooked = 0

	for rows.Next() {
		var startTime, endTime, status, roomName string
		if err := rows.Scan(&startTime, &endTime, &roomName, &status); err != nil {
			return nil, err
		}

		if response.RoomName == "" {
			response.RoomName = roomName
		}

		response.Schedules = append(response.Schedules, struct {
			StartTime string `json:"startTime"`
			EndTime   string `json:"endTime"`
			Status    string `json:"status"`
		}{
			StartTime: startTime,
			EndTime:   endTime,
			Status:    status,
		})

		response.TotalBooked++
	}

	return &response, nil
}
