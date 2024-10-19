package views

import (
	"fmt"
	"main/models"
)

func ShowMenu(menu *models.Menu) {
	fmt.Println("\033[33mDaftar menu di menu :\033[0m\n")
	if len(menu.GetFoods()) == 0 {
		fmt.Println("Tidak ada menu di menu.")
		return
	}
	fmt.Printf("\033[35m%-15s%-15s%-20s%-10s%-10s\033[0m\n", "Nama", "Jenis", "Harga", "Qty", "status")
	for _, s := range menu.GetFoods() {
		var status string
		var color string

		if s.Status {
			status = "Available"
			color = "\033[32m" // Warna hijau
		} else {
			status = "Unavailable"
			color = "\033[31m" // Warna merah
		}
		fmt.Printf("%s%-15s%-15s%-20.2f%-10d%-10s\033[0m\n", color, s.Name, s.Kinds, s.Price, s.Qty, status)
		// fmt.Printf("%-10s%-20s%-20.2f%-10d%s%-10s\033[0m\n", s.Name, s.Kinds, s.Price, s.Qty, color, status)
	}
}

