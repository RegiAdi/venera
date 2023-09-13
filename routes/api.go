package routes

import (
	"github.com/RegiAdi/hatchet/controllers"
	"github.com/RegiAdi/hatchet/handlers"
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/middleware/shrine"
	"github.com/RegiAdi/hatchet/repositories"
	"github.com/RegiAdi/hatchet/services"
	"github.com/gofiber/fiber/v2"
)

func API(appKernel *kernel.AppKernel) {
	API := appKernel.Server.Group("/api")

	API.Group("/auth")
	auth := API.Group("/auth")

	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

	appKernel.Server.Use(shrine.New())

	API.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	auth.Get("/logout", controllers.Logout)

	userRepository := repositories.NewUserRepository(appKernel.DB)
	userController := controllers.NewUserController(userRepository)
	API.Get("/userinfo", userController.GetUserInfo)

	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(appKernel, userService)
	API.Get("/me", userHandler.GetUserInfoHandler)

	API.Get("/products", controllers.GetProducts)
	API.Get("/products/:id", controllers.GetProduct)
	API.Post("/products", controllers.CreateProduct)
	API.Put("/products/:id", controllers.UpdateProduct)
	API.Delete("/products/:id", controllers.DeleteProduct)
}
