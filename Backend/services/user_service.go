package services

import (
    "errors"
    "fmt"
    
    "gym-backend/dao"
    "gym-backend/domain"
    "gym-backend/utils" // Para acceder a utils.DB
    "gorm.io/gorm"
)

type UserService struct {
    userDAO     *dao.UserDAO
    activityDAO *dao.ActivityDAO
}

func NewUserService() *UserService {
    return &UserService{
        userDAO:     dao.NewUserDAO(),
        activityDAO: dao.NewActivityDAO(),
    }
}

//TRANSACCIÓN COMPLETA - Inscripción con validaciones
func (s *UserService) EnrollInActivity(userID, activityID uint) error {
    // Usar transacción GORM para operación atómica
    return utils.DB.Transaction(func(tx *gorm.DB) error {
        // 1. Verificar que la actividad existe y está activa
        var activity domain.Activity
        if err := tx.Where("activo = ?", true).First(&activity, activityID).Error; err != nil {
            return errors.New("actividad no encontrada")
        }
        
        // 2. Verificar cupo disponible
        if activity.CupoDisponible <= 0 {
            return errors.New("la actividad está llena")
        }
        
        // 3. Verificar que el usuario no esté ya inscrito
        var existingInscription domain.Inscription
        err := tx.Where("usuario_id = ? AND actividad_id = ?", userID, activityID).
            First(&existingInscription).Error
        
        if err == nil {
            return errors.New("ya estás inscrito en esta actividad")
        } else if !errors.Is(err, gorm.ErrRecordNotFound) {
            return fmt.Errorf("error verificando inscripción: %v", err)
        }
        
        // 4. Crear inscripción
        inscription := domain.Inscription{
            UsuarioID:   userID,
            ActividadID: activityID,
        }
        
        if err := tx.Create(&inscription).Error; err != nil {
            return fmt.Errorf("error creando inscripción: %v", err)
        }
        
        // 5. Decrementar cupo disponible
        if err := tx.Model(&activity).Update("cupo_disponible", activity.CupoDisponible-1).Error; err != nil {
            return fmt.Errorf("error actualizando cupo: %v", err)
        }
        
        return nil // Commit automático si no hay error
    })
}

// Resto de métodos usando DAO
func (s *UserService) GetUserActivities(userID uint) ([]domain.MyActivityResponse, error) {
    activities, err := s.userDAO.GetUserActivities(userID)
    if err != nil {
        return nil, err
    }
    
    var response []domain.MyActivityResponse
    for _, activity := range activities {
        response = append(response, domain.MyActivityResponse{
            ID:         activity.ID,
            Titulo:     activity.Titulo,
            Horario:    activity.Horario,
            Dia:        activity.DiaSemana,
            Instructor: activity.Instructor,
        })
    }
    
    return response, nil
}
