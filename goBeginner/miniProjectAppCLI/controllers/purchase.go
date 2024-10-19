package controllers

import (
	"fmt"
	"main/models"
)

func PurchaseMenuController(menu *models.Menu) {
	defer fmt.Println("Proses pembelian selesai.")
	var itemNames []string
	var itemName string
	var qty int
	var totalPrice float64

	fmt.Println("Masukkan nama makanan yang ingin dibeli (ketik 'bayar' untuk mengakhiri):")

	for {
		fmt.Print("Nama Makanan: ")
		fmt.Scanln(&itemName)
		if itemName == "bayar" {
			break
		}

		food, err := menu.SearchMenu(itemName)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Print("Jumlah yang ingin dibeli: ")
		fmt.Scanln(&qty)

		if qty > food.Qty {
			fmt.Println("Jumlah melebihi stok yang tersedia.")
			continue
		}

		food.Qty -= qty // Mengurangi jumlah yang tersedia
		if food.Qty <= 0 {
			food.Status = false
		}
		totalPrice += food.Price * float64(qty)
		itemNames = append(itemNames, itemName)
	}

	if totalPrice >= 500000 {
		discount := totalPrice * 0.30 // Diskon 30%
		totalPriceDisc := totalPrice - discount
		tax := totalPriceDisc * 0.11
		total := totalPriceDisc + tax
		fmt.Printf("Diskon 30%% diterapkan!\nTotal Harga: %.2f\nDiscount: %.2f\nPajak: %.2f\nTotal: %.2f\n", totalPrice,discount, tax, total)
	} else if totalPrice >= 200000 && len(itemNames) > 2 {
		discount := totalPrice * 0.10 // Diskon 10%
		totalPriceDisc := totalPrice - discount
		tax := totalPriceDisc * 0.11
		total := totalPriceDisc + tax
		fmt.Printf("Diskon 10%% diterapkan!\nTotal Harga: %.2f\nDiscount: %.2f\nPajak: %.2f\nTotal: %.2f\n", totalPrice,discount, tax, total)
	} else {
		fmt.Printf("Total harga: %.2f\n", totalPrice)
	}
}
