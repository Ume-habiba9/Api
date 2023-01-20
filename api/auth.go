package api

import (
	"net/http"

	"github.com/Ume-habiba9/Api/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func LogIn(c *gin.Context) {
	email := c.Param("email")
	passward := c.Param("passward")
	var logInDetails db.User
	if err := c.ShouldBindJSON(&logInDetails); err != nil {
		c.JSON(http.StatusInternalServerError, "Invalid User Data")
		return
	}
	_, err := db.LogInCheck(email, passward)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if (logInDetails.Email != email) || (logInDetails.Password != passward) {
		c.JSON(http.StatusOK, "Invalid Data")
	} else {
		c.JSON(http.StatusOK, "User Exists")
	}
}
