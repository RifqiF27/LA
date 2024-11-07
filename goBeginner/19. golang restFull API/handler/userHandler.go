package handlers

import (
	"book-store/collections"
	"book-store/config"
	"book-store/repository"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

var jwtSecret = []byte(config.GetJWTSecret())

func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
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

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user collections.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userFromDb, err := h.userRepo.GetUserLogin(user)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(userFromDb.ID, userFromDb.Username, userFromDb.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		Message: "Login successful",
		Token:   token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
