package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
)

type SnackService struct {
	Repo *repository.SnackRepository
}

func NewSnackService(repo *repository.SnackRepository) *SnackService {
	return &SnackService{Repo: repo}
}

func (s *SnackService) GetAllSnacks() ([]model.Snack, error) {
	snacks, err := s.Repo.GetAllSnack()
	if err != nil {
		return nil, err
	}

	if snacks == nil {
		return []model.Snack{}, nil
	}

	return snacks, err
}
