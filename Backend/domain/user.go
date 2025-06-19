package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Nombre        string         `json:"nombre" gorm:"type:varchar(255);not null"`
	Email         string         `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash  string         `json:"-" gorm:"type:varchar(255);not null"`
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
	Token string   `json:"token"`
	Role  string   `json:"role"`
	User  UserInfo `json:"user"`
}

type RegisterRequest struct {
	Nombre      string `json:"nombre" binding:"required,min=2,max=255"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	TipoUsuario bool   `json:"tipo_usuario"`
}

type RegisterResponse struct {
	Message string   `json:"message"`
	User    UserInfo `json:"user"`
}

type UserInfo struct {
	ID     uint   `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

func (u *User) GetRole() string {
	if u.TipoUsuario {
		return "administrador"
	}
	return "socio"
}

func (u *User) ToUserInfo() UserInfo {
	return UserInfo{
		ID:     u.ID,
		Nombre: u.Nombre,
		Email:  u.Email,
		Role:   u.GetRole(),
	}
}

func (User) TableName() string {
	return "users"
}
