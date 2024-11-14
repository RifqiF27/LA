package model

import "time"

type Destination struct {
	ID          int    `json:"id"`
	Location    string `json:"location"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}
type Gallery struct {
	ID          int    `json:"id"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type Event struct {
	ID            int       `json:"id"`
	DestinationID int       `json:"destination_id"`
	Name          string    `json:"name"`
	Schedule      time.Time `json:"schedule"`
	Price         float64   `json:"price"`
}

type Review struct {
	ID            int    `json:"id"`
	DestinationID int    `json:"destination_id"`
	TransactionID int    `json:"transaction_id"`
	Rating        int    `json:"rating"`
	Comment       string `json:"comment"`
}

type Transaction struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required,len=10,numeric"`
	Comment   string    `json:"comment"`
	EventID   int       `json:"event_id" validate:"required"`
	Status    string    `json:"status" validate:"required,oneof=ok cancel"`
	StatusTrx bool      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type DestinationEventRating struct {
	ID            int       `json:"id"`
	Location      string    `json:"location"`
	ImageURL      string    `json:"image_url"`
	Description   string    `json:"description"`
	Gallery       []Gallery `json:"gallery,omitempty"`
	EventName     string    `json:"event_name"`
	Schedule      time.Time `json:"schedule"`
	Price         float64   `json:"price"`
	AverageRating float64   `json:"average_rating"`
	People        int       `json:"people"`
}
