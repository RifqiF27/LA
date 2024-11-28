package model

type Shipping struct {
	Base
	Name  string  `gorm:"type:varchar(100);not null" json:"name"`
	Price float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}

type OrderShipping struct {
	Base
	OrderID            string   `gorm:"type:varchar(100);unique not nul" json:"order_id"`
	ShippingID         uint     `gorm:"type:int;not null" json:"shipping_id"`
	OriginLatLong      string   `gorm:"type:varchar(100);not null" json:"origin_latlong"`
	DestinationLatLong string   `gorm:"type:varchar(100);not null" json:"destination_latlong"`
	TotalPayment       float64  `gorm:"type:decimal(10,2);not null" json:"total_payment_shipping"`
	Shipping           Shipping `gorm:"foreignKey:ShippingID;references:ID"`
}

type HistoryDelivery struct {
	Base
	OrderShippingID uint          `gorm:"type:int;not null" json:"order_shipping_id"`
	Status          string        `gorm:"type:varchar(20);not null;default:'pending';check:status IN ('pending', 'shipped', 'completed', 'canceled')" json:"status"`
	Location        string        `gorm:"type:varchar(100);not null" json:"location"`
	OrderShipping   OrderShipping `gorm:"foreignKey:OrderShippingID;references:ID"`
}
