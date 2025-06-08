package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Claims representa las claims del JWT
type Claims struct {
	UserID      uint   `json:"user_id"`
	Email       string `json:"email"`
	TipoUsuario bool   `json:"tipo_usuario"`
	jwt.RegisteredClaims
}

// AuthMiddleware verifica el token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener token del header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token de autorización requerido",
			})
			c.Abort()
			return
		}

		// Verificar formato "Bearer token"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Formato de token inválido",
			})
			c.Abort()
			return
		}

		// Validar token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(getJWTSecret()), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token inválido",
			})
			c.Abort()
			return
		}

		// Agregar claims al contexto
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("tipo_usuario", claims.TipoUsuario)

		c.Next()
	}
}

// AdminMiddleware verifica que el usuario sea administrador
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tipoUsuario, exists := c.Get("tipo_usuario")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Usuario no autenticado",
			})
			c.Abort()
			return
		}

		if !tipoUsuario.(bool) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Acceso denegado: se requieren permisos de administrador",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// getJWTSecret obtiene la clave secreta del JWT
func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "mi_clave_secreta_super_segura" // Valor por defecto para desarrollo
	}
	return secret
}
