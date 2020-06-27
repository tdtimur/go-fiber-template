package routes

import (
	"github.com/gofiber/fiber"
)

func home(c *fiber.Ctx) {
	c.Send("Home")
}

func login(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", home)

	// Auth
	auth := api.Group("/login")
	auth.Get("/login", login)
}
