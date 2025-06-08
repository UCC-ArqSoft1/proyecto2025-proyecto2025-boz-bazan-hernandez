package models

import (
	"time"

	"gorm.io/gorm"
)

type Inscription struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	UsuarioID        uint           `json:"usuario_id" gorm:"not null"`
	ActividadID      uint           `json:"actividad_id" gorm:"not null"`
	FechaInscripcion time.Time      `json:"fecha_inscripcion" gorm:"autoCreateTime"`
	User             User           `json:"user,omitempty" gorm:"foreignKey:UsuarioID"`
	Activity         Activity       `json:"activity,omitempty" gorm:"foreignKey:ActividadID"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}

// InscriptionRequest representa la solicitud para inscribirse a una actividad
type InscriptionRequest struct {
	ActividadID uint `json:"actividad_id" binding:"required"`
}

// InscriptionResponse representa la respuesta con información de la inscripción
type InscriptionResponse struct {
	ID               uint      `json:"id"`
	UsuarioID        uint      `json:"usuario_id"`
	ActividadID      uint      `json:"actividad_id"`
	FechaInscripcion time.Time `json:"fecha_inscripcion"`
	Activity         Activity  `json:"activity"`
}

// ToResponse convierte una Inscription a InscriptionResponse
func (i *Inscription) ToResponse() InscriptionResponse {
	return InscriptionResponse{
		ID:               i.ID,
		UsuarioID:        i.UsuarioID,
		ActividadID:      i.ActividadID,
		FechaInscripcion: i.FechaInscripcion,
		Activity:         i.Activity,
	}
}
