package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email, username, password, confirmPassword string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(email, username, password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("passwords do not match")
	}

	// Check if email already exists
	existing, err := s.repo.GetByEmail(email)
	if err == nil && existing != nil {
		return errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	return s.repo.Create(&user)
}
