package model

import "time"

type User struct {
	ID       int	`json:"-"`
	Username string
	Password string `json:"password,omitempty"`
	Email    string
	Token    string
}

type Session struct {
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expired_at"`
}
