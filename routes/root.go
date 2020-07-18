package routes

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var pCtx = context.Background()
var mongoHost = os.Getenv("MONGODB_HOST")

type resp models.Response
type respList models.ResponseUsersList
type tokenResp models.JwtToken

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
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoHost))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(pCtx, 5*time.Second)
	defer ctx.Done()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	usersColl := client.Database("api").Collection("users")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		log.Println("Parser")
		log.Fatal(err)
	}
	count, _ := usersColl.CountDocuments(ctx, bson.D{{"email", user.Email}})
	if count == 0 {
		_, err := usersColl.InsertOne(ctx, user)
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

	email := c.FormValue("email")
	password := c.FormValue("password")
	var res bson.M
	err = usersColl.FindOne(ctx, bson.D{{"email", email}}).Decode(&res)
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

		res := tokenResp{
			Token: t,
		}
		if err := c.JSON(res); err != nil {
			log.Println(err)
			c.Status(500).Send(err)
			return
		}
	}
}
