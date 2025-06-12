package controllers

import (
	"net/http"
	"strconv"

	"gym-backend/domain"
	"gym-backend/services"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	activityService *services.ActivityService
}

func NewActivityController() *ActivityController {
	return &ActivityController{
		activityService: services.NewActivityService(),
	}
}

func (ctrl *ActivityController) GetAllActivities(c *gin.Context) {
	activities, err := ctrl.activityService.GetAllActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo actividades"})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (ctrl *ActivityController) GetActivityByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	activity, err := ctrl.activityService.GetActivityByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func (ctrl *ActivityController) CreateActivity(c *gin.Context) {
	var req domain.CreateActivityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	activityID, err := ctrl.activityService.CreateActivity(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Actividad creada con éxito",
		"id":      activityID,
	})
}

func (ctrl *ActivityController) UpdateActivity(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req domain.UpdateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err = ctrl.activityService.UpdateActivity(uint(id), req)
	if err != nil {
		if err.Error() == "actividad no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Actividad actualizada"})
}

func (ctrl *ActivityController) DeleteActivity(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.activityService.DeleteActivity(uint(id))
	if err != nil {
		if err.Error() == "actividad no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error eliminando actividad"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Actividad eliminada"})
}
