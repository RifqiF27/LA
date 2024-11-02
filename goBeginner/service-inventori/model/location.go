package model

type Location struct {
	ID           uint16 `json:"location_id,omitempty"`
	LocationName string `json:"location_name,omitempty"`
}
