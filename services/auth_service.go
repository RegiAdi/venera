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
	GetUserByEmail(email string) (models.User, error)
	CreateUser(request models.User) (responses.UserResponse, error)
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

// RegisterService handles user registration logic
func (authService *AuthService) RegisterService(request models.User) (responses.UserResponse, error) {
	// Check if username already exists
	existingUser, _ := authService.userRepository.GetUserByUsername(request.Username)
	if existingUser.Username != "" {
		return responses.UserResponse{}, kernel.ErrUserAlreadyExists
	}

	existingUserByEmail, _ := authService.userRepository.GetUserByEmail(request.Email)
	if existingUserByEmail.Email != "" {
		return responses.UserResponse{}, kernel.ErrEmailAlreadyExists
	}

	// Hash the password
	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return responses.UserResponse{}, err
	}
	request.Password = hashedPassword
	request.CreatedAt = helpers.GetCurrentTime()
	request.UpdatedAt = helpers.GetCurrentTime()

	userResponse, err := authService.userRepository.CreateUser(request)
	if err != nil {
		return responses.UserResponse{}, err
	}

	return userResponse, nil
}
