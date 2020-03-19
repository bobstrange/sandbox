const grpc = require('grpc')
const greets = require('../server/proto/greet_pb')
const greetService = require('../server/proto/greet_grpc_pb')
const sums = require('../server/proto/sum_pb')
const sumService = require('../server/proto/sum_grpc_pb')

/**
 * Implements the greet RPC method
 */

function greet(call, callback) {
  const greeting = new greets.GreetResponse()
  greeting.setResult(
    `Hello ${call.request.getGreeting().getFirstName()} ${call.request.getGreeting().getLastName()}`
  )
  callback(null, greeting)
}

/**
 * Implements the sum RPC method
 */
function sum(call, callback) {
  const response = new sums.SumResponse();
  response.setResult(call.request.getSumming().getFirstValue() + call.request.getSumming().getSecondValue())
  callback(null, response)
}

function main() {
  const server = new grpc.Server()
  server.addService(greetService.GreetServiceService, { greet })
  server.addService(sumService.SumServiceService, { sum })
  server.bind("127.0.0.1:50051", grpc.ServerCredentials.createInsecure())
  server.start()

  console.log('Server running on port 127.0.0.1:50051')
}
main()
