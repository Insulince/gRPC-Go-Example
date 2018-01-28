package services

import (
	"google.golang.org/grpc/metadata"
	"io"
	"golang.org/x/net/context"
	"log"
	"grpc-go-example/src/pb"
	"grpc-go-example/src"
)

type FooBarBazService struct {
}

func (fooBarBazService *FooBarBazService) Unary(context context.Context, request *pb.UnaryRequest) (response *pb.UnaryResponse, err error) {
	log.Printf("Unary: Interaction started.\n")
	defer log.Printf("Unary: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("Unary: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Printf("Unary: Unable to read metadata!")
	}

	log.Printf("Unary: Request received: \"%v\".\n", request)

	util.SimulateProcessing()
	log.Printf("Unary: Sending response to client.\n")
	return &pb.UnaryResponse{Success: true}, nil
}

func (fooBarBazService *FooBarBazService) ServerStream(request *pb.ServerStreamRequest, stream pb.FooBarBazService_ServerStreamServer) (err error) {
	log.Printf("ServerStream: Interaction started.\n")
	defer log.Printf("ServerStream: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Printf("ServerStream: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Printf("ServerStream: Unable to read metadata!")
	}

	log.Printf("ServerStream: Request received: \"%v\".\n", request)

	for i := 0; i < 3; i++ {
		util.SimulateProcessing()
		log.Printf("ServerStream: Sending response to client.\n")
		stream.Send(&pb.ServerStreamResponse{Success: true})
	}

	return nil
}

func (fooBarBazService *FooBarBazService) ClientStream(stream pb.FooBarBazService_ClientStreamServer) (err error) {
	log.Printf("ClientStream: Interaction started.\n")
	defer log.Printf("ClientStream: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Printf("ClientStream: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Printf("ClientStream: Unable to read metadata!")
	}

	for {
		clientStreamRequest, err := stream.Recv()
		if err == io.EOF {
			log.Printf("ClientStream: Client stream closed.\n")
			break
		}
		if err != nil {
			log.Printf("ClientStream: Request error: \"%v\".\n", err)
			return stream.SendAndClose(&pb.ClientStreamResponse{Success: false})
		}
		log.Printf("ClientStream: Request received: \"%v\".\n", clientStreamRequest)
	}

	util.SimulateProcessing()
	log.Printf("ClientStream: Sending response to client.\n")
	return stream.SendAndClose(&pb.ClientStreamResponse{Success: true})
}

func (fooBarBazService *FooBarBazService) BidirectionalStream(stream pb.FooBarBazService_BidirectionalStreamServer) (err error) {
	log.Printf("BidirectionalStream: Interaction started.\n")
	defer log.Printf("BidirectionalStream: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Printf("BidirectionalStream: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Printf("BidirectionalStream: Unable to read metadata!")
	}

	for {
		bidirectionalStreamRequest, err := stream.Recv()
		if err == io.EOF {
			log.Printf("BidirectionalStream: Client stream closed.\n")
			break
		}
		if err != nil {
			log.Printf("BidirectionalStream: Request error: \"%v\".\n", err)
			return stream.Send(&pb.BidirectionalStreamResponse{Success: false})
		}
		log.Printf("BidirectionalStream: Request received: \"%v\".\n", bidirectionalStreamRequest)

		util.SimulateProcessing()
		log.Printf("BidirectionalStream: Sending response to client.\n")
		stream.Send(&pb.BidirectionalStreamResponse{Success: true})
	}

	return nil
}
