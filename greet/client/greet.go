package main

import (
	"context"
	pb "github.com/korvised/go-microservice-grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Korvised",
	})
	if err != nil {
		log.Printf("Could not greet: %v", err)
	}

	log.Printf("Greet response: %v", res.Result)
}
