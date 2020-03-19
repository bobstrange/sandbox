const grpc = require('grpc')

const greets = require("../server/proto/greet_pb");
const service = require("../server/proto/greet_grpc_pb");
const sums = require("../server/proto/sum_pb");
const sumService = require("../server/proto/sum_grpc_pb");

function callGreet() {
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
}

function callGreetManyTimes() {
  const client = new service.GreetServiceClient(
    'localhost:50051',
    grpc.credentials.createInsecure()
  )

  const request = new greets.GreetManyTimesRequest()
  const greeting = new greets.Greeting()
  greeting.setFirstName('John')
  greeting.setLastName('Doe')
  request.setGreeting(greeting)

  const call = client.greetManyTimes(request, () => {})

  call.on('data', (response) => {
    console.log('Client Streaming Response: ', response.getResult())
  })

  call.on('status', (status) => {
    console.log(status)
  })

  call.on('error', (error) => {
    console.error(error)
  })

  call.on('end', () => {
    console.log('Client streaming ended')
  })
}

function callSum() {
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

function main() {
  // callGreet()
  callGreetManyTimes()
  // callSum()
}

main()
