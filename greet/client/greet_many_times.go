package main

import (
	"context"
	pb "github.com/korvised/go-microservice-grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Korvised",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Printf("Could not greet many times: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("GreetManyTimes response: %v", res.GetResult())
	}
}
