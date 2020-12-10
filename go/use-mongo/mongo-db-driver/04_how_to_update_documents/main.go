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

	id, _ := primitive.ObjectIDFromHex("5fd0e469785ebdf7eae66137")
	res, err := podcastsColl.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{Key: "$set", Value: bson.D{{Key: "author", Value: "Nic Raboy"}}}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", res.ModifiedCount)
}
