const grpc = require('grpc')

const greets = require("../server/proto/greet_pb");
const service = require("../server/proto/greet_grpc_pb");
const sums = require("../server/proto/sum_pb");
const sumService = require("../server/proto/sum_grpc_pb");

function main() {
  const client = new service.GreetServiceClient(
    'localhost:50051',
    grpc.credentials.createInsecure()
  )

  const request = new greets.GreetRequest()
  const greeting = new greets.Greeting()
  greeting.setFirstName('John')
  greeting.setLastName('Doe')
  request.setGreeting(greeting)

  client.greet(request, (error, response) => {
    if (!error) {
      console.log('Greeting Response: ', response.getResult())
    } else {
      console.error(error)
    }
  })

  const sumRequest = new sums.SumRequest()
  const summing = new sums.Summing()
  summing.setFirstValue(100)
  summing.setSecondValue(11)
  sumRequest.setSumming(summing)
  const sumClient = new sumService.SumServiceClient(
    'localhost:50051',
    grpc.credentials.createInsecure()
  )
  sumClient.sum(sumRequest, (error, response) => {
    if (!error) {
      console.log('Sum Response: ', response.getResult())
    } else {
      console.error(error)
    }
  })
}

main()
