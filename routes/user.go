package routes

import (
	"Gin/Basics/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/auth/login", controller.Login)
	router.POST("/auth/register", controller.Register)
}
