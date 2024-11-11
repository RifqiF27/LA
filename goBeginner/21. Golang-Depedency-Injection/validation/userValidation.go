package validation

import (
	"book-store/collections"
	"errors"
)

func ValidateUser(user *collections.User) error {
	if user.Username == "" {
		return errors.New("user name is required")
	}
	if user.Password <= "" {
		return errors.New("user price must be greater than zero")
	}
	return nil
}