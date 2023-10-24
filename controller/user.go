package controller

import (
	"Gin/Basics/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type errMsg struct {
	Message string `json:"message,omitempty"`
}

// ! Always keep export entities capitalized
func Login(context *gin.Context) {
	var user model.User
	if err := context.BindJSON(&user); err != nil {
		var message = errMsg{Message: "Invalid JSON data"}
		context.JSON(http.StatusBadRequest, message)
		return
	}

	if user.Email == "" || user.Password == "" {
		var message = errMsg{Message: "Please provide both email and password"}
		context.JSON(http.StatusBadRequest, message)
		return
	}

	context.JSON(http.StatusCreated, user)
}

func Register(context *gin.Context) {
	var message errMsg = errMsg{Message: "REGISTER"}
	context.IndentedJSON(http.StatusAccepted, message)
}

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
