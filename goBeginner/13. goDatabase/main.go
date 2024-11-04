package main

import (
	"fmt"
	// "log"
	"main/database"
	"main/service"

	_ "github.com/lib/pq"
)

func main() {

	db, err := database.InitDb()
	if err != nil {
		fmt.Println("Gagal menginisialisasi database:", err)
		return
	}

	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Gagal memulai transaksi:", err)
		return
	}
	
	// input Customer and Driver
	userID, role, err := service.InputDataUser(db, "testCust10", "testpass", "customer")
	if err != nil {
		tx.Rollback()
		fmt.Println("Error saat membuat user:", err)
		return
	}
	fmt.Println("User berhasil ditambahkan")

	if role == "driver" {
		err = service.InputDataDriver(tx, "Driver5", "1234567890", "Driver Address", "Type Vehicle", userID)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error saat membuat driver:", err)
			return
		}
		fmt.Println("Driver berhasil ditambahkan")

	} else if role == "customer" {
		err = service.InputDataCustomer(tx, "", "0987654321", "Customer Address", userID)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error saat membuat customer:", err)
			return
		}
		fmt.Println("Customer berhasil ditambahkan")
	}

	if err := tx.Commit(); err != nil {
		fmt.Println("Gagal melakukan commit transaksi user:", err)
		return
	}

	// tx, err := db.Begin()
	// if err != nil {
	// 	fmt.Println("Gagal memulai transaksi:", err)
	// 	return
	// }

	// input Order
	// err = service.InputDataOrder(db, 1, 1, 1, "2024-12-30 12:00:00", true)
	// if err != nil {
	// 	tx.Rollback()
	// 	fmt.Println("Error saat membuat order:", err)
	// 	return
	// }
	// if err := tx.Commit(); err != nil {
	// 	fmt.Println("Gagal melakukan commit transaksi order:", err)
	// 	return
	// }
	// fmt.Println("berhasil membuat orderan")

	// Get all order
	// if err := service.GetAllOrders(db); err != nil {
	// 	log.Fatalf("gagal mengambil semua order: %v", err)
	// }
	// // total data order every month
	// if err := service.GetAllSummaryOrdersTotal(db); err != nil {
	// 	log.Fatalf("gagal mengambil semua order total: %v", err)
	// }

	// // Total order by customer every month
	// if err := service.GetAllSummaryOrdersCustomer(db); err != nil {
	// 	log.Fatalf("gagal mengambil semua order customer: %v", err)
	// }

	// // Often order every month based on region
	// if err := service.GetAllRegionOrderSummary(db); err != nil {
	// 	log.Fatalf("gagal mengambil semua data: %v", err)
	// }

	// // Often order by time
	// if err := service.GetAllSummaryOrdersTime(db); err != nil {
	// 	log.Fatalf("gagal mengambil semua data: %v", err)
	// }

	// Driver often take orders every month
	// if err := service.GetAllSummaryOrdersDriver(db); err != nil {
	// 	log.Fatalf("gagal mengambil semua data: %v", err)
	// }

	// if err := service.GetUserIsActive(db); err != nil {
	// 	log.Fatalf("gagal mengambil semua data: %v", err)
	// }
}
