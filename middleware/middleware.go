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
	UserID string `json:"userid"`
	Email  string `json:"email"`
	Role   int    `json:"role"`
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
		token, err := jwt.ParseWithClaims(tokenString, &customClaim{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtkey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*customClaim); ok && token.Valid {
			userID := claims.UserID
			role := claims.Role
			c.Set("userid", userID)
			c.Set("role", role)
			fmt.Println("role", role)
			fmt.Println("User id is", userID)
		} else {
			c.JSON(http.StatusUnauthorized, "UserID not found!")
			c.Abort()
			return
		}
		c.Next()
	}
}
func GenerateJWT(userid, email string, role int) (string, string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &customClaim{
		UserID: userid,
		Email:  email,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	refreshExpirationTime := expirationTime.Add(148 * time.Hour)
	refreshClaims := &customClaim{
		UserID: userid,
		Email:  email,
		Role:   role,
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

func ReGenerateJWT(tokenString string) (string, error) {
	refreshToken, err := jwt.ParseWithClaims(tokenString, &customClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtkey), nil
	})
	if err != nil {
		return "", err
	}
	if refreshClaims, ok := refreshToken.Claims.(*customClaim); ok && refreshToken.Valid {
		token, _, err := GenerateJWT(refreshClaims.UserID, refreshClaims.Email, refreshClaims.Role)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", nil
}
