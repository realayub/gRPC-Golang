package main

import (
	"io"
	"log"

	pb "github.com/realayub/grpc/proto"
)

func (s *helloserver) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// When the client has finished sending messages, send back the list and close the stream
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		// Append the message to the slice
		messages = append(messages, "Hello ", req.Name)
	}
}
