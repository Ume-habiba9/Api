package main

import (
	"github.com/Ume-habiba9/Api/api"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.GET("/Cars", api.GetallCars)
	router.POST("/Cars", api.PostCar)
	router.GET("/Cars/:id", api.GetCar)
	router.DELETE("/Cars/:id", api.DeleteCar)
	router.PUT("/Cars/:id", api.UpdateCar)
	router.Run("localhost:8080")
}
