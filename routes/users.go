package routes

import (
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func usersList(c *fiber.Ctx) {
	var result []models.User
	cur, err := usersColl.Find(mg.Ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	if err = cur.All(mg.Ctx, &result); err != nil {
		log.Fatal(err)
	}
	res := respList{
		StatusCode: 200,
		Message:    "List of users",
		Result:     result,
	}
	if err := c.JSON(res); err != nil {
		c.Status(500).Send(err)
		return
	}
}
