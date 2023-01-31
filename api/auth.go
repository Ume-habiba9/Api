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
		c.JSON(http.StatusUnauthorized, "please check your email or password.")
		return
	}
	if userdata != nil {
		token, refreshToken, err := middleware.GenerateJWT(userdata.UserID, userdata.Email, userdata.Role)
		fmt.Println()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Token": token, "Refresh Token": refreshToken})
	} else {
		c.JSON(http.StatusOK, "Invalid Data")
	}
}
func RefreshToken(c *gin.Context) {
	tokenString := c.GetHeader("Secret")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		c.Abort()
		return
	}
	token, err := middleware.ReGenerateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
	} else {
		c.JSON(http.StatusAccepted, gin.H{"New Access token": token})
	}
}
