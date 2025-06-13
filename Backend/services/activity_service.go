package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

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

// Función helper para convertir time.Time a string HH:MM
func timeToHourMinute(t time.Time) string {
	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}

// Función helper para convertir string HH:MM a time.Time
func hourMinuteToTime(timeStr string) (time.Time, error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return time.Time{}, errors.New("formato inválido, use HH:MM")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil || hour < 0 || hour > 23 {
		return time.Time{}, errors.New("hora inválida")
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil || minute < 0 || minute > 59 {
		return time.Time{}, errors.New("minuto inválido")
	}

	// Crear un time.Time con fecha base y la hora/minuto especificados
	return time.Date(1970, 1, 1, hour, minute, 0, 0, time.UTC), nil
}

func (s *ActivityService) GetAllActivities() ([]domain.ActivityListResponse, error) {
	activities, err := s.activityDAO.FindAll()
	if err != nil {
		return nil, err
	}

	var response []domain.ActivityListResponse
	for _, activity := range activities {
		response = append(response, domain.ActivityListResponse{
			ID:       activity.ID,
			Titulo:   activity.Titulo,
			Horario:  timeToHourMinute(activity.Horario), // Usar función helper
			Profesor: activity.Instructor,
		})
	}

	return response, nil
}

func (s *ActivityService) GetActivityByID(id uint) (*domain.ActivityDetailResponse, error) {
	activity, err := s.activityDAO.FindByID(id)
	if err != nil {
		return nil, errors.New("actividad no encontrada")
	}

	duracionStr := strconv.Itoa(activity.Duracion) + " min"

	response := &domain.ActivityDetailResponse{
		ID:          activity.ID,
		Titulo:      activity.Titulo,
		Descripcion: activity.Descripcion,
		Dia:         activity.DiaSemana,
		Horario:     timeToHourMinute(activity.Horario), // Usar función helper
		Duracion:    duracionStr,
		Cupo:        activity.CupoMaximo,
		Categoria:   activity.Categoria,
		Instructor:  activity.Instructor,
		FotoURL:     activity.Foto,
	}

	return response, nil
}

func (s *ActivityService) CreateActivity(req domain.CreateActivityRequest) (uint, error) {
	duracionStr := strings.TrimSuffix(req.Duracion, " min")
	duracion, err := strconv.Atoi(duracionStr)
	if err != nil {
		return 0, errors.New("duración inválida")
	}

	// Convertir horario string a time.Time usando función helper
	horario, err := hourMinuteToTime(req.Horario)
	if err != nil {
		return 0, err
	}

	activity := domain.Activity{
		Titulo:         req.Titulo,
		Descripcion:    req.Descripcion,
		Categoria:      req.Categoria,
		Instructor:     req.Instructor,
		DiaSemana:      req.Dia,
		Horario:        horario, // Ahora es time.Time
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
		// Convertir horario string a time.Time usando función helper
		horario, err := hourMinuteToTime(*req.Horario)
		if err != nil {
			return err
		}
		activity.Horario = horario // Ahora es time.Time
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
