package service

import (
	"cinema/model"
	"cinema/repository"
	"errors"
)

type CinemaService interface {
    GetAllCinemas() ([]model.Cinema, error)
    GetCinemaByID(id int) (*model.Cinema, error)
    GetAvailableSeats(cinemaID int, date, time string) ([]model.Seat, error)
}

type CinemaServiceImpl struct {
    Repo repository.CinemaRepository
}

func NewCinemaService(repo repository.CinemaRepository) CinemaService {
    return &CinemaServiceImpl{Repo: repo}
}

func (s *CinemaServiceImpl) GetAllCinemas() ([]model.Cinema, error) {
    return s.Repo.GetAll()
}

func (s *CinemaServiceImpl) GetCinemaByID(id int) (*model.Cinema, error) {
    cinema,err := s.Repo.GetByID(id)
	if err != nil {
		return nil, errors.New("cinema ID not found")
	}
	
	return  cinema, nil
}

func (s *CinemaServiceImpl) GetAvailableSeats(cinemaID int, date, time string) ([]model.Seat, error) {
    return s.Repo.GetSeats(cinemaID, date, time)
}
