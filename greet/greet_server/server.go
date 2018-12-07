package main

import (
	"log"
	"net"

	"github.com/chhsu0222/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type sever struct{}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &sever{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
