package model

type Menu struct {
	Category []Category `json:"category"`
}

type Category struct {
	Name string `json:"name"`
	Item []Item `json:"item"`
}

type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Order struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Payment struct {
	Status string `json:"status"`
	Amount int    `json:"amount"`
}

type StatusOrder struct {
	Status string `json:"status"`
}

type Data struct {
	Menu        Menu        `json:"menu"`
	Payment     Payment     `json:"payment"`
	Order       []Order     `json:"order"`
	StatusOrder StatusOrder `json:"status_order"`
}

func GetMenuData() string {
	return `{
		"menu": {
		  "category": [
			{
			  "name": "Appetizer",
			  "item": [
				{
				  "name": "Cream Soup",
				  "price": 15000
				},
				{
				  "name": "Salad",
				  "price": 12000
				}
			  ]
			},
			{
			  "name": "Main Course",
			  "item": [
				{
				  "name": "Cordon Blue",
				  "price": 35000
				},
				{
				  "name": "Beef Black Paper",
				  "price": 72000
				}
			  ]
			},
			{
			  "name": "Drink",
			  "item": [
				{
				  "name": "Ice Tea",
				  "price": 8000
				},
				{
				  "name": "Orange Juice",
				  "price": 15000
				}
			  ]
			},
			{
			  "name": "DEssert",
			  "item": [
				{
				  "name": "Pudding",
				  "price": 12000
				},
				{
				  "name": "Ice Cream",
				  "price": 10000
				}
			  ]
			}
		  ]
		},
		"payment": {
		  "status": "Unpaid",
		  "amount": 0
		},
		"status_order": {
		  "status": "Process"
		}
	  }`

}
