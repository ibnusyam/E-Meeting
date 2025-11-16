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

// RESERVATION CALCULATION ENDPOINT MODELS
type ReservationCalculationSnack struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Unit     string  `json:"unit"` // uncomment jika dipakai
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type ReservationCalculationRooms struct {
	Name          string                      `json:"name"`
	PricePerHour  float64                     `json:"pricePerHour"`
	ImagesUrl     string                      `json:"imageURL"`
	Capacity      int                         `json:"capacity"`
	Type          string                      `json:"type"`
	SubTotalSnack float64                     `json:"subTotalSnack"`
	SubTotalRooms float64                     `json:"subTotalRoom"`
	StartTime     time.Time                   `json:"startTime"`
	EndTime       time.Time                   `json:"endTime"`
	Duration      int                         `json:"duration"`
	Participant   int                         `json:"participant"`
	Snack         ReservationCalculationSnack `json:"snack"`
}

type ReservationCalculationRequest struct {
	RoomID      int       `json:"room_id" query:"room_id"`
	SnackID     int       `json:"snack_id" query:"snack_id"`
	StartTime   time.Time `json:"startTime" query:"startTime"`
	EndTime     time.Time `json:"endTime" query:"endTime"`
	Participant int       `json:"participant" query:"participant"`
	UserID      int       `json:"user_id" query:"user_id"`
	Name        string    `json:"name" query:"name"`
	PhoneNumber string    `json:"phoneNumber" query:"phoneNumber"`
	Company     string    `json:"company" query:"company"`
}

type ReservationCalculationPersonalData struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Company     string `json:"company"`
}

type ReservationCalculationData struct {
	Rooms        []ReservationCalculationRooms      `json:"rooms"`
	PersonalData ReservationCalculationPersonalData `json:"personalData"`
}

type ReservationCalculationResponse struct {
	Message       string                     `json:"message"`
	Data          ReservationCalculationData `json:"data"`
	SubTotalRoom  float64                    `json:"subTotalRoom"`
	SubTotalSnack float64                    `json:"subTotalSnack"`
	Total         float64                    `json:"total"`
}

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

// RESERVATION HISTORY
type ReservationHistory struct {
	ID            int                      `json:"id"`
	Name          string                   `json:"name"`
	PhoneNumber   float64                  `json:"phoneNumber"`
	Company       string                   `json:"company"`
	SubTotalSnack float64                  `json:"subTotalSnack"`
	SubTotalRoom  float64                  `json:"subTotalRoom"`
	Total         float64                  `json:"total"`
	Status        string                   `json:"status"`
	CreatedAt     time.Time                `json:"createdAt"`
	UpdatedAt     time.Time                `json:"updatedAt"`
	Rooms         []ReservationHistoryRoom `json:"rooms"`
	UserID        int                      `json:"userId"`
}

type ReservationHistoryRoom struct {
	ID         int     `json:"id"`
	Price      float64 `json:"price"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	TotalRoom  float64 `json:"totalRoom"`
	TotalSnack float64 `json:"totalSnack"`
}

type ReservationHistoryFilter struct {
	StartDate string
	EndDate   string
	Type      string
	Status    string
	Page      int
	PageSize  int
	Username  string
	UserID    int
}

type ReservationHistoryResponse struct {
	Message   string               `json:"message"`
	Data      []ReservationHistory `json:"data"`
	Page      int                  `json:"page"`
	PageSize  int                  `json:"pageSize"`
	TotalPage int                  `json:"totalPage"`
	TotalData int                  `json:"totalData"`
}

type PersonalData struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Company     string `json:"company"`
}

type ReservationData struct {
	Rooms         []RoomReservationDetail `json:"rooms"`
	PersonalData  PersonalData            `json:"personalData"`
	SubTotalSnack float64                 `json:"subTotalSnack"`
	SubTotalRoom  float64                 `json:"subTotalRoom"`
	Total         float64                 `json:"total"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ReservationResponse struct {
	Message string          `json:"message"`
	Data    ReservationData `json:"data"`
}
