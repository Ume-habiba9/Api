package api

import (
	"fmt"
	"net/http"

	"github.com/Ume-habiba9/Api/db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	
)
func LogInCheck(c *gin.Context) {
id := c.Param("id")
email := c.Param("email")
passward := c.Param("passward")
var logindetails db.User
if err := c.ShouldBindJSON(&logindetails); err != nil {
	c.JSON(http.StatusInternalServerError, "Invalid User Data")
	return
}
_, err := db.LogIn(id, email, passward)
if err != nil {
	c.JSON(http.StatusInternalServerError, err)
	return
}
fmt.Println(logindetails)
if (logindetails.Email != email) || (logindetails.Password != passward) {
	c.JSON(http.StatusOK, "Invalid Data")
} else {
	c.JSON(http.StatusOK, "User Exists")
}
}