package model

import "time"

type Transaction struct {
	TransactionID   uint16    `json:"transaction_id"`
	ItemID          int       `json:"item_id"`
	TransactionType string    `json:"transaction_type"`
	Quantity        int       `json:"quantity"`
	Timestamp       time.Time `json:"timestamp"`
	Notes           string    `json:"notes,omitempty"`
	UserID          *int      `json:"user_id,omitempty"`
}

type TransactionHistory struct {
	TransactionID   uint16    `json:"transaction_id"`
	ItemName        string    `json:"item_name"`
	TransactionType string    `json:"transaction_type"`
	Quantity        int       `json:"quantity"`
	Timestamp       time.Time `json:"timestamp"`
	Notes           string    `json:"notes,omitempty"`
	UserID          *int      `json:"user_id,omitempty"`
}
