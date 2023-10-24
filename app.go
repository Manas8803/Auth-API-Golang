package main

import (
	"Gin/Basics/configs"
	"Gin/Basics/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//* Passing the router to all user(auth) routes.
	routes.UserRoute(router)

	//* Connecting to DB
	configs.ConnectDB()
	router.Run("localhost:9000")
}
