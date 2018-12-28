package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-bento-api/structs"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Video struct {
	collection mongo.Collection
}

func (vd *Video) FindByID(id string) structs.Video {
	var result structs.Video

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"id": id}}
	err := vd.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(fmt.Sprintf("Error retreiving the document from mongo: (id: \"%v\")", id))
	}
	return result
}

func CreateVideoStore(client *mongo.Database) Video {
	videoColl := client.Collection("video_object")
	newVideo := Video{collection: *videoColl}
	return newVideo
}
