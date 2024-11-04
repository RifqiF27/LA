package model

type Category struct {
	ID           uint16 `json:"category_id,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}
