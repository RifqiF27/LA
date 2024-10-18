package views

import (
	"fmt"
	"main/models"
)

func ShowMenu(menu *models.Menu) {
	fmt.Println("Daftar menu di menu:")
	if len(menu.Food) == 0 {
		fmt.Println("Tidak ada menu di menu.")
		return
	}
	fmt.Printf("%-10s%-20s%-20s%-10s%-10s\n", "Nama", "Jenis", "Harga", "Qty", "status")
	for _, s := range menu.Food {
		var status string
		var color string

		if s.Status {
			status = "Available"
			color = "\033[32m" // Warna hijau
		} else {
			status = "Unavailable"
			color = "\033[31m" // Warna merah
		}

		fmt.Printf("%-10s%-20s%-20.2f%-10d%s%-10s\033[0m\n", s.Name, s.Kinds, s.Price, s.Qty, color, status)
	}
}
