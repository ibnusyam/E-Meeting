package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"errors"
	"net/url"
)

type RoomCreateService interface {
	CreateRoom(req model.CreateRoomRequest) error
}

type roomCreateService struct {
	roomRepo repository.RoomCreateRepository
}

func NewRoomCreateService(roomRepo repository.RoomCreateRepository) RoomCreateService {
	return &roomCreateService{roomRepo}
}

func (s *roomCreateService) CreateRoom(req model.CreateRoomRequest) error {

	// Validate type
	validTypes := map[string]bool{
		"small":  true,
		"medium": true,
		"large":  true,
	}

	if !validTypes[req.Type] {
		return errors.New("room type is not valid")
	}

	// capacity must > 0
	if req.Capacity <= 0 {
		return errors.New("capacity must be larger more than 0")
	}

	// validate url
	_, err := url.ParseRequestURI(req.ImageURL)
	if err != nil {
		return errors.New("url not found")
	}

	room := model.Room{
		Name:      req.Name,
		Price:     req.PricePerHour,
		ImagesUrl: req.ImageURL,
		Capacity:  req.Capacity,
		Type:      req.Type,
	}

	return s.roomRepo.CreateRoom(room)
}
