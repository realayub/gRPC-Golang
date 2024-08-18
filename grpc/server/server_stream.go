package main

import (
	"log"
	"time"

	pb "github.com/realayub/grpc/proto" // Replace with your actual package path
)

// callSayHelloServerStreaming handles the server streaming RPC
func (s *helloserver) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
