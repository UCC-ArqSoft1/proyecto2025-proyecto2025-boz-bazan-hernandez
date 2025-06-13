package dao

import (
	"gym-backend/domain"
	"gym-backend/utils"
)

type ActivityDAO struct{}

func NewActivityDAO() *ActivityDAO {
	return &ActivityDAO{}
}

// Obtener todas las actividades activas
func (d *ActivityDAO) FindAll() ([]domain.Activity, error) {
	var activities []domain.Activity
	err := utils.DB.Where("activo = ?", true).Find(&activities).Error
	return activities, err
}

// Buscar actividad por ID
func (d *ActivityDAO) FindByID(id uint) (*domain.Activity, error) {
	var activity domain.Activity
	err := utils.DB.Where("activo = ?", true).First(&activity, id).Error
	return &activity, err
}

// Crear nueva actividad
func (d *ActivityDAO) Create(activity *domain.Activity) error {
	return utils.DB.Create(activity).Error
}

// Actualizar actividad
func (d *ActivityDAO) Update(activity *domain.Activity) error {
	return utils.DB.Save(activity).Error
}

// Decrementar cupo disponible
func (d *ActivityDAO) DecrementAvailableSlots(activityID uint) error {
	return utils.DB.Model(&domain.Activity{}).
		Where("id = ? AND cupo_disponible > 0", activityID).
		Update("cupo_disponible", utils.DB.Raw("cupo_disponible - 1")).Error
}

// Incrementar cupo disponible (para cancelaciones)
func (d *ActivityDAO) IncrementAvailableSlots(activityID uint) error {
	return utils.DB.Model(&domain.Activity{}).
		Where("id = ?", activityID).
		Update("cupo_disponible", utils.DB.Raw("cupo_disponible + 1")).Error
}

// Actividades con estadísticas
func (d *ActivityDAO) FindWithStats() ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	err := utils.DB.Table("activities").
		Select(`activities.*, 
                COUNT(inscripciones.id) as inscripciones_actuales,
                (activities.cupo_maximo - activities.cupo_disponible) as ocupacion`).
		Joins("LEFT JOIN inscripciones ON activities.id = inscripciones.actividad_id AND inscripciones.deleted_at IS NULL").
		Where("activities.activo = ?", true).
		Group("activities.id").
		Scan(&results).Error

	return results, err
}

// Buscar por categoría
func (d *ActivityDAO) FindByCategory(categoria string) ([]domain.Activity, error) {
	var activities []domain.Activity
	err := utils.DB.Where("categoria = ? AND activo = ?", categoria, true).Find(&activities).Error
	return activities, err
}

// Buscar por día de la semana
func (d *ActivityDAO) FindByDay(diaSemana string) ([]domain.Activity, error) {
	var activities []domain.Activity
	err := utils.DB.Where("dia_semana = ? AND activo = ?", diaSemana, true).Find(&activities).Error
	return activities, err
}
