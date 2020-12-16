package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client *mongo.Client
)

func handler(event events.SNSEvent) error {
	log.Println("Debug: start handler()")
	log.Printf("Debug: event: %v\n", event)
	log.Println("Debug: done handler()")
	return nil
}

func main() {
	log.Println("Debug: start main()")

	// url := "http://172.17.0.1:8000"
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal("Failed to get: ", err)
	// }
	// defer resp.Body.Close()

	// buf, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal("Failed to read: ", err)
	// }
	// log.Println(string(buf))

	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://user:password@172.17.0.1:47017"),
	)

	if err != nil {
		log.Fatal("Fail to create client")
	}

	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal("Fail to connect")
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Fail to ping")
	}

	log.Println("Debug: lambda.Start()")
	lambda.Start(handler)
}
