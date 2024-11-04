package model

type Todo struct {
	ID     uint16 `json:"todo_id,omitempty"`
	UserID int `json:"user_id,omitempty"`
	Thread string `json:"thread,omitempty"`
	Status string `json:"status,omitempty"`
}
