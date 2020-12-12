package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cli, err := mongo.NewClient(options.Client().ApplyURI("mongodb://user:password@localhost:37017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = cli.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)

	db := cli.Database("quickstart")
	podcastsColl := db.Collection("podcasts")
	// episodesColl := db.Collection("episodes")

	id, _ := primitive.ObjectIDFromHex("5fd0e45b62af2dde1fa29046")

	var podcast bson.M
	podcastsColl.FindOne(ctx, bson.M{"_id": id})

	if podcastsColl.FindOne(ctx, bson.M{"_id": id}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Debug: FindOne(): ", podcast)

	res, err := podcastsColl.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"author", "John Doe"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", res.ModifiedCount)
}
