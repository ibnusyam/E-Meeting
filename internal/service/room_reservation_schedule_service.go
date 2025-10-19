package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
)

type RoomReservationScheduleService struct {
	Repo *repository.RoomReservationScheduleRepository
}

func NewRoomReservationScheduleService(repo *repository.RoomReservationScheduleRepository) *RoomReservationScheduleService {
	return &RoomReservationScheduleService{Repo: repo}
}

func (s *RoomReservationScheduleService) GetRoomReservationSchedules(roomID int, startDate string) (*model.RoomReservationResponse, error) {
	return s.Repo.GetRoomReservationSchedules(roomID, startDate)
}
