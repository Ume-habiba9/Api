package main

import (
	"github.com/Ume-habiba9/Api/api"
	"github.com/Ume-habiba9/Api/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.GET("/Cars", middleware.AuthMiddleware(), api.GetallCars)
	router.GET("/Cars/:id", middleware.AuthMiddleware(), api.GetCar)
	router.PUT("/Users/:id", middleware.AuthMiddleware(), api.UpdateUser)
	router.GET("/Users", middleware.AuthMiddleware(), api.Getallusers)
	router.POST("/Cars", middleware.AuthMiddleware(), api.PostCar)
	router.GET("/Users/:id", middleware.AuthMiddleware(), api.GetUser)
	router.DELETE("/Cars/:id", middleware.AuthMiddleware(), api.DeleteCar)
	router.DELETE("/Users/:id", middleware.AuthMiddleware(), api.DeleteUser)
	router.PUT("/Cars/:id", middleware.AuthMiddleware(), api.UpdateCar)
	router.POST("/signup/admin", api.PostAdmin)
	router.POST("/signup", api.Postuser)
	router.POST("/login", api.LogIn)
	router.POST("/refreshtoken", api.RefreshToken)
	router.Run("localhost:8080")
}
