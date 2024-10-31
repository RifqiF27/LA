package main

import (
	"fmt"
	"main/database"
	"main/handler"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		fmt.Println("Gagal menginisialisasi database:", err)
		return
	}
	defer db.Close()

	for {

		var endpoint string
		fmt.Print("Masukkan endpoint: ")
		fmt.Scan(&endpoint)

		switch endpoint {
		case "login":
			handler.Login(db)
		case "register":
			handler.Register(db)
		case "logout":
			handler.Logout()
		default:
			fmt.Println("Endpoint tidak dikenal")
		}
	}
}
