package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Mg = createDb(os.Getenv("MONGODB_HOST"))
var UsersColl = Mg.Client.Database("api").Collection("users")

type database struct {
	Ctx    context.Context
	Client *mongo.Client
}

func createDb(uri string) *database {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return &database{ctx, client}
}
