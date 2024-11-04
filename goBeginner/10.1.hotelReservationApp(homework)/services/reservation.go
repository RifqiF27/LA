package services

import (
	"fmt"
	"main/models"
	"main/utils"
	"time"
)

var reservations []models.Reservation

func init() {
	err := utils.ReadJSON("data/reservations.json", &reservations)
	if err != nil {
		panic("Could not load reservations data.")
	}
}

// Function for admin to manage reservations
func ManageReservations() {
	var action int
	fmt.Println("\nReservation Management:")
	fmt.Println("1. View All Reservations")
	fmt.Println("2. Confirm a Reservation")
	fmt.Println("3. Cancel a Reservation")
	fmt.Print("Select an action: ")
	fmt.Scanln(&action)

	switch action {
	case 1:
		ViewAllReservations()
	case 2:
		ConfirmReservation()
	case 3:
		CancelReservationByAdmin()
	default:
		fmt.Println("Invalid action")
	}
}

func ViewAllReservations() {
	fmt.Println("\nAll Reservations:")
	for _, reservation := range reservations {
		fmt.Printf("ID: %d, Customer: %s, Room ID: %d, Date: %s, Total Payment: %d, Payment Status: %s, Reservation Status: %s\n",
			reservation.ID, reservation.CustomerName, reservation.RoomID, reservation.ReservationDate, reservation.TotalPayment, reservation.PaymentStatus, reservation.ReservationStatus)
	}
}

func ConfirmReservation() {
	var id int
	fmt.Print("Enter Reservation ID to confirm: ")
	fmt.Scanln(&id)

	for i, reservation := range reservations {
		if reservation.ID == id {
			if reservation.ReservationStatus == "Pending Approval" {
				reservations[i].ReservationStatus = "Confirmed"
				reservations[i].PaymentStatus = "Paid"
				utils.WriteJSON("data/reservations.json", reservations)
				fmt.Println("Reservation confirmed successfully!")
				return
			} else {
				fmt.Println("Reservation is already confirmed or canceled.")
				return
			}
		}
	}
	fmt.Println("Error: Reservation not found.")
}

func CancelReservationByAdmin() {
	var id int
	fmt.Print("Enter Reservation ID to cancel: ")
	fmt.Scanln(&id)

	for i, reservation := range reservations {
		if reservation.ID == id {
			reservations[i].ReservationStatus = "Canceled"
			reservations[i].PaymentStatus = "Refunded"
			utils.WriteJSON("data/reservations.json", reservations)
			fmt.Println("Reservation canceled successfully!")
			return
		}
	}
	fmt.Println("Error: Reservation not found.")
}

func MakeReservation(customerName string) {
	var reservation models.Reservation
	reservation.CustomerName = customerName
	reservation.CreatedAt = time.Now().Format(time.RFC3339)

	fmt.Print("Enter Room ID: ")
	fmt.Scanln(&reservation.RoomID)
	fmt.Print("Enter Total Payment: ")
	fmt.Scanln(&reservation.TotalPayment)

	reservation.PaymentStatus = "Pending"
	reservation.ReservationStatus = "Pending Approval"

	reservation.ID = len(reservations) + 1
	reservations = append(reservations, reservation)
	utils.WriteJSON("data/reservations.json", reservations)

	fmt.Println("Reservation made successfully!")
}

func CancelReservation() {
	var id int
	fmt.Print("Enter Reservation ID to cancel: ")
	fmt.Scanln(&id)
	for i, reservation := range reservations {
		if reservation.ID == id && reservation.ReservationStatus == "Pending Approval" {
			reservations = append(reservations[:i], reservations[i+1:]...)
			utils.WriteJSON("data/reservations.json", reservations)
			fmt.Println("Reservation canceled successfully!")
			return
		}
	}
	fmt.Println("Error canceling reservation: reservation not found or already confirmed")
}

func ViewMyReservations(customerName string) {
	fmt.Println("My Reservations:")
	for _, r := range reservations {
		if r.CustomerName == customerName {
			fmt.Printf("ID: %d, Room ID: %d, Total Payment: %d, Status: %s\n",
				r.ID, r.RoomID, r.TotalPayment, r.ReservationStatus)
		}
	}
}