package routes

import (
	"github.com/RegiAdi/pos-mobile-backend/controllers"
	"github.com/RegiAdi/pos-mobile-backend/middleware/shrine"
	"github.com/gofiber/fiber/v2"
)

func API(app *fiber.App) {
	api := app.Group("/api")

	api.Group("/auth")
	auth := api.Group("/auth")

	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

	app.Use(shrine.New())

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	auth.Get("/logout", controllers.Logout)
	api.Get("/userinfo", controllers.GetUserInfo)

	api.Post("/products", controllers.CreateProduct)
}
