package middleware

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Secret")
		if header != "CapregSoft" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
