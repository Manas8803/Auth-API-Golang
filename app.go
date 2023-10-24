package main

import (
	"Gin/Basics/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//* Passing the router to all user(auth) routes.
	routes.UserRoute(router)
	router.Run("localhost:9000")
}
