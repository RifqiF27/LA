package handler

import (
	"book-store/collections"
	"book-store/config"
	"book-store/service"
	"encoding/json"
	"fmt"

	// "fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	service service.UserService
	tmpl    *template.Template
}

var jwtSecret = []byte(config.GetJWTSecret())

func NewAuthHandler(service service.UserService) *AuthHandler {

	tmpl, err := template.ParseFiles("view/login.html")
	if err != nil {
		panic(err)
	}
	return &AuthHandler{service: service, tmpl: tmpl}
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

func GenerateJWT(userID int, username, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	err := h.tmpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// var user collections.User

	// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	// 	fmt.Println(err,"<<<")
	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	// 	return
	// }

	if r.Method == http.MethodGet{
		fmt.Println("login endpoint accessed")
		h.LoginPage(w, r)
		return
	}

	if r.Method == http.MethodPost{
		var user collections.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

		// user.Username = r.FormValue("username")
		// user.Password = r.FormValue("password")
	
		// user := collections.User{Username: username, Password: password}
	
	
		userFromDb, err := h.service.LoginService(user)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
	
		token, err := GenerateJWT(userFromDb.ID, userFromDb.Username, userFromDb.Role)
		fmt.Println("Response data:", token)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}
	
		response := LoginResponse{
			Message: "Login successful",
			Token:   token,
		}
		w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

	}
}
