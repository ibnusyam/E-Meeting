package repository

import (
	"E-Meeting/model"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type RoomRepository struct {
	DB *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{DB: db}
}

func (repo *RoomRepository) GetAllRoom(name, roomType string, capacity, page, pageSize int) ([]model.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	baseQuery := `SELECT id, name, capacity, price, type, COALESCE(images_url, ''), created_at, updated_at FROM rooms where 1=1`

	args := []interface{}{}

	argIndex := 1

	if name != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(name) LIKE LOWER($%d)", argIndex)
		args = append(args, name)
		argIndex++
	}

	if roomType != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(type) LIKE ($%d)", argIndex)
		args = append(args, roomType)
		argIndex++
	}

	if capacity > 0 {
		baseQuery += fmt.Sprintf(" AND capacity >= $%d", argIndex)
		args = append(args, capacity)
		argIndex++
	}

	// Hitung total data (tanpa LIMIT)
	countQuery := "SELECT COUNT(*) FROM (" + baseQuery + ") AS count_query"
	var total int
	if err := repo.DB.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, err
	}

	// Pagination
	offset := (page - 1) * pageSize
	baseQuery += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pageSize, offset)

	rows, err := repo.DB.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []model.Room

	for rows.Next() {
		var room model.Room

		err := rows.Scan(&room.ID, &room.Name, &room.Capacity, &room.Price, &room.Type, &room.ImagesUrl, &room.CreatedAt, &room.UpdatedAt)

		if err != nil {
			log.Printf("Error ambil data rooms: %v\n", err)

			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
