package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"errors"
)

type RoomCreateService interface {
	CreateRoom(req *model.CreateRoomRequest, imageURL string) error
}

type roomCreateService struct {
	roomRepo repository.RoomCreateRepository
}

func NewRoomCreateService(roomRepo repository.RoomCreateRepository) RoomCreateService {
	return &roomCreateService{
		roomRepo: roomRepo,
	}
}

func (s *roomCreateService) CreateRoom(req *model.CreateRoomRequest, imageURL string) error {

	// valid type
	validTypes := map[string]bool{"small": true, "medium": true, "large": true}
	if !validTypes[req.Type] {
		return errors.New("invalid room type")
	}

	// capacity > 0
	if req.Capacity <= 0 {
		return errors.New("capacity must be larger than 0")
	}

	// create room model
	room := model.Room{
		Name:      req.Name,
		Price:     req.PricePerHour,
		Capacity:  req.Capacity,
		Type:      req.Type,
		ImagesUrl: imageURL,
	}

	return s.roomRepo.CreateRoom(room)
}
