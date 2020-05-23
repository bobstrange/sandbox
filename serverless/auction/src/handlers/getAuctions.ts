import {
  APIGatewayProxyHandler,
} from "aws-lambda";
import "source-map-support/register";
import { DynamoDB } from "aws-sdk";
import commonMiddleware from "../lib/commonMiddleware";
import createError from "http-errors";
import { DocumentClient } from "aws-sdk/clients/dynamodb";
import validator from '@middy/validator'

const dynamodb = new DynamoDB.DocumentClient();



const getAuctions: APIGatewayProxyHandler = async (event, _context) => {
  const { status } = event.queryStringParameters

  let auctions

  const params: DocumentClient.QueryInput = {
    TableName: process.env.AUCTIONS_TABLE_NAME,
    IndexName: "statusAndEndDate",
    KeyConditionExpression: "#status = :status",
    ExpressionAttributeValues: {
      ":status": status,
    },
    ExpressionAttributeNames: {
      "#status": "status",
    },
  };

  try {
    const result = await dynamodb.query(params).promise()
    auctions = result.Items
  } catch (error) {
    console.error(error)
    throw new createError.InternalServerError(error)
  }

  return {
    statusCode: 200,
    body: JSON.stringify(auctions),
  };
};

const getAuctionsSchema = {
  properties: {
    queryStringParameters: {
      type: 'object',
      properties: {
        status: {
          type: 'string',
          enum: ['OPEN', 'CLOSED'],
          default: 'OPEN'
        }
      }
    }
  },
  required: [
    'queryStringParameters'
  ]
}

export const handler = commonMiddleware(getAuctions)
  .use(
    validator({
      inputSchema: getAuctionsSchema,
      ajvOptions: {
        useDefaults: true
      }
    })
  )
