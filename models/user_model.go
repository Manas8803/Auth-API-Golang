package model

import (
	"Gin/Basics/configs"
	"crypto/rand"
	"fmt"
	"net/smtp"

	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	OTP      string `json:"otp"`
}

type OTP struct {
	Email string `json:"email" validate:"required"`
	OTP   string `json:"otp" validate:"required"`
}
type Register struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func CheckPassword(providedPassword string, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) GenerateOTP() error {
	randomBytes := make([]byte, 2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return err
	}

	otp := fmt.Sprintf("%06d", int(randomBytes[0])<<8|int(randomBytes[1])%1000000)
	user.OTP = otp

	return nil
}

func SendOTP(email string, otp string) error {

	auth := smtp.PlainAuth("", configs.EMAIL(), configs.PASSWORD(), "smtp.gmail.com")

	to := []string{email}

	message := []byte(
		"To:" + email + "\r\n" +
			"Subject: OTP for Registration\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
			"<html>" +
			"<head>" +
			"<title>OTP for Registration</title>" +
			"</head>" +
			"<body style=\"font-family: Arial, sans-serif;\">" +
			"<div style=\"padding: 20px;\">" +
			"<h1 style=\"color: #333;\">Welcome to our Service!</h1>" +
			"<p style=\"font-size: 16px;\">Your OTP for registration is: <strong>" + otp + "</strong></p>" +
			"<p>Ignore if you are not registered.</p>" +
			"</div>" +
			"</body>" +
			"</html>")

	err := smtp.SendMail("smtp.gmail.com:587", auth, configs.EMAIL(), to, message)

	return err
}
