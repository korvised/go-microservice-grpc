package main

import (
	"context"
	pb "github.com/korvised/go-microservice-grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with request: %v", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName + "!",
	}, nil
}
