package routes

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/db"
	"gitlab.com/tdtimur/go-fiber-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"time"
)

type resp models.Response
type respList models.ResponseUsersList

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
	email := c.FormValue("email")
	password := c.FormValue("password")
	var res bson.M
	err := usersColl.FindOne(mg.Ctx, bson.D{{"email", email}}).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}

	if dbPassword := res["password"]; dbPassword != password {
		res := resp{
			StatusCode: 400,
			Message:    "Password did not match",
		}
		if err := c.JSON(res); err != nil {
			c.Status(500).Send(err)
			return
		}
	} else {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = email
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			log.Println(err)
			c.SendStatus(fiber.StatusInternalServerError)
			return
		}

		if err := c.JSON(fiber.Map{"token": t}); err != nil {
			log.Println(err)
			c.Status(500).Send(err)
			return
		}
	}
}
