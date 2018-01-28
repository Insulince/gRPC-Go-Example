package main

import (
	"net"
	"google.golang.org/grpc"
	"fmt"
	"os"
	"log"
	"grpc-go-example/src/pb"
	"grpc-go-example/src/server/services"
	"grpc-go-example/src"
)


func main() () {
	listen, err := net.Listen("tcp", ":"+util.Port)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error opening tcp port \"%v\": \"%v\".\n", util.Port, err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFooBarBazServiceServer(grpcServer, new(services.FooBarBazService))
	log.Printf("Server started on TCP port \"%v\".\n", util.Port)
	grpcServer.Serve(listen)
}
