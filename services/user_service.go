package services

import "github.com/RegiAdi/hatchet/responses"

type userRepository interface {
	GetUserByAPIToken(APIToken string) (responses.UserResponse, error)
}

type userService struct {
	userRepository userRepository
}

func NewUserService(userRepository userRepository) *userService {
	return &userService{
		userRepository,
	}
}

func (userService *userService) GetUserDetail(APIToken string) (responses.UserResponse, error) {
	return userService.userRepository.GetUserByAPIToken(APIToken)
}
