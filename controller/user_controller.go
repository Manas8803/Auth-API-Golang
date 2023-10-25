package controller

import (
	"Gin/Basics/auth"
	"Gin/Basics/configs"
	db "Gin/Basics/db/sqlconfig"
	"Gin/Basics/model"
	"Gin/Basics/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func Login(r *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var req model.User

	//* Checking for invalid json format
	if err := r.BindJSON(&req); err != nil {
		respondWithError(r, http.StatusBadRequest, "Invalid JSON data", err)
		return
	}

	//* Validating the presence of all the required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		respondWithError(r, http.StatusBadRequest, "Please provide with sufficient credentials", validationErr)
		return
	}

	queries := db.New(configs.CONN)
	//* Checking whether the user is registered
	user, err := queries.GetUser(ctx, req.Email)
	if err != nil {
		respondWithError(r, http.StatusBadGateway, "Internal Server Error", err)
		return
	}

	//* Verifying password
	credentialsError := model.CheckPassword(req.Password, user.Password)
	if credentialsError != nil {
		respondWithError(r, http.StatusUnauthorized, "Invalid Credentials", credentialsError)
		return
	}

	//* Generating Token
	token, err := auth.GenerateJWT(user.Email)
	if err != nil {
		respondWithError(r, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	r.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusAccepted, Message: "success", Data: map[string]interface{}{"token": token}})
}

func Register(r *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	var user model.User
	defer cancel()

	//* Checking for invalid json format
	if err := r.BindJSON(&user); err != nil {
		respondWithError(r, http.StatusBadRequest, "Invalid JSON data", err)
		return
	}

	//* Validating the presence of all the required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		respondWithError(r, http.StatusBadRequest, "Please provide with sufficient credentials", validationErr)
		return
	}

	//* Hashing Password
	if err := user.HashPassword(user.Password); err != nil {
		respondWithError(r, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	queries := db.New(configs.CONN)

	_, err := queries.CreateUser(ctx, db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	//* Checking for errors in inserting in the DB
	if err != nil {
		respondWithError(r, http.StatusInternalServerError, "Error in inserting the document", err)
		return
	}

	r.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": user}})
}

func respondWithError(ctx *gin.Context, statusCode int, message string, err error) {
	ctx.JSON(statusCode, responses.UserResponse{
		Status:  statusCode,
		Message: message,
		Data: map[string]interface{}{
			"data": err.Error(),
		},
	})
}
