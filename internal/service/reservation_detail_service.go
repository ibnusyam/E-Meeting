package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"errors"
)

type ReservationDetailService interface {
	GetReservationByID(id int, userID int, role string) (*model.ReservationData, error)
}

type reservationDetailService struct {
	repo repository.ReservationDetailRepository
}

func NewReservationDetailService(repo repository.ReservationDetailRepository) ReservationDetailService {
	return &reservationDetailService{repo: repo}
}

func (s *reservationDetailService) GetReservationByID(id int, userID int, role string) (*model.ReservationData, error) {
	// Check authorization
	if role == "customer" {
		// Customer can only see their own reservation
		ownerID, err := s.repo.GetReservationOwnerID(id)
		if err != nil {
			return nil, err
		}

		if ownerID != userID {
			return nil, errors.New("unauthorized to access this reservation")
		}
	}
	// Admin can see all reservations, so no additional check needed

	// Get reservation data
	reservationData, err := s.repo.GetReservationByID(id)
	if err != nil {
		return nil, err
	}

	return reservationData, nil
}
