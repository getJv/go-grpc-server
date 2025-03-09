# go-grpc-server

This is a GRPC server with basic features to support my laboratories, written articles, and videos
it provides a sample of unary call, client streaming, server streaming and bidirectional streaming.
All the samples are "Hello world like" but once you have the project structure is possible to extend
it to any purpose.  

visit my [portfolio][] to find more about many topics.

[portfolio]: https://jpmorais.com.br

---

# Pre requisites
* Linux / WSL
* Docker
* Make
* [Protoc](https://grpc.io/docs/protoc-installation/)
* [grpcURL](https://github.com/fullstorydev/grpcurl)

# Available Make commands

* `make go_proto_generate`: it reads the protobuff files and gen the types and services
* `make build_go_grpc_server`: it generates the linux bin og the grpc server
* `make start_grpc_server`: it starts the server and listen to connections
* `make build_docker`: it build the image with the grpc server binary

## Docker image availiable

There is a [docker images](https://hub.docker.com/repository/docker/getjv/go-grpc-server/general) available build with the last app version for use:
```bash
docker run --rm --name grpc -p 50051:50051 getjv/go-grpc-server
```

## Exploring server calls with grpcURL

gRPCurl is a super handy tool, and the main commands are:

* `grpcurl -plaintext [::]:50051 list` – lists the available services.
* `grpcurl -plaintext [::]:50051 describe <ServiceName>` – provides more information about the available service.
* `grpcurl -plaintext -d '<JSON-Payload>' [::]:50051 <ServiceName.MethodName>` – makes a direct call to a method with a payload.

## Call samples:

*Hello sample:* it's a unary call with single request single response.
```bash
grpcurl -plaintext -d '{"name": "jhonatan"}' [::]:50051 helloworld.Greeter.SayHello
```

*Server Streaming sample:* it's a server streaming sample with 1 request and 10 messages from server 
```bash
grpcurl -d '{"name": "Cliente"}' -plaintext localhost:50051 helloworld.Greeter/StreamGreetings
```

*Client Streaming sample:* it's a client streaming sample with 10 client requests and 1 message from server
Here the steps are a bit different:
1. Open the channel with: 
```bash
grpcurl -plaintext -d @ localhost:50051 helloworld.Greeter/EchoGreetings
```
2. the terminal will wait the input give the objects bellow and press `enter`:
```bash
{"name": "Alice"}
{"name": "Bob"}
{"name": "Charlie"}
```
3. Now press `CTRL+D` to notify all was sent. Server will send you the final confirmation.

*Bi-directional streaming:* this is the most advanced sample an open channel to both sides keeping interacting
to test execute:
1. Open the channel with: 
```bash
grpcurl -plaintext -d @ localhost:50051 helloworld.Greeter/OpenGreetings
```
2. the terminal will wait the input give the objects bellow and press `enter`:
```bash
{"name": "Alice"}
```
3. After each `enter` Server will send you a confirmation.
4. Now press `CTRL+D` to notify you are done. 








