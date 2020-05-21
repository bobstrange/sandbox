import { APIGatewayProxyHandler } from "aws-lambda";
import "source-map-support/register";
import { DynamoDB } from "aws-sdk";
import commonMiddleware from '../lib/commonMiddleware'

import createError from "http-errors";
import { AttributeMap } from "aws-sdk/clients/dynamodb";

const dynamodb = new DynamoDB.DocumentClient();

export const getAuctionById = async (id: string): Promise<AttributeMap> => {
  let auction

  try {
    const result = await dynamodb.get({
      TableName: process.env.AUCTIONS_TABLE_NAME,
      Key: { id }
    }).promise()
    auction = result.Item
  } catch (error) {
    console.error(error);
    throw new createError.InternalServerError(error);
  }

  if (!auction) {
    throw new createError.NotFound(`Auction with ID: "${id}" not found`)
  }

  return auction
}

const getAuction: APIGatewayProxyHandler = async (event, _context) => {
  const { id } = event.pathParameters
  const auction = await getAuctionById(id)

  return {
    statusCode: 200,
    body: JSON.stringify(auction),
  };
};

export const handler = commonMiddleware(getAuction)
