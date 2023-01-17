package main

import (
	"github.com/gin-gonic/gin"
	"jwt-gin/controller"
	"jwt-gin/models"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)
	r.Run(":8080")
}
