package store

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Store struct{}

func (st Store) Video(client) {
	return CreateVideoStore(client)
}

func CreateStore() {
	mongohost := "mongodb://localhost:27017"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, mongohost)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to mongodb: %s", mongohost))
	}

	dbClient := client.Database("testing") //.Collection("numbers")

	// {
	// Video:
	// 	CreateVideoStore(client)
	// }
	return Store
}
