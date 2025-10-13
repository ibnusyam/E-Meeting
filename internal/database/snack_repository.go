package database

import (
	"E-Meeting/internal/model"
	"context"
	"database/sql"
	"time"
)

type SnackRepository struct {
	DB *sql.DB
}

func NewSnackRepository(db *sql.DB) *SnackRepository {
	return &SnackRepository{DB: db}
}

func (repo *SnackRepository) GetAllSnacks() ([]model.Snack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, price, category FROM snacks`
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snacks []model.Snack

	for rows.Next() {
		var snack model.Snack

		err := rows.Scan(&snack.ID, &snack.Name, &snack.Price, &snack.Category)
		if err != nil {
			return nil, err
		}
		snacks = append(snacks, snack)
	}
	return snacks, nil
}
