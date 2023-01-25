package api

import (
	"fmt"
	"net/http"

	"github.com/Ume-habiba9/Api/db"
	"github.com/Ume-habiba9/Api/middleware"
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
		fmt.Println(middleware.GenerateJWT(logInDetails.Email))
		token, err := middleware.GenerateJWT(logInDetails.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Token": token})
		if _, err := fmt.Println(middleware.ValidateToken(token)); err != nil {
			return
		}

	}
}
