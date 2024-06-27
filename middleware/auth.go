package middleware

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

type Auth struct {
	userRepository UserRepository
}

func NewAuthMiddleware(userRepository UserRepository) *Auth {
	return &Auth{
		userRepository,
	}
}

func (auth *Auth) getAuthenticatedUser(APIToken string) (models.User, error) {
	return auth.userRepository.GetAuthenticatedUser(APIToken)
}

func (auth *Auth) isAPITokenExpired(user models.User) bool {
	return helpers.GetCurrentTime().After(user.TokenExpiresAt)
}

func (auth *Auth) setAPITokenToExpired(user models.User) error {
	return auth.userRepository.UpdateAPITokenExpirationTime(user.ID)
}

func (auth *Auth) setAPITokenLastUsedTime(user models.User) error {
	return auth.userRepository.UpdateAPITokenLastUsedTime(user.ID)
}

func (auth *Auth) Handler() fiber.Handler {
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

		user, err := auth.getAuthenticatedUser(APIToken)
		if err != nil {
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusUnauthorized,
				Status:     kernel.StatusFailed,
				Message:    "Unauthorized access",
				Data:       nil,
			})
		}

		if auth.isAPITokenExpired(user) {
			err := auth.setAPITokenToExpired(user)
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

		err = auth.setAPITokenLastUsedTime(user)

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
