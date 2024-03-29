package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "github.com/rhitambanerjee/GOlang-JWT-Project/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "token required"})
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("Email", claims.Email)
		c.Set("FirstName", claims.FirstName)
		c.Set("LastName", claims.LastName)
		c.Set("UserId", claims.UserId)
		c.Set("UserType", claims.UserType)

	}
}