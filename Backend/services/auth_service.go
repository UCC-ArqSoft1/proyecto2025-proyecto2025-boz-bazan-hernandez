package services

import (
	"errors"

	"gym-backend/dao"
	"gym-backend/domain"
	"gym-backend/utils"
)

type AuthService struct {
	userDAO *dao.UserDAO
}

func NewAuthService() *AuthService {
	return &AuthService{
		userDAO: dao.NewUserDAO(),
	}
}

func (s *AuthService) Login(req domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := s.userDAO.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("credenciales inválidas")
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.TipoUsuario)
	if err != nil {
		return nil, errors.New("error generando token")
	}

	response := &domain.LoginResponse{
		Token: token,
		Role:  user.GetRole(),
	}

	return response, nil
}
