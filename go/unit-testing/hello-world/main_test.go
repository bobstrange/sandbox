package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockedPutItem struct {
	dynamodbiface.DynamoDBAPI
	Response dynamodb.PutItemOutput
}

func (m mockedPutItem) PutItem(in *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &m.Response, nil
}

func TestHandler(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		m := mockedPutItem{
			Response: dynamodb.PutItemOutput{},
		}
		d := deps{
			ddb:   m,
			table: "SomeTable",
		}
		_, err := d.handler(events.APIGatewayProxyRequest{})
		if err != nil {
			t.Fatal("Error")
		}
	})
}
