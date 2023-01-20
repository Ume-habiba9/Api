package api

import (
	"fmt"
	"net/http"

	"github.com/Ume-habiba9/Api/db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

func Postuser(c *gin.Context) {
	var newUser db.User
	database := db.DBConnect()
	defer database.Close()
	id := uuid.NewV4()
	newUser.UserID = id.String()
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	err := db.PostUserinDB(db.User(newUser))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newUser)
}
func Getallusers(c *gin.Context) {
	database := db.DBConnect()
	defer database.Close()
	users, err := db.GetUsersfromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := db.GetUserfromDB(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
func DeleteUser(c *gin.Context) {
	i := c.Param("id")
	err := db.DeleteUserfromDB(i)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted!!"})
}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userdata db.User
	if err := c.ShouldBindJSON(&userdata); err != nil {
		fmt.Println("errr ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err := db.UpdateUserinDB(id, db.User(userdata))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, "Updated Successfully")
}


