package main

import (
	"log"

	"github.com/RegiAdi/pos-mobile-backend/bootstrap"
	"github.com/RegiAdi/pos-mobile-backend/config"
	"github.com/RegiAdi/pos-mobile-backend/middleware/shrine"
	"github.com/RegiAdi/pos-mobile-backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	bootstrap.Run()

	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	app.Use(shrine.New())
	
	api := app.Group("/api")

	routes.AuthRoute(api.Group("/auth"))

    err := app.Listen(":" + config.GetAppPort())

    if err != nil {
        log.Fatal("App failed to start.")
        panic(err)
    }
}