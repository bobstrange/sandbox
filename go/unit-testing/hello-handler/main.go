package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/segmentio/ksuid"
)

type deps struct {
	ddb   dynamodbiface.DynamoDBAPI
	table string
}

type Order struct {
	ID   string `dynamodbav:"id"`
	Date time.Time
}

func (d *deps) handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if d.ddb == nil {
		sess := session.Must(session.NewSession())
		ddb := dynamodb.New(sess)
		d.ddb = ddb
		d.table = "SomeTable"
	}

	k, err := ksuid.NewRandom()
	if err != nil {
		fmt.Printf("Error generating KSUID: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	var o Order
	o.ID = k.String()
	o.Date = time.Now()

	item, err := dynamodbattribute.MarshalMap(o)
	if err != nil {
		fmt.Println("Could not call dynamodbattribute.MarshalMap(o)")
		fmt.Println(o)
		return events.APIGatewayProxyResponse{}, err
	}

	input := dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(d.table),
	}

	response, err := d.ddb.PutItem(&input)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	body, err := json.Marshal(response.Attributes)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	d := deps{}

	lambda.Start(d.handler)
}
