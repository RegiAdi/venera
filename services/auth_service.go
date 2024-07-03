package services

import (
	"github.com/RegiAdi/venera/helpers"
	"github.com/RegiAdi/venera/kernel"
	"github.com/RegiAdi/venera/models"
	"github.com/RegiAdi/venera/responses"
)

type UserRepository interface {
	GetUserByUsername(username string) (models.User, error)
	UpdateAPIToken(request models.User) (responses.UserLoginResponse, error)
}

type AuthService struct {
	userRepository UserRepository
}

func NewAuthService(
	userRepository UserRepository,
) *AuthService {
	return &AuthService{
		userRepository,
	}
}

func (authService *AuthService) LoginService(request models.User) (responses.UserLoginResponse, error) {
	var user models.User

	user, err := authService.userRepository.GetUserByUsername(request.Username)

	if err != nil {
		return responses.UserLoginResponse{}, kernel.ErrUserNotFound
	}

	if !helpers.CheckPasswordHash(request.Password, user.Password) {
		return responses.UserLoginResponse{}, kernel.ErrPasswordUnmatch
	}

	var userLoginResponse responses.UserLoginResponse
	userLoginResponse, err = authService.userRepository.UpdateAPIToken(user)

	return userLoginResponse, err
}
