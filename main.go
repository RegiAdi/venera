package main

import (
	"log"

	"github.com/RegiAdi/pos-mobile-backend/bootstrap"
	"github.com/RegiAdi/pos-mobile-backend/config"
	"github.com/RegiAdi/pos-mobile-backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	bootstrap.Run()
	routes.API(app)

    err := app.Listen(":" + config.GetAppPort())

    if err != nil {
        log.Fatal("App failed to start.")
        panic(err)
    }
}