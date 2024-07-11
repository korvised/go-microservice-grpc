package main

import (
	"context"
	pb "github.com/korvised/go-microservice-grpc/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Sum function was invoked with: ", in)

	return &pb.SumResponse{
		Result: in.GetA() + in.GetB(),
	}, nil
}
