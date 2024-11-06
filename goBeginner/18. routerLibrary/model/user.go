package model

type User struct {
	ID       uint16
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
