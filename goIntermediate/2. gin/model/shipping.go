package model

type Shipping struct {
	Base
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
