package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var jwtkey = []byte("CapregSoft")

type customClaim struct {
	Email    string `json:"email"`
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Secret")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
			c.Abort()
			return
		}
		err := ValidateToken(tokenString)
		if err != nil {
			return
		}
		c.Next()
	}
}
func GenerateJWT(email string) (tokenString string, err error) {
	expirationTime := time.Now().Add( 1* time.Hour)
	claims := &customClaim{
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtkey)
	return
}
func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(signedToken, &customClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtkey), nil
		})
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*customClaim)
	if !ok {
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return err
	}
	return nil
}
