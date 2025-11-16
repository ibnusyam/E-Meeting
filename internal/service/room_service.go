package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"database/sql"
	"errors"
	"net/url"
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

func (s *RoomService) UpdateRoom(id int, req model.CreateRoomRequest) error {

	// Validate type
	validTypes := map[string]bool{
		"small":  true,
		"medium": true,
		"large":  true,
	}

	if !validTypes[req.Type] {
		return errors.New("room type is not valid")
	}

	// validate capacity > 0
	if req.Capacity <= 0 {
		return errors.New("capacity must be larger more than 0")
	}

	// validate image URL
	_, err := url.ParseRequestURI(req.ImageURL)
	if err != nil {
		return errors.New("url not found")
	}

	// mapping request → model
	room := model.Room{
		Name:      req.Name,
		Price:     req.PricePerHour,
		ImagesUrl: req.ImageURL,
		Capacity:  req.Capacity,
		Type:      req.Type,
	}

	// call repository
	err = s.Repo.UpdateRoom(id, room)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows // supaya handler bisa return 404
		}
		return err
	}

	return nil
}

// cek jika room sudah ada reservasi atau belum
var ErrRoomUsed = errors.New("cannot delete rooms. room has reservation")

func (s *RoomService) DeleteRoom(id int) error {
	// cek apakah room sedang digunakan
	used, err := s.Repo.IsRoomUsed(id)
	if err != nil {
		return err
	}
	if used {
		return ErrRoomUsed
	}

	// delete room
	err = s.Repo.DeleteRoom(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		return err
	}

	return nil
}
