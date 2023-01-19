package api

import (
	"fmt"
	"net/http"

	"github.com/Ume-habiba9/Api/db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

func GetallCars(c *gin.Context) {
	database := db.DBConnect()
	defer database.Close()
	cars, err := db.GetCarsfromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, cars)
}

func PostCar(c *gin.Context) {
	var newCar db.Car
	database := db.DBConnect()
	defer database.Close()
	id := uuid.NewV4()
	newCar.ID = id.String()
	if err := c.BindJSON(&newCar); err != nil {
		return
	}
	err := db.PostcarinDB(db.Car(newCar))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newCar)
}

func GetCar(c *gin.Context) {
	id := c.Param("id")
	car, err := db.GetcarfromDB(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, car)
}

func DeleteCar(c *gin.Context) {
	i := c.Param("id")
	err := db.DeletecarfromDB(i)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted!!"})
}

func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var cardata db.Car
	if err := c.ShouldBindJSON(&cardata); err != nil {
		fmt.Println("errr ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err := db.UpdatecarinDB(id, db.Car(cardata))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, "Updated Successfully")
}
