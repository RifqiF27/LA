package collections

import "time"

type PaymentMethod struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    PhotoURL  string    `json:"photo_url"`
    IsActive  bool      `json:"is_active"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at"`
}

