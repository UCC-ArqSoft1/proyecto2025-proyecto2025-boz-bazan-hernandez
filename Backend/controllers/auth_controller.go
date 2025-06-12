package controllers

import (
    "net/http"
    
    "gym-backend/domain"
    "gym-backend/services"
    "github.com/gin-gonic/gin"
)

type AuthController struct {
    authService *services.AuthService
}

func NewAuthController() *AuthController {
    return &AuthController{
        authService: services.NewAuthService(),
    }
}

func (ctrl *AuthController) Login(c *gin.Context) {
    var req domain.LoginRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }
    
    response, err := ctrl.authService.Login(req)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
        return
    }
    
    c.JSON(http.StatusOK, response)
}
