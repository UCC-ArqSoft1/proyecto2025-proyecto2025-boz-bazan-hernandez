package services

import (
	"errors"
	"strings"

	"gym-management-system/config"
	"gym-management-system/models"

	"gorm.io/gorm"
)

// ActivityService maneja la lógica de actividades
type ActivityService struct{}

// NewActivityService crea una nueva instancia del servicio de actividades
func NewActivityService() *ActivityService {
	return &ActivityService{}
}

// GetAllActivities obtiene todas las actividades activas
func (s *ActivityService) GetAllActivities() ([]models.Activity, error) {
	db := config.GetDB()
	var activities []models.Activity

	if err := db.Where("activo = ?", true).Find(&activities).Error; err != nil {
		return nil, err
	}

	return activities, nil
}

// GetActivityByID obtiene una actividad por su ID
func (s *ActivityService) GetActivityByID(id uint) (*models.Activity, error) {
	db := config.GetDB()
	var activity models.Activity

	if err := db.Where("id = ? AND activo = ?", id, true).First(&activity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("actividad no encontrada")
		}
		return nil, err
	}

	return &activity, nil
}

// SearchActivities busca actividades según parámetros
func (s *ActivityService) SearchActivities(params models.SearchParams) ([]models.Activity, error) {
	db := config.GetDB()
	var activities []models.Activity

	query := db.Where("activo = ?", true)

	// Búsqueda por texto general
	if params.Query != "" {
		searchTerm := "%" + strings.ToLower(params.Query) + "%"
		query = query.Where(
			"LOWER(titulo) LIKE ? OR LOWER(descripcion) LIKE ? OR LOWER(instructor) LIKE ? OR LOWER(categoria) LIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm,
		)
	}

	// Filtro por categoría
	if params.Categoria != "" {
		query = query.Where("LOWER(categoria) = ?", strings.ToLower(params.Categoria))
	}

	// Filtro por día de semana
	if params.DiaSemana != "" {
		query = query.Where("LOWER(dia_semana) = ?", strings.ToLower(params.DiaSemana))
	}

	// Filtro por instructor
	if params.Instructor != "" {
		query = query.Where("LOWER(instructor) LIKE ?", "%"+strings.ToLower(params.Instructor)+"%")
	}

	// Paginación
	if params.Limit <= 0 {
		params.Limit = 10
	}
	if params.Page <= 0 {
		params.Page = 1
	}
	offset := (params.Page - 1) * params.Limit

	if err := query.Limit(params.Limit).Offset(offset).Find(&activities).Error; err != nil {
		return nil, err
	}

	return activities, nil
}

// CreateActivity crea una nueva actividad
func (s *ActivityService) CreateActivity(req *models.ActivityRequest) (*models.Activity, error) {
	db := config.GetDB()

	activity := req.ToActivity()

	if err := db.Create(&activity).Error; err != nil {
		return nil, errors.New("error al crear la actividad")
	}

	return &activity, nil
}

// UpdateActivity actualiza una actividad existente
func (s *ActivityService) UpdateActivity(id uint, req *models.ActivityRequest) (*models.Activity, error) {
	db := config.GetDB()

	var activity models.Activity
	if err := db.Where("id = ? AND activo = ?", id, true).First(&activity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("actividad no encontrada")
		}
		return nil, err
	}

	// Actualizar campos
	activity.Titulo = req.Titulo
	activity.Descripcion = req.Descripcion
	activity.Categoria = req.Categoria
	activity.Instructor = req.Instructor
	activity.DiaSemana = req.DiaSemana
	activity.Horario = req.Horario
	activity.Duracion = req.Duracion
	activity.Foto = req.Foto

	// Si se cambia el cupo máximo, ajustar el disponible
	if req.CupoMaximo != activity.CupoMaximo {
		inscripciones := s.countInscriptions(id)
		activity.CupoMaximo = req.CupoMaximo
		activity.CupoDisponible = req.CupoMaximo - inscripciones
		if activity.CupoDisponible < 0 {
			activity.CupoDisponible = 0
		}
	}

	if err := db.Save(&activity).Error; err != nil {
		return nil, errors.New("error al actualizar la actividad")
	}

	return &activity, nil
}

// DeleteActivity elimina una actividad (soft delete)
func (s *ActivityService) DeleteActivity(id uint) error {
	db := config.GetDB()

	var activity models.Activity
	if err := db.Where("id = ? AND activo = ?", id, true).First(&activity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("actividad no encontrada")
		}
		return err
	}

	// Marcar como inactiva
	activity.Activo = false
	if err := db.Save(&activity).Error; err != nil {
		return errors.New("error al eliminar la actividad")
	}

	return nil
}

// countInscriptions cuenta las inscripciones activas de una actividad
func (s *ActivityService) countInscriptions(activityID uint) int {
	db := config.GetDB()
	var count int64

	db.Model(&models.Inscription{}).Where("actividad_id = ?", activityID).Count(&count)
	return int(count)
}

// GetActivityWithInscriptionInfo obtiene una actividad con información de inscripciones
func (s *ActivityService) GetActivityWithInscriptionInfo(id uint, userID uint) (*models.ActivityResponse, error) {
	activity, err := s.GetActivityByID(id)
	if err != nil {
		return nil, err
	}

	inscripciones := s.countInscriptions(id)
	puedoInscribirme := s.canUserEnroll(userID, id)

	response := models.ActivityResponse{
		Activity:         *activity,
		Inscripciones:    inscripciones,
		PuedoInscribirme: puedoInscribirme && activity.CupoDisponible > 0,
	}

	return &response, nil
}

// canUserEnroll verifica si un usuario puede inscribirse en una actividad
func (s *ActivityService) canUserEnroll(userID, activityID uint) bool {
	db := config.GetDB()
	var count int64

	db.Model(&models.Inscription{}).
		Where("usuario_id = ? AND actividad_id = ?", userID, activityID).
		Count(&count)

	return count == 0 // Puede inscribirse si no está ya inscripto
}
