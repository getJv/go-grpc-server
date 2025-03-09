package main

import (
	"flag"
	"fmt"
	"getjv.github.com/go-grpc-server/services"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The GreeterSrvImpl port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	services.RegisterGreeterSrv(s)

	reflection.Register(s)

	log.Printf("GreeterSrvImpl listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
