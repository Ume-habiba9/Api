package main

import (
	"github.com/Ume-habiba9/Api/api"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.GET("/Cars", api.GetallCars)
	router.GET("/Users", api.Getallusers)
	router.POST("/Cars", api.PostCar)
	router.POST("/Users", api.Postuser)
	router.GET("/Cars/:id", api.GetCar)
	router.GET("/Users/:id", api.GetUser)
	router.DELETE("/Cars/:id", api.DeleteCar)
	router.DELETE("/Users/:id", api.DeleteUser)
	router.PUT("/Cars/:id", api.UpdateCar)
	router.PUT("/Users/:id", api.UpdateUser)
	router.POST("/Users/:id/:email/:passward", api.LogInCheck)
	router.Run("localhost:8080")
}
