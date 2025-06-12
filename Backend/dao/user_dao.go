package dao

import (
	"gym-backend/domain"
	"gym-backend/utils"
)

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (d *UserDAO) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := utils.DB.Where("email = ? AND activo = ?", email, true).First(&user).Error
	return &user, err
}

func (d *UserDAO) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := utils.DB.Where("activo = ?", true).First(&user, id).Error
	return &user, err
}

func (d *UserDAO) Create(user *domain.User) error {
	return utils.DB.Create(user).Error
}

func (d *UserDAO) Update(user *domain.User) error {
	return utils.DB.Save(user).Error
}

func (d *UserDAO) GetUserActivities(userID uint) ([]domain.Activity, error) {
	var activities []domain.Activity

	err := utils.DB.Table("activities").
		Joins("JOIN inscripciones ON activities.id = inscripciones.actividad_id").
		Where("inscripciones.usuario_id = ? AND activities.activo = ?", userID, true).
		Find(&activities).Error

	return activities, err
}

func (d *UserDAO) IsUserEnrolled(userID, activityID uint) (bool, error) {
	var count int64
	err := utils.DB.Model(&domain.Inscription{}).
		Where("usuario_id = ? AND actividad_id = ?", userID, activityID).
		Count(&count).Error

	return count > 0, err
}

func (d *UserDAO) EnrollUser(userID, activityID uint) error {
	inscription := domain.Inscription{
		UsuarioID:   userID,
		ActividadID: activityID,
	}
	return utils.DB.Create(&inscription).Error
}
