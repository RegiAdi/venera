package shrine

import (
	"github.com/RegiAdi/hatchet/helpers"
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/models"
	"github.com/RegiAdi/hatchet/responses"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	GetAuthenticatedUser(APIToken string) (models.User, error)
	UpdateAPITokenExpirationTime(userID string) error
	UpdateAPITokenLastUsedTime(userID string) error
}

type Shrine struct {
	userRepository UserRepository
}

func New(userRepository UserRepository) *Shrine {
	return &Shrine{
		userRepository,
	}
}

func (shrine *Shrine) getAuthenticatedUser(APIToken string) (models.User, error) {
	return shrine.userRepository.GetAuthenticatedUser(APIToken)
}

func (shrine *Shrine) isAPITokenExpired(user models.User) bool {
	if helpers.GetCurrentTime().After(user.TokenExpiresAt) {
		return true
	}

	return false
}

func (shrine *Shrine) setAPITokenToExpired(user models.User) error {
	return shrine.userRepository.UpdateAPITokenExpirationTime(user.ID)
}

func (shrine *Shrine) setAPITokenLastUsedTime(user models.User) error {
	return shrine.userRepository.UpdateAPITokenLastUsedTime(user.ID)
}

func (shrine *Shrine) Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqHeader := c.GetReqHeaders()
		APIToken := reqHeader["Authorization"]

		if APIToken == "" {
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusUnauthorized,
				Status:     kernel.StatusFailed,
				Message:    "Unauthorized access",
				Data:       nil,
			})
		}

		user, err := shrine.getAuthenticatedUser(APIToken)
		if err != nil {
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusUnauthorized,
				Status:     kernel.StatusFailed,
				Message:    "Unauthorized access",
				Data:       nil,
			})
		}

		if shrine.isAPITokenExpired(user) {
			err := shrine.setAPITokenToExpired(user)
			if err != nil {
				return responses.SendResponse(c, responses.BaseResponse{
					StatusCode: kernel.StatusBadRequest,
					Status:     kernel.StatusFailed,
					Message:    "Can't delete API Token",
					Data:       nil,
				})
			}

			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusUnauthorized,
				Status:     kernel.StatusFailed,
				Message:    "API Token expired",
				Data:       nil,
			})

		}

		err = shrine.setAPITokenLastUsedTime(user)

		if err != nil {
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusBadRequest,
				Status:     kernel.StatusFailed,
				Message:    "Failed to update API Token Last Used Time",
				Data:       nil,
			})
		}

		return c.Next()
	}
}
