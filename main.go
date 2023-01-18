package main

import (
	"github.com/gin-gonic/gin"
	"jwt-gin/controller"
	"jwt-gin/middlewares"
	"jwt-gin/models"
)

func main() {

	models.ConnectDataBase()
	r := gin.Default()

	public := r.Group("/api")
	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controller.CurrentUser)

	err := r.Run(":8080")
	if err != nil {
		return
	}

}
