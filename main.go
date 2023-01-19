package main

import (
	"fmt"
	"net/http"

	"github.com/Ume-habiba9/Api/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

func main() {
	router := gin.Default()
	router.GET("/Cars", getallCars)
	router.POST("/Cars", postCar)
	router.GET("/Cars/:id", getCar)
	router.DELETE("/Cars/:id", deleteCar)
	router.PUT("/Cars/:id", updateCar)
	router.Run("localhost:8080")
}
