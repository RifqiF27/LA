package service

import (
	"travelika/model"
	"travelika/repository"
)

type DestinationService struct {
	Repo repository.DestinationRepository
}

func NewDestinationService(repo repository.DestinationRepository) *DestinationService {
	return &DestinationService{Repo: repo}
}

func (s *DestinationService) GetAllEvents(eventName, location string, date string, orderBy string, orderAsc bool, limit, page int) ([]model.DestinationEventRating, int, int, error) {
	return s.Repo.GetAllEvents(eventName, location, date, orderBy, orderAsc, limit, page)
}
