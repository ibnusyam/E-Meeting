package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"errors"
	"regexp"

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

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	return hasLower && hasUpper && hasNumber && hasSpecial
}

func (s *userService) Register(email, username, password, confirmPassword string) error {

	if !isValidPassword(password) {
		return errors.New("password must contain at least one uppercase, one lowercase, one number, one special character, and be at least 8 characters long")
	}

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
