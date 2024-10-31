package handler

import (
	"database/sql"
	"fmt"
	"main/reposiitory"
	"main/service"
	"main/utils"
	"strings"
)

func Register(db *sql.DB) {
	var username, password, role, name string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)

	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	fmt.Print("Masukkan role: ")
	fmt.Scan(&role)
	if strings.ToLower(role) == "admin" {
		fmt.Print("Masukkan name: ")
		fmt.Scan(&name)	
	}

	repo := repository.NewUserRepo(db)
	userService := service.NewUserService(repo)

	err := userService.RegisterService(username, password, role, name)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	if role == "admin" {
		utils.SendJSONResponse(201, "register success "+name, nil)
	} else {
		utils.SendJSONResponse(201, "register success "+username, nil)
	}

}
