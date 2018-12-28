package store

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Store struct {
	client *mongo.Database
	Video  Video
}

func CreateStore() Store {
	mongohost := "mongodb://localhost:27017"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, mongohost)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to mongodb: %s", mongohost))
	}

	dbClient := client.Database("warehouse") //.Collection("numbers")
	videoStore := CreateVideoStore(dbClient)
	newStore := Store{
		client: dbClient,
		Video:  videoStore,
	}

	return newStore
}
