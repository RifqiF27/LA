package collections

import "time"

type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderItem struct {
    ID       int     `json:"id"`
    OrderID  int     `json:"order_id"`
    BookID   int     `json:"book_id"`
    Quantity int     `json:"quantity"`
    Price    float64 `json:"price"`
}