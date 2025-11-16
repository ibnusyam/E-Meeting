package repository

import (
	"E-Meeting/model"
	"database/sql"
	"errors"
	"time"
)

type PasswordResetRepository struct {
	DB *sql.DB
}

func NewPasswordResetRepository(db *sql.DB) *PasswordResetRepository {
	return &PasswordResetRepository{DB: db}
}

func (r *PasswordResetRepository) CheckEmailExists(email string) (*model.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = $1 LIMIT 1"
	row := r.DB.QueryRow(query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *PasswordResetRepository) CreateToken(userID int, token string) error {
	query := `
		INSERT INTO password_resets (user_id, token, created_at, expired_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.DB.Exec(query, userID, token, time.Now(), time.Now().Add(15*time.Minute))
	return err
}

func (r *PasswordResetRepository) UpdatePassword(id int, hashedPassword string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, hashedPassword, id)
	return err
}
