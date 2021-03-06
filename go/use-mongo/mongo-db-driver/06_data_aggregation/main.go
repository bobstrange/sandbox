package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/k0kubun/pp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Podcast podcast structure
type Podcast struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Author string             `json:"author" bson:"author,omitempty"`
	Tags   []interface{}      `json:"tags" bson:"tags,omitempty"`
}

// Episode episode structure
type Episode struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `json:"podcast_id" bson:"podcast,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Duration    int32              `json:"duration" bson:"duration,omitempty"`
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://user:password@localhost:37017"))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)

	db := cli.Database("quickstart")
	podcastsColl := db.Collection("podcasts")
	episodesColl := db.Collection("episodes")

	// Insert a podcast
	pInput := `
      {
        "id": "5e3b37e51c9d4400004117e6",
        "title": "Sample podcast",
        "author": "Adam Smith",
        "tags": [
          { "name": "tag1", "value": "tagv1" },
          { "name": "tag2", "value": "tagv2" }
        ]
      }
	`

	var podcast *Podcast
	if err := json.Unmarshal([]byte(pInput), &podcast); err != nil {
		log.Fatal(err)
	}
	pp.Print(podcast)

	res, err := podcastsColl.InsertOne(ctx, podcast)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Debug: inserted podcast _id: %v\n", res.InsertedID)

	// Insert some episodes
	file, err := os.Open("episodes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var episodes []interface{}
	if err := json.Unmarshal(input, &episodes); err != nil {
		log.Fatal(err)
	}
	pp.Print(episodes)

	episodesRes, err := episodesColl.InsertMany(ctx, episodes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Debug: inserted episodes _ids: %v\n", episodesRes.InsertedIDs)

}
