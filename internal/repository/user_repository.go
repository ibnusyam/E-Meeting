package repository

import (
	"E-Meeting/model"

	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

type UserRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
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

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	query := `INSERT INTO users (email, username, password, role) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, user.Email, user.Username, user.Password, "customer")
	return err
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, username, password FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
