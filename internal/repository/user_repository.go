package repository

import (
	"E-Meeting/model"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, username, password FROM users WHERE username = $1`
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
