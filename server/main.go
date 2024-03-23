package main

import (
	"log"
	"net"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

// define the port
const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	//create a grpc server
	grpcServer := grpc.NewServer()

	//Register the greetservice
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("Server started at %v", lis.Addr())

	//lis is the port, the grpc server need to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
