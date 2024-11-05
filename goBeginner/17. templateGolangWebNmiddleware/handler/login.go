package handler

import (
	"database/sql"
	// "encoding/json"
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
			fmt.Println("Rendering login form")

			err := tmpl.ExecuteTemplate(w, "layout.html",  map[string]bool{"IsLogin": true})
			if err != nil {
				fmt.Println(err, ">>>>>")
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
			// http.Redirect(w, r, "/register", http.StatusSeeOther)
			return

		}

		if r.Method != http.MethodPost {
			utils.SendJSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
			return
		}

		if _, err := os.Stat("session.json"); err == nil {
			utils.SendJSONResponse(w, http.StatusForbidden, "User already logged in", nil)
			return
		}
		if r.Method == http.MethodPost {

			// user := model.User{}
			// err := json.NewDecoder(r.Body).Decode(&user)
			// if err != nil {
			// 	fmt.Println(err, "<<<")
			// 	utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
			// 	return
			// }

			username := r.FormValue("username")
			password := r.FormValue("password")

			if username == "" || password == "" {
				utils.SendJSONResponse(w, http.StatusBadRequest, "Username and password are required", nil)
				return
			}

			user := model.User{Username: username, Password: password}

			repo := repository.NewUserRepo(db)
			adminService := service.NewUserService(repo)

			_, err := adminService.LoginService(user)
			if err != nil {
				utils.SendJSONResponse(w, http.StatusNotFound, err.Error(), nil)
				return
			}


			// sessionData := map[string]interface{}{
			// 	"ID":       admin.ID,
			// 	"Username": admin.Username,
			// 	"Role":     admin.Role,
			// }

			// err = utils.WriteJSONFile("session.json", sessionData)
			// if err != nil {
			// 	http.Error(w, "Error saving session", http.StatusInternalServerError)
			// 	return
			// }
			// if tmpl != nil {
			// 	tmpl.Execute(w, sessionData)
			// }

			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
			w.Header().Set("role", "admin_token")
			w.Header().Set("token", "12345")
			http.Redirect(w, r, "/todo/get-todo", http.StatusSeeOther)
			// utils.SendJSONResponse(w, http.StatusOK, "Login successful", admin)

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
