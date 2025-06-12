package dao

import (
	"gym-backend/domain"
	"gym-backend/utils"
)

type ActivityDAO struct{}

func NewActivityDAO() *ActivityDAO {
	return &ActivityDAO{}
}

func (d *ActivityDAO) FindAll() ([]domain.Activity, error) {
	var activities []domain.Activity
	err := utils.DB.Where("activo = ?", true).Find(&activities).Error
	return activities, err
}

func (d *ActivityDAO) FindByID(id uint) (*domain.Activity, error) {
	var activity domain.Activity
	err := utils.DB.Where("activo = ?", true).First(&activity, id).Error
	return &activity, err
}

func (d *ActivityDAO) Create(activity *domain.Activity) error {
	return utils.DB.Create(activity).Error
}

func (d *ActivityDAO) Update(activity *domain.Activity) error {
	return utils.DB.Save(activity).Error
}

func (d *ActivityDAO) Delete(id uint) error {
	return utils.DB.Model(&domain.Activity{}).Where("id = ?", id).Update("activo", false).Error
}

func (d *ActivityDAO) DecrementAvailableSlots(activityID uint) error {
	return utils.DB.Model(&domain.Activity{}).
		Where("id = ? AND cupo_disponible > 0", activityID).
		Update("cupo_disponible", utils.DB.Raw("cupo_disponible - 1")).Error
}
