package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"database/sql"
	"errors"
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
