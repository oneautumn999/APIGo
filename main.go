package main

import (
	"btpn/controllers"
	"btpn/initializers"
	"btpn/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/users/register", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.PUT("users/edit/:id", middleware.RequireAuth, controllers.UserUpdate)
	r.DELETE("users/delete/:id", middleware.RequireAuth, controllers.UserDelete)
	r.GET("/users/login", middleware.RequireAuth, controllers.Validate)

	r.POST("/users/photo", middleware.RequireAuth, controllers.Photostambah)
	r.DELETE("/users/photo/delete/:id", middleware.RequireAuth, controllers.Photosdelete)
	r.PUT("/users/photo/update/:id", middleware.RequireAuth, controllers.PhotoUpdate)
	r.Run()
}
