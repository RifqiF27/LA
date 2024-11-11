package collections

import "encoding/json"

type Customer struct {
	ID      int             `json:"id"`
	UserID  int             `json:"user_id"`
	Name    string          `json:"name"`
	Phone   string          `json:"phone"`
	Address json.RawMessage `json:"address"`
}