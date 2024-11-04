package service

import (
	"errors"
	"service-inventory/model"
	"service-inventory/repository"
)

type LocationService struct {
	RepoLocation repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) *LocationService {
	return &LocationService{RepoLocation: repo}
}

func (cs *LocationService) AddLocationService( name string) error {
	if name == "" {
		return errors.New("Location name cannot be empty")
	}
	

	location := model.Location{
		LocationName: name,
	}
	_, err := cs.RepoLocation.Create(&location)
	if err != nil {
		return errors.New("failed to create location: " + err.Error())
	}

	return nil
}