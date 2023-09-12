package handlers

import (
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/repositories"
	"github.com/RegiAdi/hatchet/responses"
	"github.com/RegiAdi/hatchet/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	appKernel *kernel.AppKernel
}

func NewUserHandler(appKernel *kernel.AppKernel) *UserHandler {
	return &UserHandler{
		appKernel,
	}
}

func (userHandler *UserHandler) GetUserInfoHandler(c *fiber.Ctx) error {
	var userResponse responses.UserResponse
	reqHeader := c.GetReqHeaders()
	APIToken := reqHeader["Authorization"]

	userRepository := repositories.NewUserRepository(userHandler.appKernel.DB)
	userService := services.NewUserService(userRepository)

	userResponse, err := userService.GetUserDetail(APIToken)

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
