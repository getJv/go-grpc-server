package services

import (
	"context"
	"fmt"
	"getjv.github.com/go-grpc-server/generated/protos"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

type GreeterSrvImpl struct {
	protos.UnimplementedGreeterServer
}

func RegisterGreeterSrv(grpcServer *grpc.Server) {
	protos.RegisterGreeterServer(grpcServer, &GreeterSrvImpl{})
}

func (s *GreeterSrvImpl) SayHello(_ context.Context, in *protos.HelloRequest) (*protos.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &protos.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *GreeterSrvImpl) StreamGreetings(req *protos.HelloRequest, stream protos.Greeter_StreamGreetingsServer) error {
	log.Printf("StreamGreetings starting to: %s", req.GetName())
	for i := 1; i <= 10; i++ {
		response := &protos.HelloReply{
			Message: fmt.Sprintf("Hello %s, msg number: %d", req.GetName(), i),
		}

		if err := stream.Send(response); err != nil {
			log.Printf("Erro while sending message: %v", err)
			return err
		}

		time.Sleep(1 * time.Second)
	}

	log.Println("StreamGreetings done.")
	return nil
}

func (s *GreeterSrvImpl) EchoGreetings(stream protos.Greeter_EchoGreetingsServer) error {
	log.Println("EchoGreetings started...")

	var allNames string // Stores names received from the client

	for {
		// Receive the next message from the stream
		req, err := stream.Recv()
		if err == io.EOF {
			// Stream has been closed by the client: Send a consolidated single response
			reply := &protos.HelloReply{
				Message: fmt.Sprintf("Hello to everyone: %s", allNames),
			}
			return stream.SendAndClose(reply)
		}
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return err // Terminate on error
		}

		// Process the received message (append name to the list)
		log.Printf("Received: %s", req.GetName())
		if allNames == "" {
			allNames = req.GetName()
		} else {
			allNames = fmt.Sprintf("%s, %s", allNames, req.GetName())
		}
	}

}

func (s *GreeterSrvImpl) OpenGreetings(stream protos.Greeter_OpenGreetingsServer) error {
	log.Println("OpenGreetings started...")

	for {
		// Receive the next message from the client stream
		req, err := stream.Recv()
		if err == io.EOF {
			// Stream has been closed by the client
			log.Println("OpenGreetings: Client closed the stream.")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return err
		}

		// Process the received message and generate a reply
		log.Printf("Received: %s", req.GetName())
		reply := &protos.HelloReply{
			Message: fmt.Sprintf("Hello %s! Your greeting is acknowledged.", req.GetName()),
		}

		// Send the reply back to the client
		if err := stream.Send(reply); err != nil {
			log.Printf("Error sending reply: %v", err)
			return err
		}

		log.Printf("Reply sent: %s", reply.Message)
	}
}
