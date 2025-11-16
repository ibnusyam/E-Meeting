package repository

import (
	"E-Meeting/model"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type ReservationHistoryRepository interface {
	GetHistory(ctx context.Context, filter model.ReservationHistoryFilter) ([]model.ReservationHistory, int, error)
	GetHistoryRooms(ctx context.Context, reservationIDs []int) (map[int][]model.ReservationHistoryRoom, error)
}

type reservationHistoryRepository struct {
	db *sql.DB
}

func NewReservationHistoryRepository(db *sql.DB) ReservationHistoryRepository {
	return &reservationHistoryRepository{db: db}
}

func (r *reservationHistoryRepository) GetHistory(ctx context.Context, filter model.ReservationHistoryFilter) ([]model.ReservationHistory, int, error) {
	var conditions []string
	var args []interface{}
	argCount := 1

	baseQuery := `
        SELECT 
            reservations.id, 
            name, 
            reserver_phone_number, 
            company_name, 
            sub_total_snack, 
            sub_total_rooms, 
            total, 
            status_reservation, 
            created_at, 
            updated_at,
			user_id
        FROM reservations
        WHERE 1=1
    `

	// Apply username filter first (FIXED)
	if filter.UserID != 0 {
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", argCount)) // ✅ Fixed
		args = append(args, filter.UserID)
		argCount++
	}

	if filter.StartDate != "" {
		conditions = append(conditions, fmt.Sprintf("created_at >= $%d", argCount))
		args = append(args, filter.StartDate)
		argCount++
	}

	if filter.EndDate != "" {
		conditions = append(conditions, fmt.Sprintf("created_at <= $%d", argCount))
		args = append(args, filter.EndDate)
		argCount++
	}

	if filter.Type != "" {
		conditions = append(conditions, fmt.Sprintf("EXISTS (SELECT 1 FROM reservation_details LEFT JOIN rooms ON rooms.id = reservation_details.room_id WHERE reservation_id = reservations.id AND rooms.type = $%d)", argCount))
		args = append(args, filter.Type)
		argCount++
	}

	if filter.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status_reservation = $%d", argCount))
		args = append(args, filter.Status)
		argCount++
	}

	// Apply conditions to base query
	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Count total data (MOVED AFTER conditions are applied)
	countQuery := "SELECT COUNT(*) FROM (" + baseQuery + ") as count_query"
	var totalData int
	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalData)
	if err != nil {
		return nil, 0, err
	}

	// Remove the checkTotalReservationByID section - it's not needed

	// Rest of the code remains the same...

	// Add pagination
	baseQuery += " ORDER BY created_at DESC"
	if filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCount, argCount+1)
		args = append(args, filter.PageSize, offset)
	}

	rows, err := r.db.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var histories []model.ReservationHistory
	for rows.Next() {
		var h model.ReservationHistory
		err := rows.Scan(
			&h.ID,
			&h.Name,
			&h.PhoneNumber,
			&h.Company,
			&h.SubTotalSnack,
			&h.SubTotalRoom,
			&h.Total,
			&h.Status,
			&h.CreatedAt,
			&h.UpdatedAt,
			&h.UserID,
		)
		if err != nil {
			return nil, 0, err
		}
		histories = append(histories, h)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return histories, totalData, nil
}

func (r *reservationHistoryRepository) GetHistoryRooms(ctx context.Context, reservationIDs []int) (map[int][]model.ReservationHistoryRoom, error) {
	if len(reservationIDs) == 0 {
		return make(map[int][]model.ReservationHistoryRoom), nil
	}

	placeholders := make([]string, len(reservationIDs))
	args := make([]interface{}, len(reservationIDs))
	for i, id := range reservationIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT 
			reservation_details.reservation_id,
			reservation_details.id, 
			reservation_details.price, 
			reservation_details.name, 
			rooms.type, 
			total_room, 
			total_snack
		FROM reservation_details
		LEFT JOIN rooms on rooms.id = reservation_details.room_id
		WHERE reservation_id IN (%s)
		ORDER BY reservation_id, id
	`, strings.Join(placeholders, ", "))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roomsMap := make(map[int][]model.ReservationHistoryRoom)
	for rows.Next() {
		var reservationID int
		var room model.ReservationHistoryRoom
		err := rows.Scan(
			&reservationID,
			&room.ID,
			&room.Price,
			&room.Name,
			&room.Type,
			&room.TotalRoom,
			&room.TotalSnack,
		)
		if err != nil {
			return nil, err
		}
		roomsMap[reservationID] = append(roomsMap[reservationID], room)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roomsMap, nil
}
