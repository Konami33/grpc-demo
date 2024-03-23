package main

import (
	"log"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	//connection close
	defer conn.Close()

	//creating a new client
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"A", "B", "C"},
	}

	//callSayHello(client) //unart api

	//callSayHelloServerStream(client, names) //serverstreaming api

	//callSayHelloClientStream(client, names) //clientStreaming api

	callSayHelloBidirectionalStream(client, names)




}