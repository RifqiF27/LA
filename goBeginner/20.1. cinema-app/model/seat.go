package model

type Seat struct {
	ID       int    `json:"id"`
	SeatID   string `json:"seatId"`
	CinemaID int    `json:"cinema_id"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Status   string `json:"status"`
}

type SeatResponse struct {
	SeatID   string `json:"seatId"`
	Status   string `json:"status"`
}
