package usercontroller

import (
	"dashboard-ecommerce-team2/config"
	"dashboard-ecommerce-team2/database"
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	utils "dashboard-ecommerce-team2/util"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	Service service.Service
	Log     *zap.Logger
	Cacher  database.Cacher
	Config  config.Configuration
}

func NewUserController(service service.Service, log *zap.Logger, cacher database.Cacher, config config.Configuration) *UserController {
	return &UserController{
		Service: service,
		Log:     log,
		Cacher:  cacher,
		Config:  config,
	}
}

// CreateUserController godoc
// @Summary      Create a new user
// @Description  Register a new user with a provided request body
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        registerRequest  body     models.RegisterRequest  true  "User Registration Request Body"
// @Success      201             {object} helper.HTTPResponse   "User created successfully"
// @Failure      400             {object} helper.HTTPResponse   "Invalid request body"
// @Failure      500             {object} helper.HTTPResponse   "Failed to create user"
// @Router       /auth/register [post]
func (ctrl *UserController) CreateUserController(c *gin.Context) {
	var registerReq models.RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}

	err := ctrl.Service.User.CreateUser(registerReq)
	if err != nil {
		ctrl.Log.Error("Failed to create user", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to create user", http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(c, nil, "User created successfully", http.StatusCreated)
}

// LoginController godoc
// @Summary      User Login
// @Description  Authenticate a user and return a token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        loginRequest  body     models.LoginRequest  true  "Login Request Body"
// @Success      200           {object} helper.HTTPResponse   "User logged in successfully"
// @Failure      400           {object} helper.HTTPResponse   "Invalid request body"
// @Failure      401           {object} helper.HTTPResponse   "Failed to login user"
// @Router       /auth/login [post]
func (ctrl *UserController) LoginController(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}
	user, err := ctrl.Service.User.Login(loginReq)
	if err != nil {
		ctrl.Log.Error("Failed to login user", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to login user", http.StatusUnauthorized)
		return
	}

	userIDStr := helper.IntToString(user.ID)
	key := fmt.Sprintf("UserID-%s", userIDStr)
	token := helper.GenerateToken(userIDStr, ctrl.Config.SecretKey)
	loginResponse := utils.LoginResponse{
		ID:    key,
		Role:  user.Role,
		Token: token,
	}

	ctrl.Cacher.SaveToken(key, token)
	helper.ResponseOK(c, loginResponse, "User logged in successfully", http.StatusOK)
}

// CheckEmailUserController godoc
// @Summary      Check if email is already registered
// @Description  Verify if a user with the given email already exists in the system
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        email  body     models.CheckEmailRequest  true  "Email to check"
// @Success      200    {object} helper.HTTPResponse     "Email check result"
// @Failure      400    {object} helper.HTTPResponse     "Invalid request body"
// @Failure      500    {object} helper.HTTPResponse     "Failed to check user email"
// @Router       /auth/check-email [post]
func (ctrl *UserController) CheckEmailUserController(c *gin.Context) {
	request := models.CheckEmailRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}

	existedUser, err := ctrl.Service.User.CheckUserEmail(request.Email)
	if err != nil {
		ctrl.Log.Error("Failed to check user email", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to check user email", http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(c, existedUser, "User email exists", http.StatusOK)
}

// ResetUserPasswordController godoc
// @Summary      Reset user password
// @Description  Reset the password for a user using a provided request body
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        resetRequest  body     models.LoginRequest  true  "User password reset request body"
// @Success      200           {object} helper.HTTPResponse   "User password reset successfully"
// @Failure      400           {object} helper.HTTPResponse   "Invalid request body"
// @Failure      500           {object} helper.HTTPResponse   "Failed to reset user password"
// @Router       /auth/reset-password [PATCH]
func (ctrl *UserController) ResetUserPasswordController(c *gin.Context) {
	var req struct {
		Email       string `json:"email" binding:"required,email"`
		OTP         string `json:"otp" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseError(c, err.Error(), "Permintaan tidak valid", http.StatusBadRequest)
		return
	}

	cacheKey := fmt.Sprintf("otp:%s", req.Email)
	storedOtp, err := ctrl.Cacher.Get(cacheKey)
	if err != nil {
		helper.ResponseError(c, "OTP tidak ditemukan atau sudah kedaluwarsa", "OTP tidak valid", http.StatusUnauthorized)
		return
	}

	if storedOtp != req.OTP {
		helper.ResponseError(c, "OTP tidak valid", "OTP salah", http.StatusUnauthorized)
		return
	}

	// if err := ctrl.Cacher.Delete(cacheKey); err != nil {
	// 	fmt.Printf("Peringatan: Gagal menghapus OTP untuk %s dari Redis: %v\n", req.Email, err)
	// }

	hashedPassword := helper.HashPassword(req.NewPassword)
	if err := ctrl.Service.User.ResetUserPassword(req.Email, storedOtp, hashedPassword); err != nil {
		helper.ResponseError(c, "Gagal memperbarui password", err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, nil, "Password berhasil direset", http.StatusOK)
}

func (ctrl *UserController) RequestReset(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseError(c, err.Error(), "Permintaan tidak valid", http.StatusBadRequest)
		return
	}

	user, err := ctrl.Service.User.CheckUserEmail(req.Email)
	if err != nil || user == nil {
		helper.ResponseError(c, "Email tidak terdaftar", "Email tidak valid", http.StatusNotFound)
		return
	}

	otp := helper.GenerateOTP(6)
	cacheKey := fmt.Sprintf("otp:%s", req.Email)
	expiration := 5 * time.Minute

	log.Println("cacheKey: ", cacheKey, "<<<<<<")
	log.Println("otp: ", otp, "<<<<<<")

	if err := ctrl.Cacher.Set(cacheKey, otp, expiration); err != nil {
		helper.ResponseError(c, "Gagal menyimpan OTP ke Redis", "Kesalahan server", http.StatusInternalServerError)
		return
	}

	go func() {
		err := helper.SendDummyEmail(req.Email, "Reset Password OTP", fmt.Sprintf("Your OTP is: %s", otp))
		if err != nil {
			fmt.Printf("Gagal mengirim email ke %s: %v\n", req.Email, err)
		}
	}()

	helper.ResponseOK(c, nil, "OTP berhasil dikirim ke email", http.StatusOK)
}
