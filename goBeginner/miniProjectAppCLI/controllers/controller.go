package controllers

import (
	"fmt"
	"main/models"
	"main/utils"
	"strings"
)

func AddMenuController(menu models.Item) {
	defer fmt.Println("Proses update menu selesai")
	var name string
	var kinds string
	var price float64
	var qty int
	// var status bool
	var err error

	name, err = utils.InputStr("Masukan nama makanan!\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Scanln(&name)

	kinds, err = utils.InputStr("Masukan jenis makanan!\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Scanln(&kinds)

	price, err = utils.InputFloat("Masukkan harga!\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	qty, err = utils.InputInt("Masukkan qty\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// status, err = utils.InputBool("Masukkan status\n")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	food := &models.Food{
		Name:   name,
		Kinds:  kinds,
		Price:  price,
		Qty:    qty,
		Status: true,
	}

	menu.AddMenu(food)
	fmt.Println("Menu berhasil ditambahkan")
}

func SearchMenuController(menu models.Item) {
	defer fmt.Println("Proses pencarian menu selesai.")
	var name string

	fmt.Print("Masukkan nama menu yang ingin dicari: ")
	fmt.Scanln(&name)

	f, err := menu.SearchMenu(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\033[35m%-15s%-15s%-20s%-10s%-10s\033[0m\n", "Nama", "Jenis", "Harga", "Qty", "status")
	fmt.Println(strings.Repeat("-", 75))

	var color string
	var status string
	if f.Status {
		status = "Available"
		color = "\033[32m" // Warna hijau
	} else {
		status = "Unavailable"
		color = "\033[31m" // Warna merah
	}
	fmt.Printf("%s%-15s%-15s%-20.2f%-10d%-10s\033[0m\n", color, f.Name, f.Kinds, f.Price, f.Qty, status)
	fmt.Println(strings.Repeat("-", 75))
}

func DeleteMenuController(menu models.Item) {
	defer fmt.Println("Proses hapus menu selesai.")
	var name string
	fmt.Print("Masukkan nama menu yang ingin dihapus: ")
	fmt.Scanln(&name)

	err := menu.DeleteMenu(name)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Menu berhasil dihapus!")
	}
}

func UpdateMenuController(menu models.Item) {
	defer fmt.Println("Proses update menu selesai")
	var name string
	var kinds string
	var price float64
	var qty int
	var status bool
	var err error

	fmt.Print("Masukan nama yang ingin diupdate: ")
	fmt.Scanln(&name)

	f, err := menu.SearchMenu(name)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("\033[35m%-15s%-15s%-20s%-10s%-10s\033[0m\n", "Nama", "Jenis", "Harga", "Qty", "status")
	fmt.Println(strings.Repeat("-", 75))
	var color string
	var statusStr string
	if f.Status {
		statusStr = "Available"
		color = "\033[32m" // Warna hijau
	} else {
		statusStr = "Unavailable"
		color = "\033[31m" // Warna merah
	}
	fmt.Printf("%s%-15s%-15s%-20.2f%-10d%-10s\033[0m\n", color, f.Name, f.Kinds, f.Price, f.Qty, statusStr)
	fmt.Println(strings.Repeat("-", 75))

	name, err = utils.InputStr("Masukan nama makanan yang ingin diupdate!\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	kinds, err = utils.InputStr("Masukan jenis makanan yang ingin diupdate!\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	price, err = utils.InputFloat("Masukkan harga yang ingin diupdate!\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	qty, err = utils.InputInt("Masukkan qty yang ingin diupdate!\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// status, err = utils.InputBool("Masukkan status yang ingin diupdate!\n")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	f.Name = name
	f.Kinds = kinds
	f.Price = price
	f.Qty = qty
	if f.Qty <= 0 {
		f.Status = false
	}

	err = menu.UpdateMenu(name, kinds, price, qty, status)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Menu berhasil diupdate")
	}

}
