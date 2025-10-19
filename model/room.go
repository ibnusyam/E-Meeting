package model

import "time"

type Room struct {
	ID        int
	Name      string
	Capacity  int
	Price     float64
	Type      string
	ImagesUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// RoomRequest untuk Swagger documentation
type RoomRequest struct {
	ID          int       `json:"id" example:"1"`
	StartTime   time.Time `json:"startTime" example:"2025-10-17T12:00:00Z"`
	EndTime     time.Time `json:"endTime" example:"2025-10-17T14:00:00Z"`
	Participant int       `json:"participant" example:"2"`
	SnackID     int       `json:"snackID" example:"1"`
}
