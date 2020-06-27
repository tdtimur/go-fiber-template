package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/models"
)

type resp models.Response

func home(c *fiber.Ctx) {
	res, _ := json.Marshal(
		resp{
			StatusCode: 200,
			Message:    "success",
		})
	c.Send(res)
}

func login(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", home)
	api.Get("/login", login)
}
