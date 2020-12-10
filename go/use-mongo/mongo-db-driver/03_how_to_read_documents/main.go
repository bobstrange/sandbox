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
	episodesColl := db.Collection("episodes")
	cur, err := episodesColl.Find(ctx, bson.M{})
	var episodes []bson.M

	// 結果セットが大きくない場合は、cursol.All() を使うと良い
	if err = cur.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)

	// 結果セットが大きくなりそうな場合は、逐次的にデータを取得してくる
	cur, err = episodesColl.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var epi bson.M
		if err = cur.Decode(&epi); err != nil {
			log.Fatal(err)
		}
		fmt.Println("cur.Next()", epi)
	}
}
