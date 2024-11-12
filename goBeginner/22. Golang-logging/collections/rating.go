package collections

import "time"

type Rating struct {
	ID         int       `json:"id"`
	OrderID    int       `json:"order_id"`
	CustomerID int       `json:"customer_id"`
	Rating     int       `json:"rating"`
	Review     string    `json:"review"`
	CreatedAt  time.Time `json:"created_at"`
}