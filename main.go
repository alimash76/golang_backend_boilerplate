package main

import (
	"energytoken/controllers"
	"energytoken/models"
	"energytoken/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDatabase()

	r := gin.Default()

	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JWTAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
