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
			{
				"$set",
				bson.D{
					{"author", "John Doe"},
					{"title", "Whatever"},
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	// 差分が無い場合は ModifiedCount -> 0 になる
	fmt.Printf("Updated %v Documents!\n", res.ModifiedCount)

	res, err = podcastsColl.UpdateMany(
		ctx,
		bson.M{"title": "The Polyglot Developer Podcast"},
		bson.D{
			{"$set", bson.D{{"author", "Jane Doe"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v documents!\n", res.ModifiedCount)

	res, err = podcastsColl.ReplaceOne(
		ctx,
		bson.M{"author": "Jane Doe"},
		bson.M{
			"title":  "Jane Doe's podcast",
			"author": "Jane Doe",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Debug: ReplaceOne() %v Documents\n", res.ModifiedCount)
}
