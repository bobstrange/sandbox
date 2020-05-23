import { DynamoDB } from "aws-sdk";

const dynamodb = new DynamoDB.DocumentClient()

export async function getEndedAuctions() {
  const now = new Date()
  const result = await dynamodb.query({
    TableName: process.env.AUCTIONS_TABLE_NAME,
    IndexName: 'statusAndEndDate',
    KeyConditionExpression: '#status = :status AND endingAt <= :now',
    ExpressionAttributeValues: {
      ':status': 'OPEN',
      ':now': now.toISOString()
    },
    // statusが予約語なので、、、
    ExpressionAttributeNames: {
      '#status': 'status'
    }
  }).promise()
  return result.Items
}
