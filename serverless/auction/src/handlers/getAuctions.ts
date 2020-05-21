import {
  APIGatewayProxyHandler,
} from "aws-lambda";
import "source-map-support/register";
import { DynamoDB } from "aws-sdk";
import commonMiddleware from "../lib/commonMiddleware";
import createError from "http-errors";

const dynamodb = new DynamoDB.DocumentClient();

const getAuctions: APIGatewayProxyHandler = async (_event, _context) => {
  let auctions

  try {
    const result = await dynamodb.scan({
      TableName: process.env.AUCTIONS_TABLE_NAME
    }).promise()

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

export const handler = commonMiddleware(getAuctions)
