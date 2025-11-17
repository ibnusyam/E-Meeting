package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"context"
	"errors"
	"math"
)

type ReservationHistoryService interface {
	GetHistory(ctx context.Context, filter model.ReservationHistoryFilter) (*model.ReservationHistoryResponse, error)
}

type reservationHistoryService struct {
	repo repository.ReservationHistoryRepository
}

func NewReservationHistoryService(repo repository.ReservationHistoryRepository) ReservationHistoryService {
	return &reservationHistoryService{repo: repo}
}

func (s *reservationHistoryService) GetHistory(ctx context.Context, filter model.ReservationHistoryFilter) (*model.ReservationHistoryResponse, error) {
	// Validate room type if provided
	if filter.Type != "" {
		validTypes := map[string]bool{
			"small":  true,
			"medium": true,
			"large":  true,
		}
		if !validTypes[filter.Type] {
			return nil, errors.New("room type is not valid")
		}
	}

	// Set default pagination if not provided
	if filter.Page == 0 {
		filter.Page = 1
	}

	if filter.PageSize == 0 {
		filter.PageSize = 10
	}

	// Get histories
	histories, totalData, err := s.repo.GetHistory(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Get room IDs for fetching rooms
	if len(histories) > 0 {
		reservationIDs := make([]int, len(histories))
		for i, h := range histories {
			reservationIDs[i] = h.ID
		}

		// Get rooms for all reservations
		roomsMap, err := s.repo.GetHistoryRooms(ctx, reservationIDs)
		if err != nil {
			return nil, err
		}

		// Assign rooms to histories
		for i := range histories {
			if rooms, ok := roomsMap[histories[i].ID]; ok {
				histories[i].Rooms = rooms
			} else {
				histories[i].Rooms = []model.ReservationHistoryRoom{}
			}
		}
	}

	// Calculate total pages
	totalPage := int(math.Ceil(float64(totalData) / float64(filter.PageSize)))

	response := &model.ReservationHistoryResponse{
		Message:   "success",
		Data:      histories,
		Page:      filter.Page,
		PageSize:  filter.PageSize,
		TotalPage: totalPage,
		TotalData: totalData,
	}

	return response, nil
}
