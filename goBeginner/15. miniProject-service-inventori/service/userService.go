package service

import (
	"database/sql"
	"errors"
	"fmt"
	"service-inventory/model"
	"service-inventory/repository"
)

type UserService struct {
	RepoUser repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{RepoUser: repo}
}


func GetAllUsers(db *sql.DB) error {
	userRepo := repository.NewUserRepo(db)
	var users []model.User

	err := userRepo.GetAll(&users)
	if err != nil {
		return fmt.Errorf("gagal mengambil data user: %w", err)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Role: %s\n", user.ID, user.Username, user.Role)
	}

	return nil
}

func (us *UserService) LoginService(user model.User) (*model.User, error) {

	if user.Username == "" {
		return nil, errors.New("username cannot empty")
	}
	if user.Password == "" {
		return nil, errors.New("password cannot empty")
	}

	users, err := us.RepoUser.GetUserLogin(user)

	if err != nil {
		return nil, errors.New("account not found")
	}

	return users, nil
}
