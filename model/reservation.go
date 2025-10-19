// internal/model/reservation.go
package model

import "time"

// Reservation represents the main reservation record in reservations table
type Reservation struct {
	ID                  int
	UserID              int
	Name                string
	ReserverPhoneNumber string
	CompanyName         string
	Note                string
	StatusReservation   string
	SubTotalSnack       float64
	SubTotalRooms       float64
	Total               float64
	ReservationDetails  []ReservationDetail // Added for inserting details
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

// ReservationDetail represents each room booking detail in reservation_details table
type ReservationDetail struct {
	ID               int
	ReservationID    int
	RoomID           int
	Name             string
	Price            float64
	StartTime        time.Time
	EndTime          time.Time
	Duration         int
	TotalParticipant int
	SnackID          int
	SnackName        string
	SnackPrice       float64
	TotalRoom        float64
	TotalSnack       float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// ReservationRequest is used for API input (Swagger documentation)
type ReservationRequest struct {
	UserID      string        `json:"userID" example:"1"`
	Name        string        `json:"name" example:"John Doe"`
	PhoneNumber string        `json:"phoneNumber" example:"081234567890"`
	Company     string        `json:"company" example:"ABC Company"`
	Notes       string        `json:"notes" example:"Special request"`
	Rooms       []RoomRequest `json:"rooms"`
}

// // RoomRequest is used for API input (room details in request)
// type RoomRequest struct {
// 	ID          int    `json:"id" example:"1"`
// 	StartTime   string `json:"startTime" example:"2025-10-01 10:00:00"`
// 	EndTime     string `json:"endTime" example:"2025-10-01 12:00:00"`
// 	Participant int    `json:"participant" example:"8"`
// 	SnackID     int    `json:"snackID" example:"3"`
// }

// // Room represents room master data
// type Room struct {
// 	ID        int
// 	Name      string
// 	Capacity  int
// 	Price     float64
// 	Type      string
// 	ImagesUrl string
// }

// // Snack represents snack master data
// type Snack struct {
// 	ID    int
// 	Name  string
// 	Price float64
// }

type ReservationDetailDTO struct {
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
