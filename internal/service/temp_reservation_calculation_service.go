package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"context"
	"errors"
)

type ReservationCalculationService interface {
	CalculateReservation(ctx context.Context, req model.ReservationCalculationRequest) (*model.ReservationCalculationResponse, error)
}

type reservationCalculationService struct {
	repo repository.ReservationCalculationRepository
}

func NewReservationCalculationService(repo repository.ReservationCalculationRepository) ReservationCalculationService {
	return &reservationCalculationService{repo: repo}
}

func (s *reservationCalculationService) CalculateReservation(ctx context.Context, req model.ReservationCalculationRequest) (*model.ReservationCalculationResponse, error) {
	// Validasi input
	if req.RoomID == 0 || req.StartTime.IsZero() || req.EndTime.IsZero() {
		return nil, errors.New("room_id, startTime, and endTime are required")
	}

	if req.StartTime.After(req.EndTime) || req.StartTime.Equal(req.EndTime) {
		return nil, errors.New("invalid time range")
	}

	if req.Participant <= 0 {
		return nil, errors.New("participant must be greater than 0")
	}

	// Get room data
	room, err := s.repo.GetRoomByID(ctx, req.RoomID)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, errors.New("room not found")
	}

	// Check room capacity
	isCapacityValid, err := s.repo.CheckRoomCapacity(ctx, req.RoomID, req.Participant)
	if err != nil {
		return nil, err
	}
	if !isCapacityValid {
		return nil, errors.New("over capacity")
	}

	// Check room availability
	isAvailable, err := s.repo.CheckRoomAvailability(ctx, req.RoomID, req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}
	if !isAvailable {
		return nil, errors.New("booking bentrok")
	}

	// Calculate duration in hours
	duration := int(req.EndTime.Sub(req.StartTime).Hours())
	if duration == 0 {
		duration = 1 // minimum 1 hour
	}

	// Calculate room subtotal
	subTotalRoom := room.PricePerHour * float64(duration)

	// Get snack data and calculate snack subtotal
	var snack model.ReservationCalculationSnack
	var subTotalSnack float64 = 0

	if req.SnackID > 0 {
		snackData, err := s.repo.GetSnackByID(ctx, req.SnackID)
		if err != nil {
			return nil, err
		}
		if snackData != nil {
			snack = *snackData
			subTotalSnack = snack.Price * float64(req.Participant)
		}
	}

	// Build response
	roomData := model.ReservationCalculationRooms{
		Name:          room.Name,
		PricePerHour:  room.PricePerHour,
		ImagesUrl:     room.ImagesUrl,
		Capacity:      room.Capacity,
		Type:          room.Type,
		SubTotalSnack: subTotalSnack,
		SubTotalRooms: subTotalRoom,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		Duration:      duration,
		Participant:   req.Participant,
		Snack:         snack,
	}

	personalData := model.ReservationCalculationPersonalData{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Company:     req.Company,
	}

	total := subTotalRoom + subTotalSnack

	response := &model.ReservationCalculationResponse{
		Message: "Calculation successful",
		Data: model.ReservationCalculationData{
			Rooms:        []model.ReservationCalculationRooms{roomData},
			PersonalData: personalData,
		},
		SubTotalRoom:  subTotalRoom,
		SubTotalSnack: subTotalSnack,
		Total:         total,
	}

	return response, nil
}
