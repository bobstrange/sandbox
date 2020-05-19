import { APIGatewayProxyHandler } from 'aws-lambda';
import 'source-map-support/register';

const createAuction: APIGatewayProxyHandler = async (event, _context) => {
  const { title } = JSON.parse(event.body)
  const now = new Date()

  const auction = {
    title,
    status: 'OPEN',
    createdAt: now.toISOString()
  }

  return {
    statusCode: 201,
    body: JSON.stringify(auction)
  }
}
export const handler = createAuction
