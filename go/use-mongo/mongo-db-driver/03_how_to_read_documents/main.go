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

	// 結果セットは、 []bson.M でとりあえず受けることができる
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

	// 1 件取得してくる場合は、 FindOne() を使う
	podcastColl := db.Collection("podcasts")
	var podcast bson.M
	if err = podcastColl.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Debug: podcastColl.FindOne().Decode():", podcast)

	// 条件をつけて検索
	cur, err = episodesColl.Find(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = cur.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Debug: episodesFiltered:", episodesFiltered)

	// ソート は findOptions で設定する
	fOpts := options.Find()
	fOpts.SetSort(bson.D{{Key: "duration", Value: -1}})
	cur, err = episodesColl.Find(ctx, bson.D{{Key: "duration", Value: bson.D{{Key: "$gt", Value: 24}}}}, fOpts)
	if err != nil {
		log.Fatal(err)
	}
	var episodesSorted []bson.M
	if err = cur.All(ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Debug episodesSorted: ", episodesSorted)
}
