package api

import (
	"net/http"

	"github.com/Ume-habiba9/Api/db"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func PostAdmin(c *gin.Context) {
	var newUser db.User
	database := db.DBConnect()
	defer database.Close()
	id := uuid.NewV4()
	newUser.UserID = id.String()
	newUser.Role = 2
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	err := db.PostAdmininDB(db.User(newUser))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newUser)
}
