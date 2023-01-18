package routes

import (
	"github.com/RegiAdi/pos-mobile-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(route fiber.Router) {
	route.Post("/login", controllers.Login)
	route.Post("/register", controllers.Register)
}