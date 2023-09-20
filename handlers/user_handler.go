package handlers

import (
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/responses"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	GetUserDetail(APIToken string) (responses.UserResponse, error)
}

type UserHandler struct {
	appKernel   *kernel.AppKernel
	userService UserService
}

func NewUserHandler(
	appKernel *kernel.AppKernel,
	userService UserService,
) *UserHandler {
	return &UserHandler{
		appKernel,
		userService,
	}
}

func (userHandler *UserHandler) GetUserInfoHandler(c *fiber.Ctx) error {
	var userResponse responses.UserResponse
	reqHeader := c.GetReqHeaders()
	APIToken := reqHeader["Authorization"]

	userResponse, err := userHandler.userService.GetUserDetail(APIToken)

	if err != nil {
		return responses.SendResponse(c, responses.BaseResponse{
			StatusCode: kernel.StatusNotFound,
			Status:     kernel.StatusFailed,
			Message:    "User not found",
			Data:       nil,
		})
	}

	return responses.SendResponse(c, responses.BaseResponse{
		StatusCode: kernel.StatusOK,
		Status:     kernel.StatusSuccess,
		Message:    "User retrieved successfully",
		Data:       userResponse,
	})
}
