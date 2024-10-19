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
