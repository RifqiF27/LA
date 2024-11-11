package service

import (
	"book-store/collections"
	"book-store/repository"
)

type UserService interface {
	LoginService(user collections.User) (*collections.User, error)
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (us *userService) LoginService(user collections.User) (*collections.User, error) {
	return us.repo.GetUserLogin(user)
}
