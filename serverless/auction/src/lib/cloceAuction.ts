import { DynamoDB } from "aws-sdk";
import { DocumentClient } from "aws-sdk/clients/dynamodb";

const dynamodb = new DynamoDB.DocumentClient()

export async function closeAuction(auction) {
  const params: DocumentClient.UpdateItemInput = {
    TableName: process.env.AUCTIONS_TABLE_NAME,
    Key: { id: auction.id },
    UpdateExpression: 'set #status = :status',
    ExpressionAttributeValues: {
      ':status': 'CLOSED'
    },
    ExpressionAttributeNames: {
      '#status': 'status'
    }
  }

  const result = await dynamodb.update(params).promise()
  return result
}
