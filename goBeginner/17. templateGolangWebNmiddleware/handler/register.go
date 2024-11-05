package handler

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	"html/template"
	"net/http"
	// "todo/model"
	"todo/repository"
	"todo/service"
	"todo/utils"
)

func Register(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Rendering registration form")
			err := tmpl.ExecuteTemplate(w, "layout.html",  map[string]bool{"IsRegister": true})
			if err != nil {
				fmt.Println(err, ">>>>>")
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}

			return
		}

		if r.Method != http.MethodPost {
			utils.SendJSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
			return
		}

		if r.Method == http.MethodPost {
			// user := model.User{}
			// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			// 	fmt.Println(err,"<<<<")
			//     utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
			//     return
			// }
			username := r.FormValue("username")
			password := r.FormValue("password")
			role := r.FormValue("role")

			if username == "" || password == "" || role == "" {
				utils.SendJSONResponse(w, http.StatusBadRequest, "username, password, and role are required", nil)
				return
			}

			repo := repository.NewUserRepo(db)
			userService := service.NewUserService(repo)

			err := userService.RegisterService(username, password, role)
			if err != nil {
				utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
				return
			}
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")

			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	}
}
