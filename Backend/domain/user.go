package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Nombre        string         `json:"nombre" gorm:"not null"`
	Email         string         `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash  string         `json:"-" gorm:"not null"`
	TipoUsuario   bool           `json:"tipo_usuario" gorm:"not null"`
	FechaCreacion time.Time      `json:"fecha_creacion" gorm:"autoCreateTime"`
	Activo        bool           `json:"activo" gorm:"default:true"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`

	Inscripciones []Inscription `json:"inscripciones,omitempty" gorm:"foreignKey:UsuarioID"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

func (u *User) GetRole() string {
	if u.TipoUsuario {
		return "administrador"
	}
	return "socio"
}
