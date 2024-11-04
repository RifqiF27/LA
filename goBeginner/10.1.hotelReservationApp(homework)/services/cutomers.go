package services

import (
	"errors"
	"fmt"
	"main/models"
	"main/utils"
)

var customers []models.Customer

func init() {
	err := utils.ReadJSON("data/customers.json", &customers)
	if err != nil {
		panic("Could not load customers data.")
	}
}

func ManageCustomers() {
	var action int
	fmt.Println("Customer Management:")
	fmt.Println("1. Add Customer")
	fmt.Println("2. View Customers")
	fmt.Println("3. Edit Customer")
	fmt.Println("4. Delete Customer")
	fmt.Print("Select an action: ")
	fmt.Scanln(&action)

	switch action {
	case 1:
		AddCustomerMenu()
	case 2:
		ViewCustomers()
	case 3:
		UpdateCustomerMenu()
	case 4:
		DeleteCustomerMenu()
	default:
		fmt.Println("Invalid action")
	}
}

func AddCustomerMenu() {
	var customer models.Customer
	fmt.Print("Enter Customer Name: ")
	fmt.Scanln(&customer.Name)
	fmt.Print("Enter Customer Phone: ")
	fmt.Scanln(&customer.Phone)
	fmt.Print("Enter Customer Email: ")
	fmt.Scanln(&customer.Email)

	if err := AddCustomer(customer); err != nil {
		fmt.Println("Error adding customer:", err)
	} else {
		fmt.Println("Customer added successfully!")
	}
}

func AddCustomer(customer models.Customer) error {
	if customer.Name == "" || customer.Phone == "" {
		return errors.New("customer name and phone are required")
	}
	customer.ID = len(customers) + 1
	customers = append(customers, customer)
	return utils.WriteJSON("data/customers.json", customers)
}

func ViewCustomers() {
	fmt.Println("Customers:")
	for _, c := range customers {
		fmt.Printf("ID: %d, Name: %s, Phone: %s, Email: %s\n", c.ID, c.Name, c.Phone, c.Email)
	}
}

func DeleteCustomerMenu() {
	var id int
	fmt.Print("Enter Customer ID to delete: ")
	fmt.Scanln(&id)

	if err := DeleteCustomer(id); err != nil {
		fmt.Println("Error deleting customer:", err)
	} else {
		fmt.Println("Customer deleted successfully!")
	}
}

func DeleteCustomer(id int) error {
	for i, customer := range customers {
		if customer.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			return utils.WriteJSON("data/customers.json", customers)
		}
	}
	return errors.New("customer not found")
}

func UpdateCustomerMenu() {
	var id int
	fmt.Print("Enter Customer ID to update: ")
	fmt.Scanln(&id)

	for i, customer := range customers {
		if customer.ID == id {
			var updatedCustomer models.Customer
			fmt.Print("Enter new Customer Name (leave blank to keep current): ")
			fmt.Scanln(&updatedCustomer.Name)
			fmt.Print("Enter new Customer Phone (leave blank to keep current): ")
			fmt.Scanln(&updatedCustomer.Phone)
			fmt.Print("Enter new Customer Email (leave blank to keep current): ")
			fmt.Scanln(&updatedCustomer.Email)

			// Update hanya field yang diisi
			if updatedCustomer.Name != "" {
				customer.Name = updatedCustomer.Name
			}
			if updatedCustomer.Phone != "" {
				customer.Phone = updatedCustomer.Phone
			}
			if updatedCustomer.Email != "" {
				customer.Email = updatedCustomer.Email
			}

			customers[i] = customer
			if err := utils.WriteJSON("data/customers.json", customers); err != nil {
				fmt.Println("Error updating customer:", err)
			} else {
				fmt.Println("Customer updated successfully!")
			}
			return
		}
	}
	fmt.Println("Error updating customer: customer not found")
}
