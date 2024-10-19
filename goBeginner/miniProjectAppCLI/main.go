package main

import (
	"fmt"
	"main/controllers"
	"main/models"
	"main/utils"
	"main/views"
	"reflect"
)

func main() {

	var menu models.Menu
	menu.Init()
	var option int

	for {
		fmt.Println("\n**Aplikasi Manajemen Makanan**")
		fmt.Println("1. Tambah Menu")
		fmt.Println("2. Cari Menu")
		fmt.Println("3. Hapus Menu")
		fmt.Println("4. Update Menu")
		fmt.Println("5. Beli Makanan")
		fmt.Println("6. Tampilkan Menu")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scanln(&option)

		fmt.Printf("Tipe data option: %s\n\n", reflect.TypeOf(option))
		utils.ClearScreen()

		switch option {
		case 1:
			controllers.AddMenuController(&menu)
		case 2:
			controllers.SearchMenuController(&menu)
		case 3:
			controllers.DeleteMenuController(&menu)
		case 4:
			controllers.UpdateMenuController(&menu)
		case 5:
			controllers.PurchaseMenuController(&menu)
		case 6:
			views.ShowMenu(&menu)
			fmt.Printf("\nTipe data: %s\n", reflect.TypeOf(menu.GetFoods()))
		case 7:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
