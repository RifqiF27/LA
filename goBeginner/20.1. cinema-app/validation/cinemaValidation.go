package validation

import (
	"cinema/model"
	"errors"
)

func ValidateCinema(cinema *model.Cinema) error {
	if cinema.Name == "" {
		return errors.New("cinema is required")
	}
	if cinema.Location == "" {
		return errors.New("location is required")
	}

	return nil
}