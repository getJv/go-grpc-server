PROTO_FOLDER_PATH=protos
GENERATED_FOLDER_NAME=generated
GO_GRPC_SERVER_PATH=${PWD}
GO_GRPC_SERVER_BINARY=go-grpc-server

go_proto_generate:
	@echo "Generating pb and grpc files for go..."
	mkdir -p ${GO_GRPC_SERVER_PATH}/${GENERATED_FOLDER_NAME} && \
    protoc --go_out=paths=source_relative:${GO_GRPC_SERVER_PATH}/${GENERATED_FOLDER_NAME} \
           --go-grpc_out=paths=source_relative:${GO_GRPC_SERVER_PATH}/${GENERATED_FOLDER_NAME} \
           ${PROTO_FOLDER_PATH}/*.proto
	@echo "Done!"

build_go_grpc_server: go_proto_generate
	@echo "Building build_go_grpc_server binary..."
	cd ${GO_GRPC_SERVER_PATH} && env GOOS=linux CGO_ENABLED=0 go build -o ${GO_GRPC_SERVER_BINARY} ./cmd/api
	@echo "Done!"

start_grpc_server: build_go_grpc_server
	@echo "Starting GRPC server..."
	cd ${GO_GRPC_SERVER_PATH} && go run ./cmd/api/main.go

build_docker: build_go_grpc_server
	@echo "Building docker image with new binary version"
	docker build -f ./go-grpc-server.dockerfile -t  getjv/go-grpc-server .
	docker images | grep getjv/go-grpc-server
	@echo "Done!"

