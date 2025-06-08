package services

import (
	"errors"

	"gym-management-system/config"
	"gym-management-system/models"

	"gorm.io/gorm"
)

// UserService maneja la lógica de usuarios
type UserService struct{}

// NewUserService crea una nueva instancia del servicio de usuarios
func NewUserService() *UserService {
	return &UserService{}
}

// GetUserByID obtiene un usuario por su ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	db := config.GetDB()
	var user models.User

	if err := db.Where("id = ? AND activo = ?", id, true).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	return &user, nil
}

// GetAllUsers obtiene todos los usuarios (solo para administradores)
func (s *UserService) GetAllUsers() ([]models.UserResponse, error) {
	db := config.GetDB()
	var users []models.User

	if err := db.Where("activo = ?", true).Find(&users).Error; err != nil {
		return nil, err
	}

	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	return userResponses, nil
}

// GetUserInscriptions obtiene las inscripciones de un usuario
func (s *UserService) GetUserInscriptions(userID uint) ([]models.InscriptionResponse, error) {
	db := config.GetDB()
	var inscriptions []models.Inscription

	if err := db.Preload("Activity").
		Where("usuario_id = ?", userID).
		Find(&inscriptions).Error; err != nil {
		return nil, err
	}

	var responses []models.InscriptionResponse
	for _, inscription := range inscriptions {
		if inscription.Activity.Activo { // Solo mostrar actividades activas
			responses = append(responses, inscription.ToResponse())
		}
	}

	return responses, nil
}

// CreateInscription crea una nueva inscripción
func (s *UserService) CreateInscription(userID uint, req *models.InscriptionRequest) (*models.InscriptionResponse, error) {
	db := config.GetDB()

	// Verificar que la actividad existe y está activa
	var activity models.Activity
	if err := db.Where("id = ? AND activo = ?", req.ActividadID, true).First(&activity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("actividad no encontrada")
		}
		return nil, err
	}

	// Verificar que hay cupo disponible
	if activity.CupoDisponible <= 0 {
		return nil, errors.New("no hay cupo disponible")
	}

	// Verificar que el usuario no esté ya inscripto
	var existingInscription models.Inscription
	if err := db.Where("usuario_id = ? AND actividad_id = ?", userID, req.ActividadID).First(&existingInscription).Error; err == nil {
		return nil, errors.New("ya estás inscripto en esta actividad")
	}

	// Iniciar transacción
	tx := db.Begin()

	// Crear inscripción
	inscription := models.Inscription{
		UsuarioID:   userID,
		ActividadID: req.ActividadID,
	}

	if err := tx.Create(&inscription).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error al crear la inscripción")
	}

	// Actualizar cupo disponible
	if err := tx.Model(&activity).Update("cupo_disponible", activity.CupoDisponible-1).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error al actualizar el cupo")
	}

	// Confirmar transacción
	tx.Commit()

	// Cargar la actividad en la inscripción para la respuesta
	if err := db.Preload("Activity").First(&inscription, inscription.ID).Error; err != nil {
		return nil, errors.New("error al cargar los datos de la inscripción")
	}

	response := inscription.ToResponse()
	return &response, nil
}

// CancelInscription cancela una inscripción
func (s *UserService) CancelInscription(userID, inscriptionID uint) error {
	db := config.GetDB()

	// Buscar la inscripción
	var inscription models.Inscription
	if err := db.Where("id = ? AND usuario_id = ?", inscriptionID, userID).First(&inscription).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("inscripción no encontrada")
		}
		return err
	}

	// Iniciar transacción
	tx := db.Begin()

	// Eliminar inscripción
	if err := tx.Delete(&inscription).Error; err != nil {
		tx.Rollback()
		return errors.New("error al cancelar la inscripción")
	}

	// Actualizar cupo disponible
	var activity models.Activity
	if err := tx.First(&activity, inscription.ActividadID).Error; err != nil {
		tx.Rollback()
		return errors.New("error al encontrar la actividad")
	}

	if err := tx.Model(&activity).Update("cupo_disponible", activity.CupoDisponible+1).Error; err != nil {
		tx.Rollback()
		return errors.New("error al actualizar el cupo")
	}

	// Confirmar transacción
	tx.Commit()

	return nil
}

// UpdateUserProfile actualiza el perfil del usuario
func (s *UserService) UpdateUserProfile(userID uint, nombre, email string) (*models.UserResponse, error) {
	db := config.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	// Verificar si el email ya existe (si es diferente al actual)
	if email != user.Email {
		var existingUser models.User
		if err := db.Where("email = ? AND id != ?", email, userID).First(&existingUser).Error; err == nil {
			return nil, errors.New("el email ya está en uso")
		}
	}

	// Actualizar datos
	user.Nombre = nombre
	user.Email = email

	if err := db.Save(&user).Error; err != nil {
		return nil, errors.New("error al actualizar el perfil")
	}

	response := user.ToResponse()
	return &response, nil
}
