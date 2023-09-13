package handlers

import (
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/responses"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	GetUserByAPIToken(APIToken string) (responses.UserResponse, error)
}

type UserService interface {
	GetUserDetail(APIToken string) (responses.UserResponse, error)
}

type UserHandler struct {
	appKernel      *kernel.AppKernel
	userRepository UserRepository
	userService    UserService
}

func NewUserHandler(
	appKernel *kernel.AppKernel,
	userRepository UserRepository,
	userService UserService,
) *UserHandler {
	return &UserHandler{
		appKernel,
		userRepository,
		userService,
	}
}

func (userHandler *UserHandler) GetUserInfoHandler(c *fiber.Ctx) error {
	var userResponse responses.UserResponse
	reqHeader := c.GetReqHeaders()
	APIToken := reqHeader["Authorization"]

	userResponse, err := userHandler.userService.GetUserDetail(APIToken)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "Failed",
			"message": "User not found",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "User retrieved successfully",
		"data":    userResponse,
	})
}
