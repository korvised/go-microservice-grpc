package main

import (
	"context"
	pb "github.com/korvised/go-microservice-grpc/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	// Sum
	sumRequest := &pb.SumRequest{
		A: 3,
		B: 10,
	}

	sumResponse, err := c.Sum(context.Background(), sumRequest)
	if err != nil {
		log.Fatalf("Failed to call Sum: %v", err)
	}
	log.Printf("Sum Result: %v", sumResponse.GetResult())
}
