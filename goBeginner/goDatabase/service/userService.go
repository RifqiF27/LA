package service

import (
	"database/sql"
	"errors"
	"fmt"
	"main/model"
	"main/repository"
)


func InputDataUser(db *sql.DB, username, password, role string) (int, string, error) {
	if username == "" {
		return 0,"", errors.New("username tidak boleh kosong")
	}
	if password == "" {
		return 0,"", errors.New("password tidak boleh kosong")
	}
	if role == "" {
		return 0,"", errors.New("role tidak boleh kosong")
	}

	// tx, err := db.Begin()
	// if err != nil {
	// 	return 0,"", fmt.Errorf("gagal memulai transaksi: %w", err)
	// }

	userRepo := repository.NewUserRepo(db)
	user := model.User{
		Username: username,
		Password: password,
		Role:     role,
	}

	fmt.Println(user, "<<<<")

	err := userRepo.Create(&user)
	if err != nil {
		// tx.Rollback() 
		return 0,"", fmt.Errorf("gagal membuat user: %w", err)
	}
	
	fmt.Println("berhasil input data user dengan id ", user.ID)
	return int(user.ID),user.Role, nil
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

func GetUserIsActive(db *sql.DB) error {
	userRepo := repository.NewUserRepo(db)
	var users []model.UserLog

	err := userRepo.GetUserActive(&users)
	if err != nil {
		return fmt.Errorf("gagal mengambil data: %w", err)
	}

	for _, user := range users {
		fmt.Printf("Active: %v, InActive: %v\n", user.Status, user.Status)
	}

	return nil
}
