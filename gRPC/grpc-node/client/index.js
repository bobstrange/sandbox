const grpc = require('grpc')

const services = require('../server/proto/dummy_grpc_pb')

function main() {
  console.log('Hello from Client')
  const client = new services.DummyServiceClient(
    'localhost:50051',
    grpc.credentials.createInsecure()
  )
  console.log('Client: ', client)
}

main()
