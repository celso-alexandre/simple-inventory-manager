package middlewares

import (
	"net/http"

	"github.com/celso-alexandre/simple-inventory-manager/server/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		payload, err := utils.VerifyJwtToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("payload", payload)
		c.Next()
	}
}

func RetrieveAuthPayload(c *gin.Context) utils.JwtPayload {
	return c.MustGet("payload").(utils.JwtPayload)
}
