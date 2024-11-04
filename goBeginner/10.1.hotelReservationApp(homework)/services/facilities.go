package services

import (
	"errors"
	"fmt"
	"main/models"
	"main/utils"
)

var facilities []models.Facility

func init() {
	err := utils.ReadJSON("data/facilities.json", &facilities)
	if err != nil {
		panic("Could not load facilities data.")
	}
}

func ManageFacilities() {
	var action int
	fmt.Println("Facility Management:")
	fmt.Println("1. Add Facility")
	fmt.Println("2. View Facilities")
	fmt.Println("3. Delete Facility")
	fmt.Print("Select an action: ")
	fmt.Scanln(&action)

	switch action {
	case 1:
		var facility models.Facility
		fmt.Print("Enter Facility Name: ")
		fmt.Scanln(&facility.Name)
		fmt.Print("Enter Facility Price: ")
		fmt.Scanln(&facility.Price)
		if err := AddFacility(facility); err != nil {
			fmt.Println("Error adding facility:", err)
		} else {
			fmt.Println("Facility added successfully!")
		}
	case 2:
		ViewFacilities()
	case 3:
		var facility models.Facility
		fmt.Print("Enter Facility Name: ")
		fmt.Scanln(&facility.Name)
		fmt.Print("Enter Facility Price: ")
		fmt.Scanln(&facility.Price)
		if err := UpdateFacility(facility); err != nil {
			fmt.Println("Error update facility:", err)
		} else {
			fmt.Println("Facility update successfully!")
		}
	case 4:
		DeleteFacility()
	default:
		fmt.Println("Invalid action")
	}
}

func AddFacility(facility models.Facility) error {
	if facility.Name == "" || facility.Price <= 0 {
		return errors.New("facility name and price are required")
	}
	facility.ID = len(facilities) + 1
	facilities = append(facilities, facility)
	return utils.WriteJSON("data/facilities.json", facilities)
}

func ViewFacilities() {
	fmt.Println("Facilities:")
	for _, f := range facilities {
		fmt.Printf("ID: %d, Name: %s, Price: %d\n", f.ID, f.Name, f.Price)
	}
}

func DeleteFacility() {
	var id int
	fmt.Print("Enter Facility ID to delete: ")
	fmt.Scanln(&id)
	for i, facility := range facilities {
		if facility.ID == id {
			facilities = append(facilities[:i], facilities[i+1:]...)
			utils.WriteJSON("data/facilities.json", facilities)
			fmt.Println("Facility deleted successfully!")
			return
		}
	}
	fmt.Println("Error deleting facility: facility not found")
}

func UpdateFacility(updatedFacility models.Facility) error {
    for i, facility := range facilities {
        if facility.ID == updatedFacility.ID {
            facilities[i] = updatedFacility
            return utils.WriteJSON("data/facilities.json", facilities)
        }
    }
    return errors.New("facility not found")
}
