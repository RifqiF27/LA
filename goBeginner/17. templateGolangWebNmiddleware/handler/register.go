package handler

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"todo/model"
	"todo/repository"
	"todo/service"
	"todo/utils"
)

func Register(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Hanya izinkan metode POST
		if r.Method != http.MethodPost {
			utils.SendJSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
			return
		}

		user := model.User{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
			return
		}

		// Validasi data
		if user.Username == "" || user.Password == "" || user.Role == "" {
			utils.SendJSONResponse(w, http.StatusBadRequest, "username, password, and role are required", nil)
			return
		}

		// Inisialisasi repository dan service
		repo := repository.NewUserRepo(db)
		userService := service.NewUserService(repo)

		// Panggil RegisterService untuk menambahkan pengguna baru
		err := userService.RegisterService(user.Username, user.Password, user.Role)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		utils.SendJSONResponse(w, http.StatusCreated, "Register success", nil)

		if tmpl != nil {
            tmpl.Execute(w, nil)
        }
	}
}
