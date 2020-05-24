import { APIGatewayProxyEvent, Handler, APIGatewayProxyResult } from 'aws-lambda';
import 'source-map-support/register';
import { v4 as uuid } from 'uuid'
import { DynamoDB } from 'aws-sdk'
import commonMiddleware from "../lib/commonMiddleware";
import createError from 'http-errors'
import { Assign } from 'utility-types'
import validator from '@middy/validator';

type SetBodyToType<A extends object, B extends object> = Assign<A, Record<"body", B>>

const dynamodb = new DynamoDB.DocumentClient()

const createAuction: Handler<
  SetBodyToType<APIGatewayProxyEvent, { title?: string }>,
  APIGatewayProxyResult
> = async (event, _context) => {
  const { title } = event.body

  if (!title) {
    throw new createError.BadRequest(`${title} is required`)
  }

  const now = new Date()
  const endDate = new Date()
  endDate.setHours(now.getHours() + 1)

  const auction = {
    id: uuid(),
    title,
    status: 'OPEN',
    createdAt: now.toISOString(),
    endingAt: endDate.toISOString(),
    highestBid: {
      amount: 0
    }
  }

  try {
    await dynamodb.put({
      TableName: process.env.AUCTIONS_TABLE_NAME,
      Item: auction
    }).promise()
  } catch (error) {
    console.error(error)
    throw new createError.InternalServerError(error)
  }

  return {
    statusCode: 201,
    body: JSON.stringify(auction)
  }
}

const createAuctionSchema = {
  properties: {
    body: {
      type: "object",
      properties: {
        title: {
          type: "string"
        },
      },
      required: ["title"]
    },
  },
  required: ["body"],
}

export const handler = commonMiddleware(createAuction)
  .use(
    validator({
      inputSchema: createAuctionSchema,
    })
  )
