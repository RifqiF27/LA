package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"todo/model"
	"todo/repository"
	"todo/service"
	"todo/utils"
)

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if _, err := os.Stat("session.json"); err == nil {
			utils.SendJSONResponse(w, http.StatusForbidden, "User already logged in", nil)
			return
		}

		user := model.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
			return
		}

		repo := repository.NewUserRepo(db)
		adminService := service.NewUserService(repo)

		admin, err := adminService.LoginService(user)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusNotFound, err.Error(), nil)
			return
		}

		utils.SendJSONResponse(w, http.StatusOK, "Login successful", admin)

		sessionData := map[string]interface{}{
			"ID":       admin.ID,
			"Username": admin.Username,
			"Role":     admin.Role,
		}

		err = utils.WriteJSONFile("session.json", sessionData)
		if err != nil {
			http.Error(w, "Error saving session", http.StatusInternalServerError)
			return
		}
	}
}
