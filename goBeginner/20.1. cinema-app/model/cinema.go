package model

// import "time"

type Cinema struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Seats    int    `json:"seats,omitempty"`
	// CreateAt time.Time
}
