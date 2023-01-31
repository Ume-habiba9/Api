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
	userID, _ := c.Get("userid")
	cars, err := db.GetCarsfromDB(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
	userID, _ := c.Get("userid")
	newCar.UserID = userID.(string)
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
	userid, _ := c.Get("userid")
	fmt.Println("user in get cars", userid)
	car, err := db.GetcarfromDB(id, userid.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, car)
}

func DeleteCar(c *gin.Context) {
	i := c.Param("id")
	userid, _ := c.Get("userid")
	err := db.DeletecarfromDB(i, userid.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted!!"})
}

func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var cardata db.Car
	cardata.ID = id
	userID, _ := c.Get("userid")
	cardata.UserID = userID.(string)
	if err := c.ShouldBindJSON(&cardata); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.IndentedJSON(http.StatusOK, "Updated Successfully")
	}
	err := db.UpdatecarinDB(id, userID.(string), db.Car(cardata))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}
}
