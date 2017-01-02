package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mefellows/follow/config"
	"github.com/mefellows/follow/controllers"
)

func main() {
	c := config.NewConfig()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Perform migrations
	healthController := controllers.NewHealthController(c)
	locationController := controllers.NewLocationController(c)
	userController := controllers.NewUserController(c)
	uiController := controllers.NewUIController(c)

	// UI
	r.GET("/map/:id", uiController.Get)

	// Routes
	r.GET("/health", healthController.Get)

	// Shopping list routes
	r.GET("/locations/:id", locationController.Get)
	r.POST("/locations", locationController.Create)

	// Users
	r.GET("/users/:id", userController.GetSingle)
	r.GET("/users", userController.Get)

	r.Run()
}
