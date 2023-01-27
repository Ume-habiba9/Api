package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var jwtkey = []byte("CapregSoft")

type customClaim struct {
	Email string `json:"email"`
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
func GenerateJWT(email string) (string, string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &customClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	refreshExpirationTime := expirationTime.Add(148 * time.Hour)
	refreshClaims := &customClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtkey)
	if err != nil {
		return "", "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
}

func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(signedToken, &customClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtkey), nil
		})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(*customClaim); ok && token.Valid {
		fmt.Println("Signature is Valid")
		if err != nil {
			return err
		}
	}
	return nil
}
func RefreshJWT(tokenString string) (string, error) {
	refreshToken, err := jwt.ParseWithClaims(tokenString, &customClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtkey), nil
	})
	if err != nil {
		return " ", err
	}
	if refreshClaims, ok := refreshToken.Claims.(*customClaim); ok && refreshToken.Valid {
		token, _, err := GenerateJWT(refreshClaims.Email)
		fmt.Println("New Token Generated:", token)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	// fmt.Println(refreshToken)
	return "", nil
}
