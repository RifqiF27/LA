package model

type Item struct {
	ID         uint16 `json:"item_id,omitempty"`              
	ItemCode   string `json:"item_code,omitempty"`
	ItemName   string `json:"item_name,omitempty"`
	Stock      int    `json:"stock,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
	LocationID int    `json:"location_id,omitempty"`
}