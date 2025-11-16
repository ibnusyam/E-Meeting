package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/internal/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Login(username, password string) (string, string, error) {
	user, err := s.Repo.GetByUsername(username)
	if err != nil {
		return "", "", errors.New("login failed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("login failed")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
