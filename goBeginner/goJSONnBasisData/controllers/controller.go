package controllers

import (
	"encoding/json"
	"fmt"
	"main/models"
	"main/utils"
	"os"
)

func ManageOrder() {
	menu := model.GetMenuData()
	var data model.Data

	err := json.Unmarshal([]byte(menu), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}
	for {
		fmt.Println("\n--- Restaurant Reservation ---")
		fmt.Println("\n1. Order")
		fmt.Println("2. Edit Order")
		fmt.Println("3. Order List")
		fmt.Println("4. Payment")
		fmt.Println("5. Status Order")
		fmt.Println("6. Exit")

		// var option int
		fmt.Print("Choose menu: ")
		option, err := utils.InputInt()
		if err != nil {

		}
		utils.ClearScreen()
		switch option {
		case 1:
			var categoryIdx, itemIdx int
			fmt.Println("\nCategory")
			for i, category := range data.Menu.Category {
				fmt.Printf("%d. %s\n", i+1, category.Name)
			}
			categoryIdx, err = utils.InputInt()
			if err != nil {

			}
			if categoryIdx < 1 || categoryIdx > len(data.Menu.Category) {
				fmt.Println("Category Invalid.")
				continue
			}
			fmt.Println("Choose Item:")
			for i, item := range data.Menu.Category[categoryIdx-1].Item {
				fmt.Printf("%d. %s - Rp %d\n", i+1, item.Name, item.Price)
			}
			itemIdx, err = utils.InputInt()
			if err != nil {

			}
			if itemIdx < 1 || itemIdx > len(data.Menu.Category[categoryIdx-1].Item) {
				fmt.Println("Item invalid.")
				continue
			}
			item := data.Menu.Category[categoryIdx-1].Item[itemIdx-1]
			data.Order = append(data.Order, model.Order{Name: item.Name, Price: item.Price})
			data.Payment.Amount += item.Price
			fmt.Printf("Order %s successfully added.\n", item.Name)
		case 2:
			// Edit Order
			if len(data.Order) == 0 {
				fmt.Println("No orders yet.")
				continue
			}
			fmt.Println("\nSelect the order you want to edit:")
			for i, p := range data.Order {
				fmt.Printf("%d. %s - Rp %d\n", i+1, p.Name, p.Price)
			}
			var orderIdx int
			orderIdx, err = utils.InputInt()
			if err != nil {

			}
			if orderIdx < 1 || orderIdx > len(data.Order) {
				fmt.Println("Order invalid.")
				continue
			}
			var categoryIdx, itemIdx int
			fmt.Println("\nSelect new category:")
			for i, category := range data.Menu.Category {
				fmt.Printf("%d. %s\n", i+1, category.Name)
			}
			categoryIdx, err = utils.InputInt()
			if err != nil {

			}
			if categoryIdx < 1 || categoryIdx > len(data.Menu.Category) {
				fmt.Println("Category invalid.")
				continue
			}
			fmt.Println("Select new Item:")
			for i, item := range data.Menu.Category[categoryIdx-1].Item {
				fmt.Printf("%d. %s - Rp %d\n", i+1, item.Name, item.Price)
			}
			itemIdx, err = utils.InputInt()
			if err != nil {

			}
			if itemIdx < 1 || itemIdx > len(data.Menu.Category[categoryIdx-1].Item) {
				fmt.Println("Item invalid.")
				continue
			}
			item := data.Menu.Category[categoryIdx-1].Item[itemIdx-1]
			data.Payment.Amount = data.Payment.Amount - data.Order[orderIdx-1].Price + item.Price
			data.Order[orderIdx-1] = model.Order{Name: item.Name, Price: item.Price}
			fmt.Printf("The order has been changed to %s.\n", item.Name)
		case 3:
			if len(data.Order) == 0 {
				fmt.Println("No orders yet.")
			} else {
				fmt.Println("Your order:")
				for _, p := range data.Order {
					fmt.Printf("- %s: Rp %d\n", p.Name, p.Price)
				}
				fmt.Printf("Amount: Rp %d\n", data.Payment.Amount)
			}
		case 4:
			fmt.Println("\nStatus Payment: \n1. Unpaid\n2. Paid")
			var statusPayment int
			statusPayment, err = utils.InputInt()
			if err != nil {

			}
			if statusPayment == 1 {
				data.Payment.Status = "Unpaid"
			} else if statusPayment == 2 {
				data.Payment.Status = "Paid"
			} else {
				fmt.Println("Invalid.")
			}
		case 5:
			fmt.Println("\nStatus Order: \n1. Process\n2. Delivered\n3. Done")
			var statusOrder int
			statusOrder, err = utils.InputInt()
			if err != nil {

			}
			if statusOrder == 1 {
				data.StatusOrder.Status = "Process"
			} else if statusOrder == 2 {
				data.StatusOrder.Status = "Delivered"
			} else if statusOrder == 3 {
				data.StatusOrder.Status = "Done"
			} else {
				fmt.Println("Invalid.")
			}
		case 6:
			fmt.Println("Exit.")
			return
		default:
			fmt.Println("Invalid.")
		}

		resultOrder, _ := json.MarshalIndent(data.Order, "", "  ")
		resultPayment, _ := json.MarshalIndent(data.Payment, "", "  ")
		resultStatus, _ := json.MarshalIndent(data.StatusOrder, "", "  ")
		fmt.Println("\nOrder Data in JSON format:")
		fmt.Printf("Order list: %s\nPayment: %s\nStatus Order: %s", string(resultOrder), string(resultPayment), string(resultStatus))
	}
}
