package userservice

import (
	"dashboard-ecommerce-team2/database"
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"
	"errors"
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(userInput models.RegisterRequest) error
	Login(input models.LoginRequest) (*models.User, error)
	CheckUserEmail(email string) (*models.User, error)
	ResetUserPassword(email, otp, password string) error
	RequestResetOTP(email string) error

}

type userService struct {
	Repo repository.Repository
	Log  *zap.Logger
	Cacher database.Cacher
}

func NewUserService(repo repository.Repository, log *zap.Logger, cacher database.Cacher) UserService {
	return &userService{Repo: repo, Log: log, Cacher: cacher}
}

// CheckUserEmail implements UserService.
func (u *userService) CheckUserEmail(email string) (*models.User, error) {
	return u.Repo.User.GetByEmail(email)
}

// CreateUser implements UserService.
func (u *userService) CreateUser(userInput models.RegisterRequest) error {
	newUserInput := models.User{
		Email:    userInput.Email,
		Password: helper.HashPassword(userInput.Password),
		Role:     "staff",
		Name:     userInput.Name,
	}
	return u.Repo.User.Create(newUserInput)
}

// Login implements UserService.
func (u *userService) Login(input models.LoginRequest) (*models.User, error) {
	user, err := u.Repo.User.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	// Check if the user's password matches the input password
	if !helper.CheckPassword(input.Password, user.Password) {
		return nil, errors.New("invalid user password")
	}
	return user, nil
}

// ResetUserPassword implements UserService.
func (u *userService) ResetUserPassword(email, otp, password string) error {
	
	cacheKey := fmt.Sprintf("otp:%s", email)
	storedOtp, err := u.Cacher.Get(cacheKey)
	if err != nil {
		log.Println(err, "<<<<<<<")
		return fmt.Errorf("OTP tidak ditemukan atau sudah kedaluwarsa")
	}

	// Validate OTP
	if storedOtp != otp {
		return fmt.Errorf("OTP tidak valid")
	}

	// Delete OTP from Redis after validation
	// if err := u.Cacher.Delete(cacheKey); err != nil {
	// 	fmt.Printf("Peringatan: Gagal menghapus OTP untuk %s dari Redis: %v\n", email, err)
	// }

	// Hash new password and update it in the database
	hashedPassword := helper.HashPassword(password)
	if err := u.Repo.User.UpdatePassword(email, hashedPassword); err != nil {
		return fmt.Errorf("gagal memperbarui password: %v", err)
	}

	return nil
}

func (svc *userService) RequestResetOTP(email string) error {
	// Validate if the email exists in the database
	user, err := svc.Repo.User.GetByEmail(email)
	if err != nil || user == nil {
		return fmt.Errorf("email tidak terdaftar")
	}

	// Generate OTP
	otp := helper.GenerateOTP(6)
	cacheKey := fmt.Sprintf("otp:%s", email)
	expiration := 5 * time.Minute

	// Save OTP to Redis
	if err := svc.Cacher.Set(cacheKey, otp, expiration); err != nil {
		return fmt.Errorf("gagal menyimpan OTP ke Redis: %v", err)
	}

	// Simulate sending email
	go func() {
		err := helper.SendDummyEmail(email, "Reset Password OTP", fmt.Sprintf("Your OTP is: %s", otp))
		if err != nil {
			fmt.Printf("Gagal mengirim email ke %s: %v\n", email, err)
		}
	}()

	return nil
}

