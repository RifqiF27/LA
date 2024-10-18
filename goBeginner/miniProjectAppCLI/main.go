package main

import (
	"fmt"
	"main/controllers"
	"main/models"
	"main/views"
	// "reflect"
)

func main() {
	// items := models.Menu{}
	// var name, kinds string
	// var price float64
	// var qty int
	// var status bool

	// fmt.Println("Masukkan nama makanan:")
	// fmt.Scanln(&name)
	// fmt.Println("Masukkan deskripsi makanan:")
	// fmt.Scanln(&kinds)
	// fmt.Println("Masukkan harga makanan:")
	// fmt.Scanln(&price)
	// fmt.Println("Masukkan qty makanan:")
	// fmt.Scanln(&qty)
	// fmt.Println("Masukkan status makanan (true untuk tersedia, false untuk tidak tersedia):")
	// fmt.Scanln(&status)
	// controllers.AddMenuController(&items, "chick","lorem",400.20,12,true)
	// controllers.AddMenuController(&items, "chick","lorem",400.20,4,true)
	// controllers.AddMenuController(&items, name, kinds, price, qty, status)
	// for _, item := range items.Food {
	// 	fmt.Printf("Name: %s, kinds: %s, Price: %.2f,Qty: %d, Status: %t\n", item.Name, item.Kinds, item.Price,item.Qty, item.Status)
	// }
	// fmt.Println(items)
	var menu models.Menu
	var pilihan int

	for {
		fmt.Println("\nAplikasi Manajemen Makanan")
		fmt.Println("1. Tambah Menu")
		fmt.Println("2. Cari Menu")
		fmt.Println("3. Hapus Menu")
		fmt.Println("4. Update Menu")
		fmt.Println("5. Tampilkan Menu")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scanln(&pilihan)

		// fmt.Println("Tipe data input:", reflect.TypeOf(pilihan))

		switch pilihan {
		case 1:
			controllers.AddMenuController(&menu)
		case 2:
			controllers.SearchMenuController(&menu)
		case 3:
			controllers.DeleteMenuController(&menu)
		case 4:
			controllers.UpdateMenuController(&menu)
		case 5:
			views.ShowMenu(&menu)
		case 6:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
