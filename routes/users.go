package routes

import (
	"context"
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func usersList(c *fiber.Ctx) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoHost))
	if err != nil {
		log.Println(err)
		return
	}
	ctx, _ := context.WithTimeout(pCtx, 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	usersColl := client.Database("api").Collection("users")

	var result []models.User
	cur, err := usersColl.Find(ctx, bson.D{{}})
	if err != nil {
		log.Println(err)
		return
	}
	if err = cur.All(ctx, &result); err != nil {
		log.Println(err)
		return
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
