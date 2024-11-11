package repository

import (
	"cinema/model"
	"database/sql"
	"errors"
	"fmt"
)

type AuthRepository interface {
	CreateUser(user *model.User) error
	GetUserLogin(user model.User) (*model.User, error)
	CreateSession(session *model.Session) error
	GetSessionByToken(token string) (*model.Session, error)
	DeleteSession(token string) error
	tokenExists(token string) (bool, error)
}

type AuthRepositoryDb struct {
	DB *sql.DB
}

func NewAuthRepositoryDb(db *sql.DB) AuthRepository {
	return &AuthRepositoryDb{DB: db}
}

func (repo *AuthRepositoryDb) CreateUser(user *model.User) error {
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"
	return repo.DB.QueryRow(query, user.Username, user.Password, user.Email).Scan(&user.ID)
}

func (repo *AuthRepositoryDb) GetUserLogin(user model.User) (*model.User, error) {
	query := `SELECT id, username, email FROM users WHERE username=$1 AND password=$2`
	var users model.User
	err := repo.DB.QueryRow(query, user.Username, user.Password).Scan(&users.ID, &users.Username, &users.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return &users, err
}

func (repo *AuthRepositoryDb) CreateSession(session *model.Session) error {
	query := "INSERT INTO sessions (user_id, token, expires_at) VALUES ($1, $2, $3)"
	_, err := repo.DB.Exec(query, session.UserID, session.Token, session.ExpiresAt)
	return err
}

func (repo *AuthRepositoryDb) GetSessionByToken(token string) (*model.Session, error) {
	var session model.Session
	query := "SELECT user_id, token, expires_at FROM sessions WHERE token=$1"
	err := repo.DB.QueryRow(query, token).Scan(&session.UserID, &session.Token, &session.ExpiresAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("session not found or expired")
	}
	return &session, err
}

func (repo *AuthRepositoryDb) DeleteSession(token string) error {
    query := "DELETE FROM sessions WHERE token=$1"
    res, err := repo.DB.Exec(query, token)
    if err != nil {
        fmt.Println("Error executing delete:", err)
        return err
    }

    rowsAffected, _ := res.RowsAffected()
    fmt.Println("Rows affected by delete:", rowsAffected) // Debugging log
    if rowsAffected == 0 {
        fmt.Println("No session found with this token.") // If no rows were affected, the token might not exist.
    }
    
    return nil
}

func (repo *AuthRepositoryDb) tokenExists(token string) (bool, error) {
    query := `SELECT COUNT(*) FROM sessions WHERE token = $1`
    var count int
    err := repo.DB.QueryRow(query, token).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}
