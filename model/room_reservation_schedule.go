package model

import "time"

type ReservationDetail struct {
	ID               int       `json:"id"`
	ReservationID    int       `json:"reservation_id"`
	RoomID           int       `json:"room_id"`
	Name             string    `json:"name"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	Status           string    `json:"status"`
	TotalParticipant int       `json:"total_participant"`
}

type RoomReservationResponse struct {
	RoomName    string `json:"roomName"`
	TotalBooked int    `json:"totalBooked"`
	Schedules   []struct {
		StartTime string `json:"startTime"`
		EndTime   string `json:"endTime"`
		Status    string `json:"status"`
	} `json:"schedules"`
}
