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

type RoomReservationDetail struct {
	Name         string    `json:"name"`
	PricePerHour float64   `json:"pricePerHour"`
	ImageURL     string    `json:"imageURL"`
	Capacity     int       `json:"capacity"`
	Type         string    `json:"type"`
	TotalSnack   float64   `json:"totalSnack"`
	TotalRoom    float64   `json:"totalRoom"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	Duration     int       `json:"duration"`
	Participant  int       `json:"participant"`
	Snack        Snack     `json:"snack,omitempty"`
}

// RoomRequest untuk Swagger documentation
type RoomRequest struct {
	ID          int       `json:"id" example:"1"`
	StartTime   time.Time `json:"startTime" example:"2025-10-17T12:00:00Z"`
	EndTime     time.Time `json:"endTime" example:"2025-10-17T14:00:00Z"`
	Participant int       `json:"participant" example:"2"`
	SnackID     int       `json:"snackID" example:"1"`
}
