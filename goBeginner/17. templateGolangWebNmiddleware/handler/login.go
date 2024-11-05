package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"todo/model"
	"todo/repository"
	"todo/service"
	"todo/utils"
)

func Login(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			err := tmpl.ExecuteTemplate(w, "layout.html", nil)
			if err != nil {
				fmt.Println(err, ">>>>>")
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
			return
		}

		if _, err := os.Stat("session.json"); err == nil {
			utils.SendJSONResponse(w, http.StatusForbidden, "User already logged in", nil)
			return
		}

		user := model.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println(err, "<<<")
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
		if tmpl != nil {
			tmpl.Execute(w, sessionData)
		}
	}
}
