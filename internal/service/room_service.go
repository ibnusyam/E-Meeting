package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
)

type RoomService struct {
	Repo *repository.RoomRepository
}

func NewRoomService(repo *repository.RoomRepository) *RoomService {
	return &RoomService{Repo: repo}
}

func (s *RoomService) GetAllRooms(name, roomType string, capacity, page, pageSize int) ([]model.Room, error) {

	rooms, err := s.Repo.GetAllRoom(name, roomType, capacity, page, pageSize)
	if err != nil {
		return nil, err
	}

	if rooms == nil {
		return []model.Room{}, nil
	}

	return rooms, err
}
