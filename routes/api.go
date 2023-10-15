package routes

import (
	"github.com/RegiAdi/hatchet/controllers"
	"github.com/RegiAdi/hatchet/handlers"
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/middleware/shrine"
	"github.com/RegiAdi/hatchet/repositories"
	"github.com/RegiAdi/hatchet/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func API(appKernel *kernel.AppKernel) {
	// repositories
	userRepository := repositories.NewUserRepository(appKernel.DB)

	// services
	userService := services.NewUserService(userRepository)

	// handlers
	userHandler := handlers.NewUserHandler(userService)

	API := appKernel.Server.Group("/api")

	API.Group("/auth")
	auth := API.Group("/auth")

	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

	shrine := shrine.New(userRepository)
	appKernel.Server.Use(shrine.Handler())
	appKernel.Server.Use(logger.New(logger.Config{
		TimeZone:      "Asia/Jakarta",
		TimeFormat:    "2006-01-02 15:04:05",
		DisableColors: true,
		Format:        `{"time":"${time}","status":${status},"method":"${method}","host":"${host}","url":"${url}","path":"${path}","reqHeaders":"${reqHeaders}","queryParams":"${queryParams}","body":"${body}","ip":"${ip}"}`,
	}))

	API.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	auth.Get("/logout", controllers.Logout)

	API.Get("/me", userHandler.GetUserInfoHandler)

	API.Get("/products", controllers.GetProducts)
	API.Get("/products/:id", controllers.GetProduct)
	API.Post("/products", controllers.CreateProduct)
	API.Put("/products/:id", controllers.UpdateProduct)
	API.Delete("/products/:id", controllers.DeleteProduct)
}
