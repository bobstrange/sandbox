package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.SNSEvent) error {
	log.Println("Debug: start handler()")
	log.Printf("Debug: event: %v\n", event)
	log.Println("Debug: done handler()")
	return nil
}

func main() {
	lambda.Start(handler)
}
