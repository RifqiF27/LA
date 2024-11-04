package main

import (
	"fmt"
	"main/services"
	"main/utils"
)

func main() {
	for {
		var username, password string

		fmt.Println("Welcome to the Hotel Reservation System")
		fmt.Print("Enter Username: ")
		fmt.Scanln(&username)
		fmt.Print("Enter Password: ")
		fmt.Scanln(&password)

		token, ctx, cancel, err := services.Login(username, password)
		if err != nil {
			fmt.Println("Login failed:", err)
			continue 
		}
		defer cancel()

		user, err := services.GetUserByUsername(username)
		if err != nil {
			fmt.Println("User not found")
			continue 
		}

		go func() {
			select {
			case <-ctx.Done():
				fmt.Println("\nSession expired. Please log in again.")
				
				return
			}
		}()

		role := user.Role

		
		if !utils.ValidateSession(token) {
			fmt.Println("Invalid session token. Please log in again.")
			continue 
		}

		if role == "admin" {
			handleAdminMenu()
		} else if role == "customer" {
			handleCustomerMenu(username)
		} else {
			fmt.Println("Invalid user role")
		}
	}
}

func handleAdminMenu() {
	for {
		fmt.Println("\nAdmin Menu:")
		fmt.Println("1. Manage Customers")
		fmt.Println("2. Manage Facilities")
		fmt.Println("3. Manage Reservations")
		fmt.Println("4. View Available Rooms")
		fmt.Println("5. Logout")
		fmt.Print("Select an option: ")
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			services.ManageCustomers()
		case 2:
			services.ManageFacilities()
		case 3:
			services.ManageReservations()
		case 4:
			services.ViewAvailableRooms()
		case 5:
			fmt.Println("Logging out...")
			return 
		default:
			fmt.Println("Invalid option")
		}
	}
}

func handleCustomerMenu(username string) {
	for {
		fmt.Println("\nCustomer Menu:")
		fmt.Println("1. Make a Reservation")
		fmt.Println("2. Cancel a Reservation (Pending Approval)")
		fmt.Println("3. View My Reservations")
		fmt.Println("4. Logout")
		fmt.Print("Select an option: ")
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			services.MakeReservation(username)
		case 2:
			services.CancelReservation()
		case 3:
			services.ViewMyReservations(username)
		case 4:
			fmt.Println("Logging out...")
			return 
		default:
			fmt.Println("Invalid option")
		}
	}
}
