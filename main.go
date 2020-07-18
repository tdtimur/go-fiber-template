package main

import (
	"context"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func Setup() *fiber.App {
	var pCtx = context.Background()
	var mongoHost = os.Getenv("MONGODB_HOST")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoHost))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(pCtx, 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
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
