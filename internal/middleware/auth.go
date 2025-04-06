package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflibima25/task-management/pkg/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// extrack token from header
		tokenString, err := auth.ExtractTokenFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "message": err.Error()})
			c.Abort()
			return
		}

		// validasi token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "message": "invalid token"})
			c.Abort()
			return
		}

		// set user info ke context untuk digunakan oleh handler
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("name", claims.Name)

		c.Next()
	}
}
