package controllers

import (
	"net/http"

	"gym-management-system/models"
	"gym-management-system/services"

	"github.com/gin-gonic/gin"
)

var authService = services.NewAuthService()

// Login maneja el inicio de sesión
func Login(c *gin.Context) {
	var req models.LoginRequest

	// Validar datos de entrada
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	// Intentar login
	response, err := authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Register maneja el registro de nuevos usuarios
func Register(c *gin.Context) {
	var req models.RegisterRequest

	// Validar datos de entrada
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	// Crear usuario
	user, err := authService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
		"user": user,
	})
}
