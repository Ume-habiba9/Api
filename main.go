package main

import (
	"github.com/Ume-habiba9/Api/api"
	"github.com/Ume-habiba9/Api/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.GET("/Cars", api.GetallCars)
	router.POST("/Users", api.Postuser)
	router.GET("/Cars/:id", api.GetCar)
	router.PUT("/Users/:id", api.UpdateUser)
	router.GET("/Users", api.Getallusers)
	router.POST("/Cars", api.PostCar)
	router.GET("/Users/:id", api.GetUser)
	router.DELETE("/Cars/:id", api.DeleteCar)
	router.DELETE("/Users/:id", api.DeleteUser)
	router.PUT("/Cars/:id", api.UpdateCar)
	router.Use(middleware.AuthMiddleware())
	router.POST("/Users/:id/:email/:passward", api.LogIn)
	router.Run("localhost:8080")
}
