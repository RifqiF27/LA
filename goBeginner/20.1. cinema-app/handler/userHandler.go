package handler

import (
	"cinema/model"
	"cinema/service"
	"cinema/utils"
	"cinema/validation"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validation.ValidateUser(&user, false); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err := h.authService.Register(user)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusConflict, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "Registration successful", nil)

}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validation.ValidateUser(&user, true); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	users, err := h.authService.Login(user)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.SendJSONResponse(w, http.StatusOK, "login success", users)
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	fmt.Println("Token received:", token)
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}
	fmt.Println("Token after removal of Bearer prefix:", token)

	if token == "" {
		utils.SendJSONResponse(w, http.StatusUnauthorized, "Token required", nil)
		return
	}

	_, err := h.authService.VerifyToken(token)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusUnauthorized, "Invalid or expired token", nil)
		return
	}

	if err := h.authService.Logout(token); err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "Logout failed", nil)
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "Logout success", nil)
}
