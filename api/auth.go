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
	var logInDetails db.User
	if err := c.ShouldBindJSON(&logInDetails); err != nil {
		c.JSON(http.StatusInternalServerError, "Invalid User Data")
		return
	}
	userdata, err := db.LogInCheck(logInDetails.Email, logInDetails.Password)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusUnauthorized, "please check your email or password.")
		return
	}
	if userdata != nil {
		fmt.Println(middleware.GenerateJWT(logInDetails.Email))
		token, refreshToken, err := middleware.GenerateJWT(logInDetails.Email)
		fmt.Println()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Token": token, "Refresh Token": refreshToken})
		if _, err := fmt.Println(middleware.ValidateToken(token)); err != nil {
			return
		}
	} else {
		c.JSON(http.StatusOK, "Invalid Data")
	}
}
func RefreshToken(c *gin.Context) {
	var token string
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhhYmliYUBnbWFpbC5jb20iLCJleHAiOjE2NzUzMzY4NTB9.3q0R4xDJ50JYXl4wjXX6C1DSV3aIEomG53qy7AW2_1E"
	refreshed, err := middleware.RefreshJWT(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)
	} else {
		c.JSON(http.StatusAccepted, gin.H{"Token Created": refreshed})
		fmt.Println(refreshed)
	}
}
