package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Car struct {
	ID       string  `json:"id"`
	CarName  string  `json:"carname"`
	CarType  string  `json:"cartype"`
	Capacity string  `json:"capacity"`
	Price    float64 `json:"price"`
}

var Cars map[uuid.UUID]Car

func init() {
	Cars = make(map[uuid.UUID]Car)
	id1 := uuid.NewV4()
	Cars[id1] = Car{id1.String(), "Tesla", "SUV", "6 People", 6.78}
	id2 := uuid.NewV4()
	Cars[id2] = Car{id2.String(), "BMW", "MPV", "4 people", 63.89}
	id3 := uuid.NewV4()
	Cars[id3] = Car{id3.String(), "Mercedes", "Hatchback", "6 people", 54.9}
}

// get all Cars
func getallCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Cars)
}

// add new Cars
func postCar(c *gin.Context) {
	var newCar Car
	id := uuid.NewV4()
	newCar.ID = id.String()
	if err := c.BindJSON(&newCar); err != nil {
		return
	}
	Cars[id] = newCar
	c.IndentedJSON(http.StatusCreated, newCar)
}

func getCar(c *gin.Context) {
	id := c.Param("id")
	for i, car := range Cars {
		if car.ID == id {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}

func deleteCar(c *gin.Context) {
	i := c.Param("id")
	for car, deletecar := range Cars {
		if deletecar.ID == i {
			delete(Cars, car)
			c.JSON(http.StatusOK, gin.H{"message": "Car deleted!!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Car not found"})
	}
}
func main() {
	router := gin.Default()
	router.GET("/Cars", getallCars)
	router.POST("/Cars", postCar)
	router.GET("/Cars/:id", getCar)
	router.DELETE("/Cars/:id", deleteCar)
	router.Run("localhost:8080")
}
