package models

type User struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"Phone_number"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type Account struct {
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	ID          int     `json:"id"`
	FromUserID  string  `json:"from_user_id"`
	ToUserID    string  `json:"to_user_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
}

type UserServices struct {
	UsersFile    string
	SessionsFile string
	Users        []User
	Session      *User
}
type AccountService struct {
	AccountsFile     string
	TransactionsFile string
	Accounts         []Account
	Transactions     []Transaction
}
