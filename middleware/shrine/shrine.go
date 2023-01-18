package shrine

import (
	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "Hello, World")
		return c.Next()
	}
}