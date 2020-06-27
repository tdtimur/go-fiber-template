package main

import (
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/routes"
	"log"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	err := app.Listen(3000)
	if err != nil {
		log.Fatal(err)
	}
}
