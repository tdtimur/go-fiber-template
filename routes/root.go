package routes

import (
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/db"
	"gitlab.com/tdtimur/go-fiber-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type resp models.Response
type respList models.ResponseList

var mg = db.Mg
var usersColl = db.UsersColl

func home(c *fiber.Ctx) {
	res := resp{
		StatusCode: 200,
		Message:    "API Home",
	}
	if err := c.JSON(res); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func register(c *fiber.Ctx) {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		log.Println("Parser")
		log.Fatal(err)
	}
	count, _ := usersColl.CountDocuments(mg.Ctx, bson.D{{"email", user.Email}})
	if count == 0 {
		_, err := usersColl.InsertOne(mg.Ctx, user)
		if err != nil {
			log.Fatal(err)
		}
		res := resp{
			StatusCode: 201,
			Message:    "User added",
		}
		if err := c.JSON(res); err != nil {
			c.Status(500).Send(err)
			return
		}
	} else {
		res := resp{
			StatusCode: 400,
			Message:    "User already exists",
		}
		if err := c.JSON(res); err != nil {
			c.Status(500).Send(err)
			return
		}
	}
}

func login(c *fiber.Ctx) {
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

func SetupRoot(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", home)
	api.Get("/login", login)
	api.Post("/register", register)
}
