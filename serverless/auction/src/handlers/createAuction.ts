import { APIGatewayProxyHandler } from 'aws-lambda';
import 'source-map-support/register';
import { v4 as uuid } from 'uuid'
import { DynamoDB } from 'aws-sdk'

const dynamodb = new DynamoDB.DocumentClient()

const createAuction: APIGatewayProxyHandler = async (event, _context) => {
  const { title } = JSON.parse(event.body)
  const now = new Date()

  const auction = {
    id: uuid(),
    title,
    status: 'OPEN',
    createdAt: now.toISOString()
  }

  await dynamodb.put({
    TableName: process.env.AUCTIONS_TABLE_NAME,
    Item: auction
  }).promise()

  return {
    statusCode: 201,
    body: JSON.stringify(auction)
  }
}
export const handler = createAuction
