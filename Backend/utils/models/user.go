package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Nombre         string         `json:"nombre" gorm:"not null"`
	Email          string         `json:"email" gorm:"unique;not null"`
	PasswordHash   string         `json:"-" gorm:"not null"` // El "-" hace que no se serialice en JSON
	TipoUsuario    bool           `json:"tipo_usuario" gorm:"default:false"` // false = socio, true = admin
	FechaCreacion  time.Time      `json:"fecha_creacion" gorm:"autoCreateTime"`
	Activo         bool           `json:"activo" gorm:"default:true"`
	Inscriptions   []Inscription  `json:"inscriptions,omitempty" gorm:"foreignKey:UsuarioID"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

// UserResponse es la estructura para respuestas que no incluyen información sensible
type UserResponse struct {
	ID            uint      `json:"id"`
	Nombre        string    `json:"nombre"`
	Email         string    `json:"email"`
	TipoUsuario   bool      `json:"tipo_usuario"`
	FechaCreacion time.Time `json:"fecha_creacion"`
	Activo        bool      `json:"activo"`
}

// ToResponse convierte un User a UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:            u.ID,
		Nombre:        u.Nombre,
		Email:         u.Email,
		TipoUsuario:   u.TipoUsuario,
		FechaCreacion: u.FechaCreacion,
		Activo:        u.Activo,
	}
}

// LoginRequest representa la estructura de datos para el login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest representa la estructura de datos para el registro
type RegisterRequest struct {
	Nombre      string `json:"nombre" binding:"required,min=2"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	TipoUsuario bool   `json:"tipo_usuario"`
}

// LoginResponse representa la respuesta del login
type LoginResponse struct {
	Token   string       `json:"token"`
	User    UserResponse `json:"user"`
	Message string       `json:"message"`
}
