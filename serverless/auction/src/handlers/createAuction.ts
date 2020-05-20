import { APIGatewayProxyEvent, Handler, APIGatewayProxyResult } from 'aws-lambda';
import 'source-map-support/register';
import { v4 as uuid } from 'uuid'
import { DynamoDB } from 'aws-sdk'
import middy from '@middy/core'
import httpJsonBodyParser from '@middy/http-json-body-parser'
import httpEventNormalizer from '@middy/http-event-normalizer'
import httpErrorHandler from '@middy/http-error-handler'
import createError from 'http-errors'
import { Assign } from 'utility-types'

type SetBodyToType<A extends object, B extends object> = Assign<A, Record<"body", B>>

const dynamodb = new DynamoDB.DocumentClient()

const createAuction: Handler<
  SetBodyToType<APIGatewayProxyEvent, { title: string }>,
  APIGatewayProxyResult
> = async (event, _context) => {
  const { title } = event.body
  const now = new Date()

  const auction = {
    id: uuid(),
    title,
    status: 'OPEN',
    createdAt: now.toISOString()
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
export const handler = middy(createAuction)
  .use(httpJsonBodyParser())
  .use(httpEventNormalizer())
  .use(httpErrorHandler())
