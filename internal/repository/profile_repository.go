package repository

import (
	"E-Meeting/model"
	"context"
	"database/sql"
	"errors"
	"time"
)

type ProfileRepository struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (repo *ProfileRepository) GetUserProfileByID(id string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := "SELECT id, role, username, email, password, phone_number, language, status_user, profile_picture, created_at, updated_at FROM USERS where id=" + id

	rows := repo.DB.QueryRowContext(ctx, query)

	var user model.User

	err := rows.Scan(&user.ID, &user.Role, &user.Username, &user.Email, &user.Password, &user.PhoneNumber, &user.Language, &user.Status, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}
