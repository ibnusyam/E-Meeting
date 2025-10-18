// service/reservation_service.go
package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"errors"
	"strconv"
)

type ReservationService struct {
	Repo *repository.ReservationRepository
}

func NewReservationService(repo *repository.ReservationRepository) *ReservationService {
	return &ReservationService{Repo: repo}
}

func (s *ReservationService) CreateReservation(req model.ReservationRequest) error {
	// Validate rooms availability
	for _, room := range req.Rooms {
		exists := s.Repo.CheckRoomExists(room.ID)
		if !exists {
			return errors.New("room not found")
		}

		isBooked := s.Repo.CheckRoomBooked(room.ID, room.StartTime, room.EndTime)

		if isBooked {
			return errors.New("room has been booked")
		}
	}

	UserID, err := strconv.Atoi(req.UserID)
	if err != nil {
		return errors.New("Error when converting userID into string")
	}

	// Build reservation details with calculations
	reservationDetails := make([]model.ReservationDetail, 0)
	var subTotalSnack, subTotalRooms float64

	for _, room := range req.Rooms {
		// Get room details
		roomDetail, err := s.Repo.GetRoomDetail(strconv.Itoa(room.ID))
		if err != nil {
			return errors.New("failed to get room detail")
		}

		// Calculate duration in hours
		duration := int(room.EndTime.Sub(room.StartTime).Hours())

		// Calculate room total
		totalRoom := roomDetail.Price * float64(duration)

		// Get snack details if snackID provided
		var snackName string
		var snackPrice, totalSnack float64
		if room.SnackID > 0 {
			snackDetail, err := s.Repo.GetSnackDetail(strconv.Itoa(room.SnackID))
			if err != nil {
				return errors.New("failed to get snack detail")
			}
			snackName = snackDetail.Name
			snackPrice = snackDetail.Price
			totalSnack = snackPrice * float64(room.Participant)
		}

		// Create reservation detail
		detail := model.ReservationDetail{
			RoomID:           room.ID,
			Name:             roomDetail.Name,
			Price:            roomDetail.Price,
			StartTime:        room.StartTime,
			EndTime:          room.EndTime,
			Duration:         duration,
			TotalParticipant: room.Participant,
			SnackID:          room.SnackID,
			SnackName:        snackName,
			SnackPrice:       snackPrice,
			TotalSnack:       totalSnack,
			TotalRoom:        totalRoom,
		}

		reservationDetails = append(reservationDetails, detail)
		subTotalSnack += totalSnack
		subTotalRooms += totalRoom
	}

	// Calculate grand total
	total := subTotalSnack + subTotalRooms

	// Build reservation object
	reservation := model.Reservation{
		UserID:              UserID,
		Name:                req.Name,
		ReserverPhoneNumber: req.PhoneNumber,
		CompanyName:         req.Company,
		Note:                req.Notes,
		StatusReservation:   "booked",
		SubTotalSnack:       subTotalSnack,
		SubTotalRooms:       subTotalRooms,
		Total:               total,
		ReservationDetails:  reservationDetails,
	}

	// Insert to database
	err = s.Repo.InsertReservation(reservation)
	if err != nil {
		return err
	}

	return nil
}
