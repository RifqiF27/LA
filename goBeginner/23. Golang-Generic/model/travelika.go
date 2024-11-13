package model

import "time"

type Destination struct {
	ID          int    `json:"id"`
	Location    string `json:"location"`
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
	ID      int  `json:"id"`
	EventID int  `json:"event_id"`
	Status  bool `json:"status"`
}

type DestinationEventRating struct {
	ID            int       `json:"id"`
	Location      string    `json:"location"`
	ImageURL      string    `json:"image_url"`
	Destination   string    `json:"destination"`
	EventName     string    `json:"event_name"`
	Schedule      time.Time `json:"schedule"`
	Price         float64   `json:"price"`
	AverageRating float64   `json:"average_rating"`
}
