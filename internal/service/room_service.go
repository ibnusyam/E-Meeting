package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
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

func (s *RoomService) UpdateRoom(id int, req model.UpdateRoomRequest, imageFile *multipart.FileHeader) error {

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

	// ambil data lama dulu
	existing, err := s.Repo.GetRoomByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		return err
	}

	imagePath := existing.ImagesUrl // default gambar lama

	room := model.Room{
		Name:      req.Name,
		Price:     req.PricePerHour,
		ImagesUrl: imagePath,
		Capacity:  req.Capacity,
		Type:      req.Type,
	}

	// kalau user upload gambar baru
	if req.ImageFile != nil {
		src, err := req.ImageFile.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), req.ImageFile.Filename)
		dstPath := "public/uploads/rooms/" + fileName

		dst, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return err
		}

		room.ImagesUrl = "/static/rooms/" + fileName
	}

	err = s.Repo.UpdateRoom(id, room)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
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
