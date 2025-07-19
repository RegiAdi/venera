package routes

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/RegiAdi/venera/handlers"
	"github.com/RegiAdi/venera/kernel"
	"github.com/RegiAdi/venera/middleware"
	"github.com/RegiAdi/venera/repositories"
	"github.com/RegiAdi/venera/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
)

var (
	Logger zerolog.Logger
)

func API(appKernel *kernel.AppKernel) {
	// repositories
	userRepository := repositories.NewUserRepository(appKernel.DB)
	productRepository := repositories.NewProductRepository(appKernel.DB)

	// services
	authService := services.NewAuthService(userRepository)
	userService := services.NewUserService(userRepository)
	productService := services.NewProductService(productRepository)

	// handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)

	API := appKernel.Server.Group("/api")

	auth := API.Group("/auth")

	auth.Post("/login", authHandler.LoginHandler)
	auth.Post("/register", authHandler.RegisterHandler)

	authMiddleware := middleware.NewAuthMiddleware(userRepository)
	appKernel.Server.Use(authMiddleware.Handler())

	appKernel.Server.Use(logger.New(logger.Config{
		TimeZone:      "Asia/Jakarta",
		TimeFormat:    "2006-01-02 15:04:05",
		DisableColors: true,
		Format:        "{\"time\":\"${time}\",\"status\":${status},\"method\":\"${method}\",\"host\":\"${host}\",\"url\":\"${url}\",\"path\":\"${path}\",\"queryParams\":\"${queryParams}\",\"body\":\"${body_request}\",\"response_body\":\"${resBody}\",\"ip\":\"${ip}\",\"latency\":\"${latency}\"}\n",
		CustomTags: map[string]logger.LogFunc{
			"body_request": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				if c.Request().Body() == nil {
					return output.WriteString("")
				}
				buffer := new(bytes.Buffer)
				if err := json.Compact(buffer, c.Request().Body()); err != nil {
					fmt.Println("error json compact: ", err)
				}
				return output.WriteString(buffer.String())
			},
		},
	}))

	API.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	API.Post("/logout", authHandler.LogoutHandler)

	API.Get("/me", userHandler.GetUserInfoHandler)

	// Product routes
	API.Get("/products", productHandler.GetProductsHandler)
	API.Get("/products/:id", productHandler.GetProductHandler)
	API.Post("/products", productHandler.CreateProductHandler)
	API.Put("/products/:id", productHandler.UpdateProductHandler)
	API.Delete("/products/:id", productHandler.DeleteProductHandler)
}
