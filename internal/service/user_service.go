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
	// Minimal 1 huruf kecil, 1 huruf besar, 1 angka, 1 simbol, panjang ≥ 8
	var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$`)
	return passwordRegex.MatchString(password)
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
