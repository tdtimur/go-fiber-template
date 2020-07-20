package routes

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"gitlab.com/tdtimur/go-fiber-template/models"
)

func SetupRoot(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", home)
	api.Post("/login", login)
	api.Post("/register", register)

	users := app.Group("/users")
	users.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(models.Config.GetString("JWT_SECRET")),
	}))
	users.Get("/list", usersList)
}
