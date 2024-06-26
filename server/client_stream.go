package main

import (
	"io"
	"log"

	pb "grpc-demo/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Names: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got stream request with name : %v", req.Name)
		messages = append(messages, "Hello " + req.Name)
	}
}