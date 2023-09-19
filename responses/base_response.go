package responses

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	StatusCode int
	Status     string
	Message    string
	Data       interface{}
}

func SendResponse(Ctx *fiber.Ctx, response BaseResponse) error {
	return Ctx.Status(response.StatusCode).JSON(fiber.Map{
		"status":  response.Status,
		"message": response.Message,
		"data":    response.Data,
	})
}
