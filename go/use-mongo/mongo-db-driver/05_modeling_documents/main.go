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

// Podcast podcast data structure
type Podcast struct {
	// bosn のアノテーションは、MongoDB Document のフィールド名と対応する
	// omitempty は、対象のフィールドのデータが存在しない場合は、MongoDB の Document のフィールドとしては存在しない
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   interface{}        `bson:"tags,omitempty"`
}

func main() {
	podcast := Podcast{
		Title:  "Test podcast 01",
		Author: "John Doe",
		Tags: []struct{ Name, Value string }{
			{Name: "Genre", Value: "Programming"},
			{Name: "Language", Value: "English"},
		},
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://user:password@localhost:37017"))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)

	db := cli.Database("quickstart")
	podcastsColl := db.Collection("podcasts")

	// 事前定義したデータ構造を使った Insert
	insertRes, err := podcastsColl.InsertOne(
		ctx,
		podcast,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Debug: InsertOne():", insertRes.InsertedID)

	// 事前定義したデータ構造を使った Find
	var podcasts []Podcast
	cur, err := podcastsColl.Find(
		ctx,
		bson.M{"author": "John Doe"},
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = cur.All(ctx, &podcasts); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Debug: retrived podcasts:", podcasts)
}
