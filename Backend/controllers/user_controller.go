<<<<<<< HEAD
package controllers

import (
	"net/http"
	"strconv"

	"gym-backend/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (ctrl *UserController) GetMyActivities(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	activities, err := ctrl.userService.GetUserActivities(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo actividades"})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (ctrl *UserController) EnrollInActivity(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	activityIDStr := c.Param("id")
	activityID, err := strconv.ParseUint(activityIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de actividad inválido"})
		return
	}

	err = ctrl.userService.EnrollInActivity(userID.(uint), uint(activityID))
	if err != nil {
		if err.Error() == "la actividad está llena" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La actividad está llena"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción exitosa"})
}
=======
package controllers

import (
    "net/http"
    "strconv"
    
    "gym-backend/services"
    "github.com/gin-gonic/gin"
)

type UserController struct {
    userService *services.UserService
}

func NewUserController() *UserController {
    return &UserController{
        userService: services.NewUserService(),
    }
}

func (ctrl *UserController) GetMyActivities(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }
    
    activities, err := ctrl.userService.GetUserActivities(userID.(uint))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo actividades"})
        return
    }
    
    c.JSON(http.StatusOK, activities)
}

func (ctrl *UserController) EnrollInActivity(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }
    
    activityIDStr := c.Param("id")
    activityID, err := strconv.ParseUint(activityIDStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de actividad inválido"})
        return
    }
    
    err = ctrl.userService.EnrollInActivity(userID.(uint), uint(activityID))
    if err != nil {
        if err.Error() == "la actividad está llena" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "La actividad está llena"})
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Inscripción exitosa"})
}
>>>>>>> e9f915f0d8d09355f0f2c17b2ed95dc1b1fad0ed
