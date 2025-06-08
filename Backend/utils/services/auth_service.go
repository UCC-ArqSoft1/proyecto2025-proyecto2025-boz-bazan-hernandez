package services

import (
	"errors"
	"os"
	"time"

	"gym-management-system/config"
	"gym-management-system/models"
	"gym-management-system/utils"

	"github.com/golang-jwt/jwt/v4"
)

// AuthService maneja la lógica de autenticación
type AuthService struct{}

// NewAuthService crea una nueva instancia del servicio de autenticación
func NewAuthService() *AuthService {
	return &AuthService{}
}

// Login autentica un usuario y genera un token JWT
func (s *AuthService) Login(email, password string) (*models.LoginResponse, error) {
	db := config.GetDB()

	// Buscar usuario por email
	var user models.User
	if err := db.Where("email = ? AND activo = ?", email, true).First(&user).Error; err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	// Verificar contraseña
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("credenciales inválidas")
	}

	// Generar token JWT
	token, err := s.GenerateJWT(user.ID, user.Email, user.TipoUsuario)
	if err != nil {
		return nil, errors.New("error al generar token")
	}

	return &models.LoginResponse{
		Token:   token,
		User:    user.ToResponse(),
		Message: "Login exitoso",
	}, nil
}

// Register registra un nuevo usuario
func (s *AuthService) Register(req *models.RegisterRequest) (*models.UserResponse, error) {
	db := config.GetDB()

	// Verificar si el email ya existe
	var existingUser models.User
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("el email ya está registrado")
	}

	// Hashear contraseña
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("error al procesar la contraseña")
	}

	// Crear nuevo usuario
	user := models.User{
		Nombre:       req.Nombre,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		TipoUsuario:  req.TipoUsuario,
		Activo:       true,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, errors.New("error al crear el usuario")
	}

	userResponse := user.ToResponse()
	return &userResponse, nil
}

// GenerateJWT genera un token JWT para un usuario
func (s *AuthService) GenerateJWT(userID uint, email string, tipoUsuario bool) (string, error) {
	// Crear claims
	claims := jwt.MapClaims{
		"user_id":      userID,
		"email":        email,
		"tipo_usuario": tipoUsuario,
		"exp":          time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
		"iat":          time.Now().Unix(),
	}

	// Crear token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar token
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "mi_clave_secreta_super_segura" // Valor por defecto para desarrollo
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken valida un token JWT
func (s *AuthService) ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "mi_clave_secreta_super_segura"
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("token inválido")
}
