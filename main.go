package main

import (
	authcontroler "hello-world/controllers/auth"
	productcontroller "hello-world/controllers/product"
	"hello-world/models"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	models.ConnectDatabase()

	api := r.Group("/api")
	api.POST("/login", authcontroler.Login)
	api.POST("/register", authcontroler.Register)
	api.Use(authcontroler.AuthMiddleware())
	api.GET("/products", productcontroller.Index)
	api.GET("/products/:id", productcontroller.Show)
	api.POST("/products", productcontroller.Create)
	api.PUT("/products/:id", productcontroller.Update)
	api.DELETE("/products/:id", productcontroller.Delete)

	


	r.Run(":8411")
}
