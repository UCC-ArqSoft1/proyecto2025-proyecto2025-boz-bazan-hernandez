package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	Titulo           string         `json:"titulo" gorm:"not null"`
	Descripcion      string         `json:"descripcion"`
	Categoria        string         `json:"categoria"`
	Instructor       string         `json:"instructor"`
	DiaSemana        string         `json:"dia_semana"`
	Horario          string         `json:"horario"` // Almacenado como string en formato HH:MM
	Duracion         int            `json:"duracion"` // En minutos
	CupoMaximo       int            `json:"cupo_maximo"`
	CupoDisponible   int            `json:"cupo_disponible"`
	Foto             string         `json:"foto"`
	FechaCreacion    time.Time      `json:"fecha_creacion" gorm:"autoCreateTime"`
	Activo           bool           `json:"activo" gorm:"default:true"`
	Inscriptions     []Inscription  `json:"inscriptions,omitempty" gorm:"foreignKey:ActividadID"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}

// ActivityRequest representa la estructura para crear/actualizar actividades
type ActivityRequest struct {
	Titulo         string `json:"titulo" binding:"required,min=3"`
	Descripcion    string `json:"descripcion"`
	Categoria      string `json:"categoria" binding:"required"`
	Instructor     string `json:"instructor" binding:"required"`
	DiaSemana      string `json:"dia_semana" binding:"required"`
	Horario        string `json:"horario" binding:"required"` // Formato HH:MM
	Duracion       int    `json:"duracion" binding:"required,min=15,max=180"`
	CupoMaximo     int    `json:"cupo_maximo" binding:"required,min=1,max=50"`
	Foto           string `json:"foto"`
}

// ToActivity convierte un ActivityRequest a Activity
func (ar *ActivityRequest) ToActivity() Activity {
	return Activity{
		Titulo:         ar.Titulo,
		Descripcion:    ar.Descripcion,
		Categoria:      ar.Categoria,
		Instructor:     ar.Instructor,
		DiaSemana:      ar.DiaSemana,
		Horario:        ar.Horario,
		Duracion:       ar.Duracion,
		CupoMaximo:     ar.CupoMaximo,
		CupoDisponible: ar.CupoMaximo, // Inicialmente disponible = máximo
		Foto:           ar.Foto,
		Activo:         true,
	}
}

// ActivityResponse representa la respuesta con información adicional
type ActivityResponse struct {
	Activity
	Inscripciones int  `json:"inscripciones"` // Cantidad de inscripciones
	PuedoInscribirme bool `json:"puedo_inscribirme"`
}

// SearchParams representa los parámetros de búsqueda
type SearchParams struct {
	Query     string `form:"q"`
	Categoria string `form:"categoria"`
	DiaSemana string `form:"dia_semana"`
	Instructor string `form:"instructor"`
	Page      int    `form:"page"`
	Limit     int    `form:"limit"`
}
