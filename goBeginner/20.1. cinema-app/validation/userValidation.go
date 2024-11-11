package validation

import (
	"cinema/model"
	"errors"
)

func ValidateUser(user *model.User, isLogin bool) error {
    if user.Username == "" {
        return errors.New("username is required")
    }
    if user.Password == "" {
        return errors.New("password is required")
    }
    if !isLogin && user.Email == "" {
        return errors.New("email is required")
    } 
	return nil
}