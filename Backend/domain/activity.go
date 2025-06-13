package domain

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Titulo         string         `json:"titulo" gorm:"not null"`
	Descripcion    string         `json:"descripcion"`
	Categoria      string         `json:"categoria" gorm:"not null"`
	Instructor     string         `json:"instructor" gorm:"not null"`
	DiaSemana      string         `json:"dia" gorm:"not null"`
	Horario        string         `json:"horario" gorm:"type:time;not null"`
	Duracion       int            `json:"duracion" gorm:"not null"`
	CupoMaximo     int            `json:"cupo" gorm:"not null"`
	CupoDisponible int            `json:"cupo_disponible" gorm:"not null"`
	Foto           string         `json:"foto_url" gorm:"column:foto"`
	FechaCreacion  time.Time      `json:"fecha_creacion" gorm:"autoCreateTime"`
	Activo         bool           `json:"activo" gorm:"default:true"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`

	Inscripciones []Inscription `json:"inscripciones,omitempty" gorm:"foreignKey:ActividadID"`
}

type CreateActivityRequest struct {
	Titulo      string `json:"titulo" binding:"required"`
	Descripcion string `json:"descripcion"`
	Categoria   string `json:"categoria" binding:"required"`
	Instructor  string `json:"instructor" binding:"required"`
	Dia         string `json:"dia" binding:"required"`
	Horario     string `json:"horario" binding:"required"`
	Duracion    string `json:"duracion" binding:"required"`
	Cupo        int    `json:"cupo" binding:"required,min=1"`
	FotoURL     string `json:"foto_url"`
}

type UpdateActivityRequest struct {
	Titulo      *string `json:"titulo"`
	Descripcion *string `json:"descripcion"`
	Categoria   *string `json:"categoria"`
	Instructor  *string `json:"instructor"`
	Dia         *string `json:"dia"`
	Horario     *string `json:"horario"`
	Duracion    *string `json:"duracion"`
	Cupo        *int    `json:"cupo"`
	FotoURL     *string `json:"foto_url"`
}

type ActivityListResponse struct {
	ID       uint   `json:"id"`
	Titulo   string `json:"titulo"`
	Horario  string `json:"horario"`
	Profesor string `json:"profesor"`
}

type ActivityDetailResponse struct {
	ID          uint   `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Dia         string `json:"dia"`
	Horario     string `json:"horario"`
	Duracion    string `json:"duracion"`
	Cupo        int    `json:"cupo"`
	Categoria   string `json:"categoria"`
	Instructor  string `json:"instructor"`
	FotoURL     string `json:"foto_url"`
}

type MyActivityResponse struct {
	ID         uint   `json:"id"`
	Titulo     string `json:"titulo"`
	Horario    string `json:"horario"`
	Dia        string `json:"dia"`
	Instructor string `json:"instructor"`
}

func (Activity) TableName() string {
	return "activities"
}
