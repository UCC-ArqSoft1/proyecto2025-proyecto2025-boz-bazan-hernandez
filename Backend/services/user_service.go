package services

import (
	"errors"
	"fmt"

	"gym-backend/dao"
	"gym-backend/domain"
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

func (s *UserService) EnrollInActivity(userID, activityID uint) error {
	activity, err := s.activityDAO.FindByID(activityID)
	if err != nil {
		return errors.New("actividad no encontrada")
	}

	if activity.CupoDisponible <= 0 {
		return errors.New("la actividad está llena")
	}

	enrolled, err := s.userDAO.IsUserEnrolled(userID, activityID)
	if err != nil {
		return err
	}
	if enrolled {
		return errors.New("ya estás inscrito en esta actividad")
	}

	err = s.userDAO.EnrollUser(userID, activityID)
	if err != nil {
		return fmt.Errorf("error al inscribirse: %v", err)
	}

	err = s.activityDAO.DecrementAvailableSlots(activityID)
	if err != nil {
		return fmt.Errorf("error actualizando cupo: %v", err)
	}

	return nil
}
