package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("req: %v\n", req)
	log.Printf("QueryStringParameters: %v\n", req.QueryStringParameters)

	name, exists := req.QueryStringParameters["name"]
	if !exists {
		log.Fatal("please specify name.")
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprint("hi there"),
			StatusCode: 400,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint("Hi ", name),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
