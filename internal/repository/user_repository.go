package repository

import (
	"E-Meeting/model"

	"database/sql"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
}

type userRepository struct {
	DB *sql.DB
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, username, password, role FROM users WHERE username = $1`
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(user *model.User) error {
	query := `INSERT INTO users (email, username, password, role) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, user.Email, user.Username, user.Password, "customer")
	return err
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, username, password FROM users WHERE email = $1`
	row := r.DB.QueryRow(query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
