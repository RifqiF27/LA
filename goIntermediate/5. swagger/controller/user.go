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

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" binding:"required" example:"password1234"`
}

// Login godoc
// @Summary Login user
// @Description Authenticate user with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "Login request payload"
// @Success 200 {object} utils.ResponseOK{data=utils.LoginResponse} "Successful login"
// @Failure 400 {object} utils.ErrorResponse "Invalid input"
// @Failure 401 {object} utils.ErrorResponse "Invalid email or password"
// @Failure 500 {object} utils.ErrorResponse "Failed to save token"
// @Router /login [post]
func (a *AuthController) Login(c *gin.Context) {

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
	userIDstr := helper.IntToString(user.ID)
	token := helper.GenerateToken()
	err = a.Cacher.SaveToken(userIDstr, token)
	if err != nil {
		helper.ResponseError(c, err.Error(), "Failed to save token", http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, gin.H{
		"id": userIDstr,
		"token": token,
	}, "Login Success", http.StatusOK)
}