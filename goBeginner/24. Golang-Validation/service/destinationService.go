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

func (s *DestinationService) GetById(id int) (*model.DestinationEventRating, error) {
	return s.Repo.GetById(id)
}

func (s *DestinationService) GetTourPlansByEventID(id int) (*model.TourPlan, error) {
	return s.Repo.GetTourPlansByEventID(id)
}
func (s *DestinationService) GetLocationsByDestinationID(id int) (*model.Location, error) {
	return s.Repo.GetLocationsByDestinationID(id)
}


