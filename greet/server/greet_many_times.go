package main

import (
	"fmt"
	pb "github.com/korvised/go-microservice-grpc/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, steam pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.GetFirstName(), i)

		steam.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
