package handlers

import (
	"errors"

	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/models"
	"github.com/RegiAdi/hatchet/responses"
	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	LoginService(request models.User) (responses.UserLoginResponse, error)
}

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(
	authService AuthService,
) *AuthHandler {
	return &AuthHandler{
		authService,
	}
}

func (authHandler *AuthHandler) LoginHandler(c *fiber.Ctx) error {
	var request models.User
	var userLoginResponse responses.UserLoginResponse

	if err := c.BodyParser(&request); err != nil {
		return responses.SendResponse(c, responses.BaseResponse{
			StatusCode: kernel.StatusBadRequest,
			Status:     kernel.StatusFailed,
			Message:    "Failed to handle request",
			Data:       err,
		})
	}

	userLoginResponse, err := authHandler.authService.LoginService(request)
	if err != nil {
		switch {
		case errors.Is(err, kernel.ErrUserNotFound):
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusNotFound,
				Status:     kernel.StatusFailed,
				Message:    "User not found",
				Data:       nil,
			})
		case errors.Is(err, kernel.ErrPasswordUnmatch):
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusUnauthorized,
				Status:     kernel.StatusFailed,
				Message:    "Password do not match",
				Data:       nil,
			})
		case errors.Is(err, kernel.ErrGenerateAPITokenFailed):
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusBadRequest,
				Status:     kernel.StatusFailed,
				Message:    "Failed to generate API Token",
				Data:       nil,
			})
		case errors.Is(err, kernel.ErrInvalidObjectID):
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusBadRequest,
				Status:     kernel.StatusFailed,
				Message:    "Invalid ObjectID",
				Data:       nil,
			})
		case errors.Is(err, kernel.ErrUserUpdateFailed):
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusBadRequest,
				Status:     kernel.StatusFailed,
				Message:    "Failed to update user data",
				Data:       nil,
			})
		default:
			return responses.SendResponse(c, responses.BaseResponse{
				StatusCode: kernel.StatusBadRequest,
				Status:     kernel.StatusFailed,
				Message:    "Something wrong happened",
				Data:       nil,
			})
		}
	}

	return responses.SendResponse(c, responses.BaseResponse{
		StatusCode: kernel.StatusOK,
		Status:     kernel.StatusSuccess,
		Message:    "User authenticated successfully",
		Data:       userLoginResponse,
	})
}
