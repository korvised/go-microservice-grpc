package main

import (
	"fmt"
	pb "github.com/korvised/go-microservice-grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request")

	result := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: result,
			})
		}
		log.Printf("Received request: %v\n", req)

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		result += fmt.Sprintf("Hello %s! \n", req.GetFirstName())
	}

	return nil
}
