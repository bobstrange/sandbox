import { APIGatewayProxyEvent, Handler, APIGatewayProxyResult } from 'aws-lambda';
import "source-map-support/register";
import { DynamoDB } from "aws-sdk";
import commonMiddleware from "../lib/commonMiddleware";
import createError from "http-errors";
import { Assign } from 'utility-types'

type SetBodyToType<A extends object, B extends object> = Assign<A, Record<"body", B>>

const dynamodb = new DynamoDB.DocumentClient();

const placeBid: Handler<
  SetBodyToType<APIGatewayProxyEvent, { amount?: number }>,
  APIGatewayProxyResult
> = async (event, _context) => {
  const { id } = event.pathParameters
  const { amount } = event.body

  let updatedAuction

  try {
    const result = await dynamodb.update({
      TableName: process.env.AUCTIONS_TABLE_NAME,
      Key: { id },
      UpdateExpression: 'set highestBid.amount = :amount',
      ExpressionAttributeValues: {
        ':amount': amount
      },
      ReturnValues: 'ALL_NEW'
    }).promise()
    updatedAuction = result.Attributes
  } catch (error) {
    console.error(error);
    throw new createError.InternalServerError(error);
  }

  return {
    statusCode: 200,
    body: JSON.stringify(updatedAuction),
  };
};

export const handler = commonMiddleware(placeBid);
