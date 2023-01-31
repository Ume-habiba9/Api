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
	role, _ := c.Get("role")
	fmt.Println(role)
	if role == 1 {
		cars, err := db.GetCarsfromDB(userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, cars)
	}
	if role == 2 {
		cars, err := db.GetCarsbyAdmin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, cars)
	}
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
	role, _ := c.Get("role")
	if role == 1 {
		car, err := db.GetcarfromDB(id, userid.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, car)
	}
	if role == 2 {
		car, err := db.GetcarbyAdmin(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, car)
	}
}

func DeleteCar(c *gin.Context) {
	i := c.Param("id")
	userid, _ := c.Get("userid")
	role, _ := c.Get("role")
	if role == 1 {
		err := db.DeletecarfromDB(i, userid.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// c.JSON(http.StatusOK, gin.H{"message": "Car deleted!!"})
	}
	if role == 2 {
		err := db.DeleteCarbyAdmin(i)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// c.JSON(http.StatusOK, gin.H{"message": "Car deleted!!"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted!!"})

}

func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var cardata db.Car
	cardata.ID = id
	userID, _ := c.Get("userid")
	role, _ := c.Get("role")
	cardata.UserID = userID.(string)
	err := c.ShouldBindJSON(&cardata)
	c.JSON(http.StatusAccepted, "Updated Successfully")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if role == 1 {
		err := db.UpdatecarinDB(id, userID.(string), db.Car(cardata))
		// c.JSON(http.StatusAccepted, "Updated Successfully")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
		}
	}
	if role == 2 {
		err := db.UpdateCarbyAdmin(id, db.Car(cardata))
		// c.JSON(http.StatusAccepted, "Updated Successfully")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
		}
	}
}
