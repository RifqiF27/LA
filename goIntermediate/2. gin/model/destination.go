package model

type Destination struct {
	Origin_lat       string
	Origin_long      string
	Destination_lat  string
	Destination_long string
}


type CostResponse struct {
	Shipping string  `json:"shipping"`
	Distance   float64 `json:"distance"`
	Cost       float64 `json:"cost"`
}

type RequestDestination struct {
	ShippingID int `json:"shipping_id" binding:"required" form:"shipping_id"`
	Qty        int `json:"qty" binding:"required" form:"qty"`
	UserID     int `json:"user_id" binding:"required" form:"user_id"`
	// OriginLongLat      string `json:"origin_long_lat" binding:"required" form:"origin_long_lat"`
	// DestinationLongLat string `json:"destination_long_lat" binding:"required" form:"destination_long_lat"`
}
