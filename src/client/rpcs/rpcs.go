package rpcs

import (
	"grpc-go-example/src/pb"
	"log"
	"grpc-go-example/src"
	"google.golang.org/grpc/metadata"
	"io"
	"golang.org/x/net/context"
)

var fooBarBazs = []pb.FooBarBaz{
	{Foo: "foo", Bar: 1, Baz: true},
	{Foo: "bar", Bar: 2, Baz: false},
	{Foo: "baz", Bar: 3, Baz: true},
}
var fooBarBaz = fooBarBazs[0]

func Unary(client pb.FooBarBazServiceClient) () {
	log.Printf("Unary: Interaction started.\n")
	defer log.Printf("Unary: Interaction complete.\n")

	util.SimulateProcessing()
	requestMetadata := metadata.New(map[string]string{
		"foo": "bar",
		"bar": "baz",
	})
	requestContext := metadata.NewOutgoingContext(context.Background(), requestMetadata)
	log.Printf("Unary: Sending request to server.\n")
	unaryResponse, err := client.Unary(requestContext, &pb.UnaryRequest{FooBarBaz: &fooBarBaz})
	if err != nil {
		log.Fatalf("Unary: Request Error: \"%v\".\n", err)
	}
	log.Printf("Unary: Response from server: \"%v\".\n", unaryResponse)
}

func ServerStream(client pb.FooBarBazServiceClient) () {
	log.Printf("ServerStream: Interaction started.\n")
	defer log.Printf("ServerStream: Interaction complete.\n")

	util.SimulateProcessing()
	requestMetadata := metadata.New(map[string]string{
		"foo": "bar",
		"bar": "baz",
	})
	requestContext := metadata.NewOutgoingContext(context.Background(), requestMetadata)
	log.Printf("ServerStream: Sending request to server.\n")
	stream, err := client.ServerStream(requestContext, &pb.ServerStreamRequest{FooBarBaz: &fooBarBaz})
	if err != nil {
		log.Fatalf("ServerStream: Request error: \"%v\".\n", err)
	}

	for {
		serverStreamResponse, err := stream.Recv()
		if err == io.EOF {
			log.Printf("ServerStream: Server stream closed.\n")
			break
		}
		if err != nil {
			log.Fatalf("ServerStream: Response error: \"%v\".\n", err)
		}
		log.Printf("ServerStream: Response received: \"%v\".\n", serverStreamResponse)
	}
}

func ClientStream(client pb.FooBarBazServiceClient) () {
	log.Printf("ClientStream: Interaction started.\n")
	defer log.Printf("ClientStream: Interaction complete.\n")

	requestMetadata := metadata.New(map[string]string{
		"foo": "bar",
		"bar": "baz",
		"baz": "foo",
	})
	requestContext := metadata.NewOutgoingContext(context.Background(), requestMetadata)
	stream, err := client.ClientStream(requestContext)
	if err != nil {
		log.Fatalf("ClientStream: Stream error: \"%v\".\n", err)
	}

	for _, fooBarBaz := range fooBarBazs {
		util.SimulateProcessing()
		log.Printf("ClientStream: Sending request to server.\n")
		err := stream.Send(&pb.ClientStreamRequest{FooBarBaz: &fooBarBaz})
		if err != nil {
			log.Printf("ClientStream: Request error: \"%v\".\n", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("ClientStream: Response error: \"%v\".\n", err)
	}
	log.Printf("ClientStream: Client stream closed.\n")
	log.Printf("ClientStream: Response received: \"%v\".\n", response.Success)
}

func BidirectionalStream(client pb.FooBarBazServiceClient) () {
	log.Printf("BidirectionalStream: Interaction started.\n")
	defer log.Printf("BidirectionalStream: Interaction complete.\n")

	requestMetadata := metadata.New(map[string]string{
		"foo": "bar",
		"bar": "baz",
	})
	requestContext := metadata.NewOutgoingContext(context.Background(), requestMetadata)
	stream, err := client.BidirectionalStream(requestContext)
	if err != nil {
		log.Fatalf("BidirectionalStream: Stream error: \"%v\".\n", err)
	}

	interactionComplete := make(chan util.Void)
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				log.Printf("BidirectionalStream: Server stream closed.\n")
				interactionComplete <- util.Void{}
				break
			}
			if err != nil {
				log.Fatalf("BidirectionalStream: Response error: \"%v\".\n", err)
			}
			log.Printf("BidirectionalStream: Response received: \"%v\".\n", response)
		}
	}()

	for _, fooBarBaz := range fooBarBazs {
		util.SimulateProcessing()
		log.Printf("BidirectionalStream: Sending request to server.\n")
		err := stream.Send(&pb.BidirectionalStreamRequest{FooBarBaz: &fooBarBaz})
		if err != nil {
			log.Fatalf("BidirectionalStream: Request error: \"%v\".\n", err)
		}
	}
	err = stream.CloseSend()
	if err != nil {
		log.Printf("BidirectionalStream: Final request error: \"%v\".\n", err)
	}
	log.Printf("BidirectionalStream: Client stream closed.\n")
	<-interactionComplete
}
