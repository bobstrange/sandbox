package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type deps struct {
	ddb   dynamodbiface.DynamoDBAPI
	table string
}

func (d deps) handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello"),
		StatusCode: 200,
	}, nil
}

func main() {
	sess := session.Must(session.NewSession())
	ddb := dynamodb.New(sess)

	d := deps{
		ddb:   ddb,
		table: "SomeTable",
	}
	lambda.Start(d.handler)
}
