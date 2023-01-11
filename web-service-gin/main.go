package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// album represents data about a record album.

type car struct {
	ID       string  `json:"id"`
	CarName  string  `json:"carname"`
	CarType  string  `json:"cartype"`
	Capacity string  `json:"capacity"`
	Price    float64 `json:"price"`
}

// albums slice to seed record album data.

var cars = []car{
	{ID: "1", CarName: "Nissan GT - R", CarType: "Sport", Capacity: "2 People", Price: 56.99},
	{ID: "2", CarName: "CR  - V", CarType: "SUV", Capacity: "4 People", Price: 17.99},
	{ID: "3", CarName: "MG ZX Excite", CarType: "Hatchback", Capacity: "4 people", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}
func postCars(c *gin.Context) {
	var newCar car
	newCar.ID = CreateUUID()
	if err := c.BindJSON(&newCar); err != nil {
		return
	}
	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}
func getaCar(c *gin.Context) {
	id := c.Param("id")
	for _, getcar := range cars {
		if getcar.ID == id {
			c.IndentedJSON(http.StatusOK, getcar)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}
func deleteCars(c *gin.Context) {
	id := c.Param("id")
	for _, deletecar := range cars {
		if deletecar.ID == id {
			c.MultipartForm()
		}
		c.JSON(http.StatusOK, gin.H{"message": "Car not found"})
	}
}
func CreateUUID() string {
	id := uuid.New()
	return id.String()
}
func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.POST("/cars", postCars)
	router.GET("/cars/id", getaCar)
	router.DELETE("/cars/id", deleteCars)
	router.Run("localhost:8080")
}
