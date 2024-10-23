package models

type Product struct {
	ID    int
	Name  string
	Price float64
}

type CartItem struct {
	Product  Product
	Quantity int
}

type Item struct {
	Cart    []CartItem
	Session bool
}

var Products = []Product{
	{ID: 1, Name: "Laptop", Price: 15000},
	{ID: 2, Name: "Mouse", Price: 200},
	{ID: 3, Name: "Keyboard", Price: 500},
	{ID: 4, Name: "SSD", Price: 3000},
}

var Items = Item{
	Cart:    []CartItem{},
	Session: false,
}
