package controller

import (
	"Gin/Basics/auth"
	"Gin/Basics/configs"
	db "Gin/Basics/db/sqlconfig"
	model "Gin/Basics/models"
	"Gin/Basics/responses"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func Login(r *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var req model.Login

	//* Checking for invalid json format
	if err := r.BindJSON(&req); err != nil {
		respondWithError(r, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	//* Validating if all the fields are present
	if validationErr := validate.Struct(&req); validationErr != nil {
		respondWithError(r, http.StatusBadRequest, "Please provide the required credentials.")
		return
	}

	queries := db.New(configs.CONN)
	//* Checking whether the user is registered
	user, err := queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(r, http.StatusNotFound, "User is not registered.")
			return
		}

		respondWithError(r, http.StatusInternalServerError, "Internal server error")
		return
	}

	//* Checking for verification of the user
	if !user.Isverified {
		go func() {
			model.SendOTP(user.Email, user.Otp)
		}()
		respondWithError(r, http.StatusUnprocessableEntity, "Email already registered. Please verify your rmail address using the OTP sent to your registered email.")
		return
	}

	//* Verifying password
	credentialsError := model.CheckPassword(req.Password, user.Password)
	if credentialsError != nil {
		respondWithError(r, http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	//* Generating Token
	var token string
	token, err = auth.GenerateJWT(user.Email)
	if err != nil {
		respondWithError(r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	r.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusAccepted, Message: "success", Data: map[string]interface{}{"token": token}})
}

func Register(r *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	var user model.User
	defer cancel()
	var queries *db.Queries
	go func() {
		queries = db.New(configs.CONN)
	}()

	//* Checking for invalid json format
	if invalidJsonErr := r.BindJSON(&user); invalidJsonErr != nil {
		respondWithError(r, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	//* Validating if all the fields are present
	if validationErr := validate.Struct(&user); validationErr != nil {
		respondWithError(r, http.StatusBadRequest, "Please provide with sufficient credentials")
		return
	}

	//* Hashing Password
	if hashPassErr := user.HashPassword(user.Password); hashPassErr != nil {
		respondWithError(r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	//* Generating OTP
	if genOtpErr := user.GenerateOTP(); genOtpErr != nil {
		respondWithError(r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	//* Sending OTP
	go func() {
		if sendEmailErr := model.SendOTP(user.Email, user.OTP); sendEmailErr != nil {
			respondWithError(r, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}()

	_, insertDBErr := queries.CreateUser(ctx, db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Otp:      user.OTP,
	})

	//* Checking for errors while inserting in the DB
	if insertDBErr != nil {

		if strings.HasPrefix(insertDBErr.Error(), "ERROR: duplicate key") {
			respondWithError(r, http.StatusConflict, "User already exists")
			return
		} else if strings.Contains(insertDBErr.Error(), "\"valid_email\"") {
			respondWithError(r, http.StatusBadRequest, "Invalid Email")
			return
		}

		respondWithError(r, http.StatusInternalServerError, "Error in inserting the document")
		return
	}

	r.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "OTP has been sent to your email"})
}

func ValidateOTP(r *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var req model.OTP
	//* Checking for invalid json format
	if err := r.BindJSON(&req); err != nil {
		respondWithError(r, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	//* Validating if all the fields are present
	if validationErr := validate.Struct(&req); validationErr != nil {
		respondWithError(r, http.StatusBadRequest, "Please provide the required credentials.")
		return
	}

	queries := db.New(configs.CONN)

	//* Checking whether user exists or not
	user, getUserErr := queries.GetUserByEmail(ctx, req.Email)
	if getUserErr != nil {
		respondWithError(r, http.StatusBadRequest, "User does not exist. Please register to generate OTP.")
		return
	}

	//* Checking if user is already verified
	if user.Isverified {
		respondWithError(r, http.StatusOK, "User already verified. Please login.")
		return
	}

	//* Validating OTP
	if user.Otp != req.OTP {
		respondWithError(r, http.StatusUnauthorized, "Invalid OTP")
		return
	}

	//* Updating user to be verified
	updateUserErr := queries.UpdateUser(ctx, req.Email)
	if updateUserErr != nil {
		fmt.Println(updateUserErr)
		respondWithError(r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	//* Generating Token
	token, tokenErr := auth.GenerateJWT(req.Email)
	if tokenErr != nil {
		respondWithError(r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	r.JSON(http.StatusAccepted, responses.UserResponse{Status: http.StatusAccepted, Message: "success", Data: map[string]interface{}{"token": token}})

}

func respondWithError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, responses.UserResponse{
		Status:  statusCode,
		Message: message,
	})
}
