package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "io"
	"main/model"
	"main/reposiitory"
	"main/service"
	"main/utils"
	"os"
)

func Login(db *sql.DB) {

	// input
	// user := model.User{}
	// file, err := os.Open("body.json")

	// if err != nil {
	// 	fmt.Println("Error : ", err)
	// }

	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&user)
	// if err != nil && err != io.EOF {
	// 	fmt.Println("error decoding JSON: ", err)
	// }

	var username, password string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)

	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	user := model.User{
		Username: username,
		Password: password,
	}

	// proses
	repo := repository.NewUserRepo(db)
	adminService := service.NewUserService(repo)

	admin, err := adminService.LoginService(user)

	if err != nil {
		utils.SendJSONResponse(404, "account not found", nil)
	} else {
		utils.SendJSONResponse(200, "login success", admin)

		sessionData := map[string]interface{}{
			"ID":       admin.ID,
			"Username": admin.Username,
			"Role":     admin.Role,
			
		}
	
		sessionJSON, err := json.MarshalIndent(sessionData, "", "  ")
		if err != nil {
			fmt.Println("Gagal membuat data sesi:", err)
			return
		}
	
		err = os.WriteFile("session.json", sessionJSON, 0644)
		if err != nil {
			fmt.Println("Gagal menyimpan sesi:", err)
			return
		}
	
		fmt.Println("Sesi berhasil disimpan dalam session.json")
	}

}