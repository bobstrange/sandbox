import { APIGatewayProxyEvent, Handler, APIGatewayProxyResult } from 'aws-lambda';
import "source-map-support/register";
import { DynamoDB } from "aws-sdk";
import commonMiddleware from "../lib/commonMiddleware";
import createError from "http-errors";
import { Assign } from 'utility-types'
import { getAuctionById } from "./getAuction";
import validator from '@middy/validator';

type SetBodyToType<A extends object, B extends object> = Assign<A, Record<"body", B>>

const dynamodb = new DynamoDB.DocumentClient();

const placeBid: Handler<
  SetBodyToType<APIGatewayProxyEvent, { amount?: number }>,
  APIGatewayProxyResult
> = async (event, _context) => {
  const { id } = event.pathParameters
  const { amount } = event.body
  const { email } = event.requestContext.authorizer

  const auction = await getAuctionById(id);

  if (auction.status !== 'OPEN') {
    throw new createError.Forbidden('You cannot bid on closed auctions!')
  }

  const highestBidAmount = auction.highestBid.amount
  if (amount <= auction.highestBid.amount) {
    throw new createError.Forbidden(`Your bid must be higher than ${highestBidAmount}`)
  }

  let updatedAuction

  try {
    const result = await dynamodb.update({
      TableName: process.env.AUCTIONS_TABLE_NAME,
      Key: { id },
      UpdateExpression: 'set highestBid.amount = :amount, highestBid.bidder = :bidder',
      ExpressionAttributeValues: {
        ':amount': amount,
        ':bidder': email
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

const placeBidSchema = {
  properties: {
    body: {
      type: 'object',
      properties: {
        amount: {
          type: 'number',
        }
      },
      required: ['amount']
    }
  },
  required: ['body']
}

export const handler = commonMiddleware(placeBid)
  .use(
    validator({
      inputSchema: placeBidSchema
    })
  )

