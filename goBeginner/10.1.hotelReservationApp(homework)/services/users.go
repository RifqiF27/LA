package services

import (
	"context"
	"errors"
	"main/models"
	"main/utils"

	"time"
)

var users []models.User

func init() {
	err := utils.ReadJSON("data/users.json", &users)
	if err != nil {
		panic("Could not load users data.")
	}
}

func Login(username, password string) (string, context.Context, context.CancelFunc, error) {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			token, ctx, cancel := utils.CreateSession(username, 10*time.Second)
			return token, ctx, cancel, nil
		}
	}
	return "", nil, nil, errors.New("invalid credentials")
}

func GetUserByUsername(username string) (models.User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}
