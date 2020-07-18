package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/routes"
	"log"
)

func Setup() *fiber.App {
	app := fiber.New()
	corsConfig := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{"*"},
	}
	app.Use(cors.New(corsConfig))
	routes.SetupRoot(app)
	return app
}

func main() {
	app := Setup()
	defer func() {
		if err := app.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()
	if err := app.Listen(3000); err != nil {
		log.Fatal(err)
	}
}
