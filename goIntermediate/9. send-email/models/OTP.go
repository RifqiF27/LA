package models

import "time"

type OTP struct {
	ID         uint      `gorm:"primaryKey"`
	Email      string    `gorm:"not null"`
	Code       string    `gorm:"not null"`
	Expiration time.Time `gorm:"not null"`
}

type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ValidateOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}
