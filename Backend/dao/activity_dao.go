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

// Soft delete - marcar como inactivo
func (d *ActivityDAO) Delete(id uint) error {
	return utils.DB.Model(&domain.Activity{}).
		Where("id = ?", id).
		Update("activo", false).Error
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
