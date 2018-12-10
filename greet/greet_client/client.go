package main

import (
	"fmt"
	"log"

	"github.com/chhsu0222/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	// Create a gRPC channel to communicate with the server
	cc, err := grpc.Dial("localhos:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer cc.Close()

	// Create a client stub to perform RPCs
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}