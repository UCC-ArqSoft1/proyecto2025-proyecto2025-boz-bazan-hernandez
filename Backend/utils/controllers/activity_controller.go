package controllers

import (
	"net/http"
	"strconv"

	"gym-management-system/models"
	"gym-management-system/services"

	"github.com/gin-gonic/gin"
)

var activityService = services.NewActivityService()

// GetActivities obtiene todas las actividades
func GetActivities(c *gin.Context) {
	activities, err := activityService.GetAllActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener actividades",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"activities": activities,
		"total": len(activities),
	})
}

// GetActivityByID obtiene una actividad por ID
func GetActivityByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Si hay usuario autenticado, obtener información adicional
	userID, exists := c.Get("user_id")
	if exists {
		activityResponse, err := activityService.GetActivityWithInscriptionInfo(uint(id), userID.(uint))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, activityResponse)
		return
	}

	// Si no hay usuario autenticado, solo devolver la actividad básica
	activity, err := activityService.GetActivityByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// SearchActivities busca actividades
func SearchActivities(c *gin.Context) {
	var params models.SearchParams

	// Bind query parameters
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parámetros de búsqueda inválidos",
		})
		return
	}

	activities, err := activityService.SearchActivities(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al buscar actividades",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"activities": activities,
		"total": len(activities),
		"page": params.Page,
		"limit": params.Limit,
	})
}

// CreateActivity crea una nueva actividad (solo administradores)
func CreateActivity(c *gin.Context) {
	var req models.ActivityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	activity, err := activityService.CreateActivity(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Actividad creada exitosamente",
		"activity": activity,
	})
}

// UpdateActivity actualiza una actividad (solo administradores)
func UpdateActivity(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req models.ActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	activity, err := activityService.UpdateActivity(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Actividad actualizada exitosamente",
		"activity": activity,
	})
}

// DeleteActivity elimina una actividad (solo administradores)
func DeleteActivity(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = activityService.DeleteActivity(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Actividad eliminada exitosamente",
	})
}
