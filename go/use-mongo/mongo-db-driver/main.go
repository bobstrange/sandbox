package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ref: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// user と password は URI に含める
	clientOptions := options.Client().ApplyURI("mongodb://user:password@localhost:37017")

	// MongoDB に接続する
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// コネクションのチェック
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection := client.Database("dev").Collection("trainers")

	john := Trainer{"John", 10, "Tokyo"}
	mike := Trainer{"Mike", 20, "Kanagawa"}
	susan := Trainer{"Susan", 25, "Saitama"}

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), john)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document:", insertResult.InsertedID)

	// Insert multiple documents
	trainers := []interface{}{mike, susan}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents:", insertManyResult.InsertedIDs)

	// Update documents
	filter := bson.D{{"name", "John"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updatedResult.MatchedCount, updatedResult.ModifiedCount)

	// Find a single document
	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	// Find multiple documents
	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Trainer

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	cur.All(context.TODO(), results)
	// for cur.Next(context.TODO()) {
	// 	var elem Trainer
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	results = append(results, &elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// 読み込みが完了したらカーソルをクローズする
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	// Delete documents
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	defer client.Disconnect(context.TODO())

}
