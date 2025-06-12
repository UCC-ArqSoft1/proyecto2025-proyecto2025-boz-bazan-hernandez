package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización requerido"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		claims, err := ValidateToken(tokenParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("tipo_usuario", claims.TipoUsuario)

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tipoUsuario, exists := c.Get("tipo_usuario")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Información de usuario no encontrada"})
			c.Abort()
			return
		}

		isAdmin, ok := tipoUsuario.(bool)
		if !ok || !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado - Se requieren permisos de administrador"})
			c.Abort()
			return
		}

		c.Next()
	}
}
