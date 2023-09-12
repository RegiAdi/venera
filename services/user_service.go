package services

import "github.com/RegiAdi/hatchet/responses"

type UserRepository interface {
	GetUserByAPIToken(APIToken string) (responses.UserResponse, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		userRepository,
	}
}

func (userService *UserService) GetUserDetail(APIToken string) (responses.UserResponse, error) {
	userResponse, err := userService.userRepository.GetUserByAPIToken(APIToken)

	return userResponse, err
}
