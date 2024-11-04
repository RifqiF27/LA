package services

import (
	"errors"
	"fmt"
	"main/models"
	"main/utils"
)

var rooms []models.Room

func init() {
    err := utils.ReadJSON("data/rooms.json", &rooms)
    if err != nil {
        panic("Could not load rooms data.")
    }
}

func ManageRooms() {
	var action int
	fmt.Println("Room Management:")
	fmt.Println("1. Add Room")
	fmt.Println("2. View Rooms")
	fmt.Println("3. Delete Room")
	fmt.Print("Select an action: ")
	fmt.Scanln(&action)

	switch action {
	case 1:
		var room models.Room
		fmt.Print("Enter Room Name: ")
		fmt.Scanln(&room.Name)
		fmt.Print("Enter room Price: ")
		fmt.Scanln(&room.Price)
		if err := AddRoom(room); err != nil {
			fmt.Println("Error adding Room:", err)
		} else {
			fmt.Println("Room added successfully!")
		}
	case 2:
		GetRooms()
	case 3:
		var room models.Room
		fmt.Print("Enter Room Name: ")
		fmt.Scanln(&room.Name)
		fmt.Print("Enter room Price: ")
		fmt.Scanln(&room.Price)
		if err := UpdateRoom(room); err != nil {
			fmt.Println("Error update Room:", err)
		} else {
			fmt.Println("Room update successfully!")
		}
	case 4:
		var room models.Room
		fmt.Print("Enter Room ID to delete: ")
		fmt.Scanln(&room.ID)
		DeleteRoom(room.ID)
	default:
		fmt.Println("Invalid action")
	}
}

func ViewAvailableRooms() {
	fmt.Println("Available Rooms:")
	for _, r := range rooms {
		if r.IsAvailable {
			fmt.Printf("ID: %d, Name: %s, Price: %d\n", r.ID, r.Name, r.Price)
		}
	}
}

func GetAvailableRooms() ([]models.Room, error) {
    availableRooms := []models.Room{}
    for _, room := range rooms {
        if room.IsAvailable {
            availableRooms = append(availableRooms, room)
        }
    }
    if len(availableRooms) == 0 {
        return nil, errors.New("no rooms available")
    }
    return availableRooms, nil
}

func SaveRooms() error {
    return utils.WriteJSON("data/rooms.json", rooms)
}

func AddRoom(room models.Room) error {
    if room.Name == "" || room.Price <= 0 {
        return errors.New("room name and price are required")
    }
    room.ID = len(rooms) + 1
    rooms = append(rooms, room)
    return utils.WriteJSON("data/rooms.json", rooms)
}

func GetRooms() []models.Room {
    return rooms
}

func UpdateRoom(updatedRoom models.Room) error {
    for i, room := range rooms {
        if room.ID == updatedRoom.ID {
            rooms[i] = updatedRoom
            return utils.WriteJSON("data/rooms.json", rooms)
        }
    }
    return errors.New("room not found")
}

func DeleteRoom(roomID int) error {
    for i, room := range rooms {
        if room.ID == roomID {
            rooms = append(rooms[:i], rooms[i+1:]...)
            return utils.WriteJSON("data/rooms.json", rooms)
        }
    }
    return errors.New("room not found")
}