package main

import (
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)
	_ = app.Listen(3000)
}
