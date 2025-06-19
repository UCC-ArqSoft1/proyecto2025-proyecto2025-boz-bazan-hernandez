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
		User:  user.ToUserInfo(),
	}

	return response, nil
}

func (s *AuthService) Register(req domain.RegisterRequest) (*domain.RegisterResponse, error) {
	// Verificar si el email ya existe
	exists, err := s.userDAO.EmailExists(req.Email)
	if err != nil {
		return nil, errors.New("error verificando email")
	}

	if exists {
		return nil, errors.New("el email ya está registrado")
	}

	// Hashear la contraseña - ajustado según tu función utils.HashPassword
	hashedPassword := utils.HashPassword(req.Password)

	// Crear nuevo usuario
	user := &domain.User{
		Nombre:       req.Nombre,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		TipoUsuario:  req.TipoUsuario,
		Activo:       true,
	}

	err = s.userDAO.Create(user)
	if err != nil {
		return nil, errors.New("error creando usuario")
	}

	response := &domain.RegisterResponse{
		Message: "Usuario registrado exitosamente",
		User:    user.ToUserInfo(),
	}

	return response, nil
}
