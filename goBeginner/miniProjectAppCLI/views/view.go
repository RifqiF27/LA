package views

import (
	"fmt"
	"main/models"
	"strings"
)

func ShowMenu(menu *models.Menu) {
	fmt.Printf("\033[33mDaftar menu di menu :\033[0m\n")
	if len(menu.GetFoods()) == 0 {
		fmt.Println("Tidak ada menu di menu.")
		return
	}
	fmt.Printf("\033[35m%-15s%-15s%-20s%-10s%-10s\033[0m\n", "Nama", "Jenis", "Harga", "Qty", "status")
	fmt.Println(strings.Repeat("-", 75))
	for _, s := range menu.GetFoods() {
		var status, color string

		if s.Status {
			status = "Available"
			color = "\033[32m" // Warna hijau
		} else {
			status = "Unavailable"
			color = "\033[31m" // Warna merah
		}
		fmt.Printf("%s%-15s%-15s%-20.2f%-10d%-10s\033[0m\n", color, s.Name, s.Kinds, s.Price, s.Qty, status)
		fmt.Println(strings.Repeat("-", 75))

	}
}
