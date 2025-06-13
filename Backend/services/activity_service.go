package services

import (
	"errors"
	"strconv"
	"strings"

	"gym-backend/dao"
	"gym-backend/domain"
)

type ActivityService struct {
	activityDAO *dao.ActivityDAO
}

func NewActivityService() *ActivityService {
	return &ActivityService{
		activityDAO: dao.NewActivityDAO(),
	}
}

func (s *ActivityService) GetAllActivities() ([]domain.Activity, error) {
	return s.activityDAO.FindAll()
}

func (s *ActivityService) GetActivityByID(id uint) (*domain.Activity, error) {
	activity, err := s.activityDAO.FindByID(id)
	if err != nil {
		return nil, errors.New("actividad no encontrada")
	}
	return activity, nil
}

func (s *ActivityService) CreateActivity(req domain.CreateActivityRequest) (uint, error) {
	duracionStr := strings.TrimSuffix(req.Duracion, " min")
	duracion, err := strconv.Atoi(duracionStr)
	if err != nil {
		return 0, errors.New("duración inválida")
	}

	activity := domain.Activity{
		Titulo:         req.Titulo,
		Descripcion:    req.Descripcion,
		Categoria:      req.Categoria,
		Instructor:     req.Instructor,
		DiaSemana:      req.Dia,
		Horario:        req.Horario,
		Duracion:       duracion,
		CupoMaximo:     req.Cupo,
		CupoDisponible: req.Cupo,
		Foto:           req.FotoURL,
		Activo:         true,
	}

	err = s.activityDAO.Create(&activity)
	if err != nil {
		return 0, err
	}

	return activity.ID, nil
}

func (s *ActivityService) UpdateActivity(id uint, req domain.UpdateActivityRequest) error {
	activity, err := s.activityDAO.FindByID(id)
	if err != nil {
		return errors.New("actividad no encontrada")
	}

	if req.Titulo != nil {
		activity.Titulo = *req.Titulo
	}
	if req.Descripcion != nil {
		activity.Descripcion = *req.Descripcion
	}
	if req.Categoria != nil {
		activity.Categoria = *req.Categoria
	}
	if req.Instructor != nil {
		activity.Instructor = *req.Instructor
	}
	if req.Dia != nil {
		activity.DiaSemana = *req.Dia
	}
	if req.Horario != nil {
		activity.Horario = *req.Horario
	}
	if req.Duracion != nil {
		duracionStr := strings.TrimSuffix(*req.Duracion, " min")
		duracion, err := strconv.Atoi(duracionStr)
		if err != nil {
			return errors.New("duración inválida")
		}
		activity.Duracion = duracion
	}
	if req.Cupo != nil {
		diferencia := *req.Cupo - activity.CupoMaximo
		activity.CupoMaximo = *req.Cupo
		activity.CupoDisponible += diferencia
		if activity.CupoDisponible < 0 {
			activity.CupoDisponible = 0
		}
	}
	if req.FotoURL != nil {
		activity.Foto = *req.FotoURL
	}

	return s.activityDAO.Update(activity)
}

func (s *ActivityService) DeleteActivity(id uint) error {
	_, err := s.activityDAO.FindByID(id)
	if err != nil {
		return errors.New("actividad no encontrada")
	}
	return s.activityDAO.Delete(id)
}
