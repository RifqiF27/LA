package main

import (
	"fmt"
	"main/controllers"
	"main/models"
	"main/utils"
	"reflect"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\033[31mProgram mengalami kesalahan yang tidak terduga dan tertangani oleh recover.\033[0m")
			fmt.Printf("\033[31mError: %v\033[0m\n", r)
		}
	}()

	var book models.MyLibrary
	book.Init()
	var option int

	for {
		fmt.Println("\n**Aplikasi Manajemen Perpustakaan**")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Hapus Buku")
		fmt.Println("3. Tampilkan Buku")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu (masukan angka untuk memilih menu): ")

		fmt.Scanln(&option)

		fmt.Printf("Tipe data option: %s\n\n", reflect.TypeOf(option))
		utils.ClearScreen()

		switch option {
		case 1:
			controllers.AddBookController(&book)
		case 2:
			controllers.RemoveBookController(&book)
		case 3:
			controllers.ShowBookController(&book)
		case 4:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("\033[031mPilihan tidak valid, coba lagi!\033[0m")
		}
	}
}
