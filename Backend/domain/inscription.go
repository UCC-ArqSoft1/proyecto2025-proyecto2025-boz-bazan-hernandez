package domain

import (
	"time"

	"gorm.io/gorm"
)

type Inscription struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	UsuarioID        uint           `json:"usuario_id" gorm:"not null"`
	ActividadID      uint           `json:"actividad_id" gorm:"not null"`
	FechaInscripcion time.Time      `json:"fecha_inscripcion" gorm:"autoCreateTime"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`

	Usuario   User     `json:"usuario,omitempty" gorm:"foreignKey:UsuarioID"`
	Actividad Activity `json:"actividad,omitempty" gorm:"foreignKey:ActividadID"`
}

func (Inscription) TableName() string {
	return "inscripciones"
}
