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
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshclaims := refreshToken.Claims.(jwt.MapClaims)
	refreshclaims["refreshExpirationTime"] = expirationTime.Add(148 * time.Hour).Unix()
	refreshclaims["email"] = claims.Email
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
	if claims, ok := token.Claims.(*customClaim); ok && token.Valid {
		fmt.Println("Signature is Valid")
		if claims.ExpiresAt < time.Now().Local().Unix() {
			return err
		}
		refreshToken, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtkey), nil
		})
		if err != nil {
			return err
		}
		if refreshclaims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
			if (refreshclaims["email"]) != claims.Email {
				return err
			}
		}

		// fmt.Println("token", refreshToken)
	}
	return nil
}
