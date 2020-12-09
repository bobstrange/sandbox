// ref: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cli, err := mongo.NewClient(options.Client().ApplyURI("mongodb://user:password@localhost:37017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = cli.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)

	db := cli.Database("quickstart")
	podcastsColl := db.Collection("podcasts")
	episodesColl := db.Collection("episodes")

	// bson.D は Document
	// bson.A は Array
	// 実行時にエラーがなければ、 InsertOneResult が返ってくる
	// res.InsertedID で、挿入された ID がわかる
	podcastRes, err := podcastsColl.InsertOne(ctx, bson.D{
		{"title", "The Polyglot Developer Podcast"},
		{"author", "Nic Raboy"},
		{"tags", bson.A{"development", "programming", "coding"}},
	})

	if err != nil {
		log.Fatal(err)
	}

	// InsertMany の場合は、Insertしようとしているドキュメントの型を表す []interface{}
	// InsertManyResult を返す
	episodesRes, err := episodesColl.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastRes.InsertedID},
			{"title", "GraphQL for API Development"},
			{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastRes.InsertedID},
			{"title", "Progressive Web Application Development"},
			{"description", "Learn about PWA development with Tara Manicsic."},
			{"duration", 32},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v documents into episode collection!\n", len(episodesRes.InsertedIDs))
}
