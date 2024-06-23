package middlewares

import (
	"net/http"
	"strings"

	"github.com/oscar503sv/gin-jwt/services"
	"github.com/oscar503sv/gin-jwt/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		// Asegúrate de que el token tenga el formato "Bearer <token>"
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// Verificar si el token está en la lista negra
		isBlacklisted, err := services.IsBlacklisted(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking token blacklist"})
			c.Abort()
			return
		}

		if isBlacklisted {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is blacklisted"})
			c.Abort()
			return
		}

		// Validar el token
		_, err = utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
