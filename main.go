package main

import (
	"Gin/Basics/configs"
	docs "Gin/Basics/docs"
	"Gin/Basics/routes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Registration API
//	@version		1.0
//	@description	This is a registration api for an application.

//	@BasePath	/api/

func main() {

	// prod := configs.RELEASE_MODE()
	// if prod == "true" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	api := router.Group("/api/v1")
	//* Passing the router to all user(auth) routes.
	routes.UserRoute(api)

	//* Connecting to DB
	configs.ConnectDB()

	router.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")
}
