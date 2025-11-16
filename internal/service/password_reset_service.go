package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/internal/utils"
	"errors"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type PasswordResetService struct {
	Repo *repository.PasswordResetRepository
}

func NewPasswordResetService(repo *repository.PasswordResetRepository) *PasswordResetService {
	return &PasswordResetService{Repo: repo}
}

func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 32)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}

func (s *PasswordResetService) GenerateResetToken(email string) (string, error) {
	user, err := s.Repo.CheckEmailExists(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("email not found")
	}

	token, err := utils.GenerateToken()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *PasswordResetService) UpdatePassword(id int, newPassword, confirmPassword string) error {
	if newPassword != confirmPassword {
		return errors.New("password confirmation is not match")
	}

	// Hash password dengan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Simpan ke DB lewat repository
	err = s.Repo.UpdatePassword(id, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
