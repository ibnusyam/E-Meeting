package service

import (
	"E-Meeting/internal/repository"
	"errors"
	"log"
	"time"
)

type DashboardService interface {
	GetDashboard(startDateStr, endDateStr string) (repository.DashboardData, error)
}

type dashboardService struct {
	repo repository.DashboardRepository
}

func NewDashboardService(repo repository.DashboardRepository) DashboardService {
	return &dashboardService{repo: repo}
}

func (s *dashboardService) GetDashboard(startDateStr, endDateStr string) (repository.DashboardData, error) {
	layout := "2006-01-02"

	log.Printf("Start Date: %v", startDateStr)
	// startDate := startDateStr
	// endDate := endDateStr
	startDate, err := time.Parse(layout, startDateStr)
	log.Printf("Start Date: %v", startDate)
	if err != nil {
		return repository.DashboardData{}, errors.New("invalid start date format (use YYYY-MM-DD)")
	}

	log.Printf("End Date: %v", endDateStr)
	endDate, err := time.Parse(layout, endDateStr)
	log.Printf("End Date: %v", endDate)
	if err != nil {
		return repository.DashboardData{}, errors.New("invalid end date format (use YYYY-MM-DD)")
	}

	if startDate.After(endDate) {
		return repository.DashboardData{}, errors.New("start date must be smaller than end date")
	}

	data, err := s.repo.GetDashboardData(startDate, endDate)
	if err != nil {
		return repository.DashboardData{}, err
	}

	return data, nil
}
