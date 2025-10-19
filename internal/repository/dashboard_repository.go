package repository

import (
	"database/sql"
	"time"
)

type DashboardRepository interface {
	GetDashboardData(startDate, endDate time.Time) (DashboardData, error)
}

type dashboardRepository struct {
	db *sql.DB
}

func NewDashboardRepository(db *sql.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

type RoomStat struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Omzet             float64 `json:"omzet"`
	PercentageOfUsage float64 `json:"percentageOfUsage"`
}

type DashboardData struct {
	TotalRoom        int        `json:"totalRoom"`
	TotalVisitor     int        `json:"totalVisitor"`
	TotalReservation int        `json:"totalReservation"`
	TotalOmzet       float64    `json:"totalOmzet"`
	Rooms            []RoomStat `json:"rooms"`
}

func (r *dashboardRepository) GetDashboardData(startDate, endDate time.Time) (DashboardData, error) {
	var data DashboardData

	// Hitung total room
	err := r.db.QueryRow(`SELECT COUNT(*) FROM rooms`).Scan(&data.TotalRoom)
	if err != nil {
		return data, err
	}

	// Hitung total visitor, total reservation, total omzet (paid transaction)
	err = r.db.QueryRow(`
		SELECT 
			COALESCE(SUM(res_det.total_participant ),0),
			COUNT(*),
			COALESCE(SUM(total),0)
		FROM reservations res
		left join reservation_details res_det on res.id=res_det.reservation_id
		WHERE room_id is not null and status_reservation = 'paid' AND res_det.created_at BETWEEN $1 AND $2
	`, startDate, endDate).Scan(&data.TotalVisitor, &data.TotalReservation, &data.TotalOmzet)
	if err != nil {
		return data, err
	}

	// Ambil omzet dan usage per room
	rows, err := r.db.Query(`
		SELECT r.id, r.name, 
			COALESCE(SUM(res.total),0) as omzet,
			ROUND((COUNT(res.id)::decimal / NULLIF((SELECT COUNT(*) FROM reservations WHERE status_reservation = 'paid'), 0)) * 100, 2) as usage_percentage
		FROM rooms r 
		LEFT JOIN reservation_details res_det ON r.id = res_det.room_id 
		left join reservations res on res_det.reservation_id = res.id AND res.created_at BETWEEN $1 AND $2 AND res.status_reservation = 'paid'
		GROUP BY r.id, r.name
	`, startDate, endDate)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var room RoomStat
		if err := rows.Scan(&room.ID, &room.Name, &room.Omzet, &room.PercentageOfUsage); err != nil {
			return data, err
		}
		data.Rooms = append(data.Rooms, room)
	}

	return data, nil
}
