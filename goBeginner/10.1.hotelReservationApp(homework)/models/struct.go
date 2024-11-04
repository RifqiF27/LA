package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Room struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	IsAvailable bool   `json:"is_available"`
}

type Reservation struct {
	ID                int    `json:"id"`
	CustomerName      string `json:"customer_name"`
	CustomerPhone     string `json:"customer_phone"`
	RoomID            int    `json:"room_id"`
	ReservationDate   string `json:"reservation_date"`
	TotalPayment      int    `json:"total_payment"`
	PaymentStatus     string `json:"payment_status"`
	ReservationStatus string `json:"reservation_status"`
	CreatedAt         string `json:"created_at"`
}

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Facility struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	IsActive bool   `json:"is_active"`
}
