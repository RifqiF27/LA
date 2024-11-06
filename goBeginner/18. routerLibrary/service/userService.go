package service

import (
	"database/sql"
	"errors"
	"fmt"
	"todo/model"
	"todo/repository"
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

func (us *UserService) RegisterService(username, password, role string) error {
	if username == "" {
		return errors.New("username tidak boleh kosong")
	}
	if password == "" {
		return errors.New("password tidak boleh kosong")
	}
	if role == "" {
		return errors.New("role tidak boleh kosong")
	}

	// if role == "admin" && name == "" {
	// 	return errors.New("nama admin tidak boleh kosong")
	// }

	exists, err := us.RepoUser.UsernameExists(username)
	if err != nil {
		return fmt.Errorf("gagal memeriksa keberadaan username: %w", err)
	}
	if exists {
		return errors.New("username sudah ada, silakan pilih yang lain")
	}

	user := model.User{
		Username: username,
		Password: password,
		Role:     role,
	}

	err = us.RepoUser.Create(&user)
	if err != nil {
		// tx.Rollback()
		return fmt.Errorf("gagal membuat user: %w", err)
	}
	// if role == "admin" {
	// 	admin := model.Admin{
	// 		Name:   name,
	// 		UserID: int(user.ID),
	// 	}

	// 	err = us.RepoUser.CreateAdmin(&admin)
	// 	if err != nil {
	// 		return fmt.Errorf("gagal membuat admin: %w", err)
	// 	}
	// }

	fmt.Println("berhasil input data user dengan id ", user.ID)
	return nil
}
