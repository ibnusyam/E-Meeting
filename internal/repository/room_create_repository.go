package repository

import (
	"E-Meeting/model"
	"database/sql"
)

type RoomCreateRepository interface {
	CreateRoom(room model.Room) error
}

type roomCreateRepository struct {
	db *sql.DB
}

func NewRoomCreateRepository(db *sql.DB) RoomCreateRepository {
	return &roomCreateRepository{db}
}

func (r *roomCreateRepository) CreateRoom(room model.Room) error {
	query := `
		INSERT INTO rooms (name, capacity, price, type, images_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`

	_, err := r.db.Exec(query,
		room.Name,
		room.Capacity,
		room.Price,
		room.Type,
		room.ImagesUrl,
	)

	return err
}
