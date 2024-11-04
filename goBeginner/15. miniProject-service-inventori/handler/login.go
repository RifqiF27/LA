package handler

import (
	"database/sql"
	"fmt"
	"os"
	"service-inventory/model"
	"service-inventory/repository"
	"service-inventory/service"
	"service-inventory/utils"
)

func Login(db *sql.DB) {
	if _, err := os.Stat("session.json"); err == nil {
		utils.SendJSONResponse(403, "User already logged in", nil)
		return
	}
	user := model.User{}
	err := utils.DecodeJSONFile("body.json", &user)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	repo := repository.NewUserRepo(db)
	adminService := service.NewUserService(repo)

	admin, err := adminService.LoginService(user)

	if err != nil {
		utils.SendJSONResponse(404, err.Error(), nil)
	} else {
		utils.SendJSONResponse(200, "login success", admin)

		sessionData := map[string]interface{}{
			"ID":       admin.ID,
			"Username": admin.Username,
			"Role":     admin.Role,
		}

		err = utils.WriteJSONFile("session.json", sessionData)
		if err != nil {
			fmt.Println("Error menyimpan sesi:", err)
			return
		}

		fmt.Println("Sesi berhasil disimpan dalam session.json")
	}

}
