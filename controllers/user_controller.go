package controllers

import (
	"github.com/RegiAdi/hatchet/responses"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	GetUserByAPIToken(APIToken string) (responses.UserResponse, error)
}

type UserController struct {
	userRepository UserRepository
}

func NewUserController(userRepository UserRepository) *UserController {
	return &UserController{
		userRepository,
	}
}

func (userController *UserController) GetUserInfo(c *fiber.Ctx) error {
	var userResponse responses.UserResponse
	reqHeader := c.GetReqHeaders()

	userResponse, err := userController.userRepository.GetUserByAPIToken(reqHeader["Authorization"])

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
