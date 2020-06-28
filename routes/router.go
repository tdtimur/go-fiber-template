package routes

import "github.com/gofiber/fiber"

func SetupRoot(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", home)
	api.Get("/login", login)
	api.Post("/register", register)
}
