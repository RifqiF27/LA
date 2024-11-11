package service

import (
	"cinema/model"
	"cinema/repository"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type AuthService interface {
	Register(user model.User) error
	Login(user model.User) (*model.User, error)
	Logout(token string) error
	VerifyToken(token string) (int, error)
}

type AuthServiceImpl struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *AuthServiceImpl) Register(user model.User) error {
	_, err := s.repo.GetUserLogin(user)
	if err == nil {
		return errors.New("username already exists")
	}

	users := &model.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}

	err = s.repo.CreateUser(users)
	if err != nil {
		return errors.New("email already exists")
	}
	return nil
}

func (s *AuthServiceImpl) Login(user model.User) (*model.User, error) {

	users, err := s.repo.GetUserLogin(user)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}
	if users.ID == 0 {
		return nil, errors.New("user ID is invalid or missing")
	}

	token, err := generateToken()
	if err != nil {
		return nil, err
	}
	session := &model.Session{
		UserID:    users.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := s.repo.CreateSession(session); err != nil {
		return nil, err
	}

	users.Token = token
	return users, nil
}

func (s *AuthServiceImpl) Logout(token string) error {

	return s.repo.DeleteSession(token)
}

func (s *AuthServiceImpl) VerifyToken(token string) (int, error) {
	session, err := s.repo.GetSessionByToken(token)
	if err != nil || session.ExpiresAt.Before(time.Now()) {
		fmt.Println("Failed to verify token:", token)
		return 0, errors.New("invalid or expired token")
	}
	fmt.Println("Token verified:", token)
	return session.UserID, nil
}
