package main

import (
	pb "github.com/korvised/go-microservice-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	//doSum(c)
	doPrimes(c)
}
