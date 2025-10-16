package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
)

type ProfileService struct {
	Repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{Repo: repo}
}

func (s *ProfileService) GetUserProfileByID(id string) (model.User, error) {
	profile, err := s.Repo.GetUserProfileByID(id)
	if err != nil {
		return profile, err
	}

	return profile, nil
}
