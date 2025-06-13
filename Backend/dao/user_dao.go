package dao

import (
	"gym-backend/domain"
	"gym-backend/utils"
)

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Buscar usuario por email (para login)
func (d *UserDAO) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := utils.DB.Where("email = ? AND activo = ?", email, true).First(&user).Error
	return &user, err
}

// Buscar usuario por ID
func (d *UserDAO) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := utils.DB.Where("activo = ?", true).First(&user, id).Error
	return &user, err
}

// Crear nuevo usuario
func (d *UserDAO) Create(user *domain.User) error {
	return utils.DB.Create(user).Error
}

// Actualizar usuario
func (d *UserDAO) Update(user *domain.User) error {
	return utils.DB.Save(user).Error
}

// Obtener actividades del usuario con JOIN
func (d *UserDAO) GetUserActivities(userID uint) ([]domain.Activity, error) {
	var activities []domain.Activity

	// JOIN entre activities e inscripciones
	err := utils.DB.Table("activities").
		Joins("JOIN inscripciones ON activities.id = inscripciones.actividad_id").
		Where("inscripciones.usuario_id = ? AND activities.activo = ?", userID, true).
		Find(&activities).Error

	return activities, err
}

// Verificar si usuario ya está inscrito en actividad
func (d *UserDAO) IsUserEnrolled(userID, activityID uint) (bool, error) {
	var count int64
	err := utils.DB.Model(&domain.Inscription{}).
		Where("usuario_id = ? AND actividad_id = ?", userID, activityID).
		Count(&count).Error

	return count > 0, err
}

// Inscribir usuario a actividad
func (d *UserDAO) EnrollUser(userID, activityID uint) error {
	inscription := domain.Inscription{
		UsuarioID:   userID,
		ActividadID: activityID,
	}
	return utils.DB.Create(&inscription).Error
}

// Obtener estadísticas del usuario
func (d *UserDAO) GetUserStats(userID uint) (map[string]interface{}, error) {
	var stats map[string]interface{}

	// Contar inscripciones activas
	var totalInscripciones int64
	err := utils.DB.Model(&domain.Inscription{}).
		Where("usuario_id = ?", userID).
		Count(&totalInscripciones).Error

	if err != nil {
		return nil, err
	}

	stats = map[string]interface{}{
		"total_inscripciones": totalInscripciones,
	}

	return stats, nil
}
