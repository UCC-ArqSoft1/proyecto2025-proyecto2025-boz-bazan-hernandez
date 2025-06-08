package controllers

import (
	"net/http"
	"strconv"

	"gym-management-system/models"
	"gym-management-system/services"

	"github.com/gin-gonic/gin"
)

var userService = services.NewUserService()

// GetUserProfile obtiene el perfil del usuario autenticado
func GetUserProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Usuario no autenticado",
		})
		return
	}

	user, err := userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user.ToResponse())
}

(NO ESTA TERMINADO TERMINAR***)
