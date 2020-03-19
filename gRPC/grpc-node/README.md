# gRPC NodeJS

## Setup

```bash
npm install --save grpc-tools google-protobuf
```

## Run server
```bash
npm run server:start
npm run client:start
```

## Make a example

```bash
npx grpc_tools_node_protoc -I=. ./proto/dummy.proto \
  --js_out=import_style=commonjs,binary:./server \
  --grpc_out=./server \
  --plugin=protoc-gen-grpc=node_modules/.bin/grpc_tools_node_protoc_plugin
```


