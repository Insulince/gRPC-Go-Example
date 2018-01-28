package main

import (
	"google.golang.org/grpc"
	"log"
	"grpc-go-example/src"
	"os"
	"grpc-go-example/src/pb"
	"grpc-go-example/src/client/rpcs"
)

var requestTypes = []string{"unary", "server-stream", "client-stream", "bidirectional-stream"}
var requestType string

func init() () {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		log.Fatalf("No request type argument provided, expected one of \"%v\".\n", requestTypes)
	}
	requestType = arguments[0]
}

func main() () {
	connection, err := grpc.Dial("localhost:"+util.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC connection error: \"%v\".\n", err)
	}
	defer connection.Close()

	client := pb.NewFooBarBazServiceClient(connection)

	switch requestType {
	case requestTypes[0]:
		rpcs.Unary(client)
	case requestTypes[1]:
		rpcs.ServerStream(client)
	case requestTypes[2]:
		rpcs.ClientStream(client)
	case requestTypes[3]:
		rpcs.BidirectionalStream(client)
	default:
		log.Fatalf("Unrecognized request type: \"%v\".\n", requestType)
	}
}
