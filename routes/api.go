package routes

import (
	"github.com/RegiAdi/hatchet/controllers"
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/middleware/shrine"
	"github.com/RegiAdi/hatchet/repositories"
	"github.com/gofiber/fiber/v2"
)

func API(app *fiber.App) {
	API := app.Group("/api")

	API.Group("/auth")
	auth := API.Group("/auth")

	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

	app.Use(shrine.New())

	API.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	auth.Get("/logout", controllers.Logout)

	DB := kernel.Mongo.DB
	userRepository := repositories.NewUserRepository(DB)

	userController := controllers.NewUserController(userRepository)

	API.Get("/userinfo", userController.GetUserInfo)

	API.Get("/products", controllers.GetProducts)
	API.Get("/products/:id", controllers.GetProduct)
	API.Post("/products", controllers.CreateProduct)
	API.Put("/products/:id", controllers.UpdateProduct)
	API.Delete("/products/:id", controllers.DeleteProduct)
}
