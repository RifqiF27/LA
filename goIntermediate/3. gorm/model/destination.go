package model

type Destination struct {
	OriginLat       string `gorm:"-" json:"origin_lat"`
	OriginLong      string `gorm:"-" json:"origin_long"`
	DestinationLat  string `gorm:"-" json:"destination_lat"`
	DestinationLong string `gorm:"-" json:"destination_long"`
}

type CostResponse struct {
	Shipping string  `json:"shipping"`
	Distance float64 `json:"distance"`
	Cost     float64 `json:"cost"`
}

type RequestDestination struct {
	ShippingID int `json:"shipping_id" binding:"required" form:"shipping_id"`
	Qty        int `json:"qty" binding:"required" form:"qty"`
	UserID     int `json:"user_id" binding:"required" form:"user_id"`
}
