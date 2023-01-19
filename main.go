package main

import (
	"github.com/Ume-habiba9/Api/Modules"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.GET("/Cars", Modules.GetallCars)
	router.POST("/Cars", Modules.PostCar)
	router.GET("/Cars/:id", Modules.GetCar)
	router.DELETE("/Cars/:id", Modules.DeleteCar)
	router.PUT("/Cars/:id", Modules.UpdateCar)
	router.Run("localhost:8080")
}
