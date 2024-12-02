package controller

import (
	"net/http"
	"voucher_system/database"
	"voucher_system/helper"
	"voucher_system/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	Service service.Service
	log     *zap.Logger
	Cacher  database.Cacher
}

func NewAuthController(service service.Service, log *zap.Logger, cacher database.Cacher) AuthController {
	return AuthController{Service: service, log: log, Cacher: cacher}
}

func (a *AuthController) Login(c *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseError(c, err.Error(), "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := a.Service.User.Login(req.Email)
	if err != nil {
		helper.ResponseError(c, "User not found", "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !helper.CheckPassword(req.Password, user.Password) {
		helper.ResponseError(c, "Invalid password", "Invalid email or password", http.StatusUnauthorized)
		return
	}
	// userIDstr := helper.IntToString(user.ID)
	token := helper.GenerateToken()
	err = a.Cacher.SaveToken(user.Email, token)
	if err != nil {
		helper.ResponseError(c, err.Error(), "Failed to save token", http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, gin.H{
		"email": user.Email,
		"token": token,
	}, "Login Success")
}
