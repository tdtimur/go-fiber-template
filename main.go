package main

import (
	"context"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"gitlab.com/tdtimur/go-fiber-template/models"
	"gitlab.com/tdtimur/go-fiber-template/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Setup() *fiber.App {
	var pCtx = context.Background()
	log.Printf(
		"Preparing... MongoDB host: %s, JWT Secret: %s",
		models.Config.GetString("MONGODB_HOST"),
		models.Config.GetString("JWT_SECRET"),
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(models.Config.GetString("MONGODB_HOST")))
	if err != nil {
		log.Println(err)
		return nil
	}
	ctx, _ := context.WithTimeout(pCtx, 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return nil
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
		log.Println(err)
		return
	}
}
