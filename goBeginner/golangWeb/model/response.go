package model

type Response struct {
	StatusCode int
	Message    string
	Page       int         `json:"page,omtodopty"`
	Limit      int         `json:"limit,omtodopty"`
	TotalTodos int         `json:"total_todos,omtodopty"`
	TotalPages int         `json:"total_pages,omtodopty"`
	Data       interface{} `json:"data,omtodopty"`
}
