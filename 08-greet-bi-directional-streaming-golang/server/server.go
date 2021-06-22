package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/vitorbari/grpc-playground/08-greet-bi-directional-streaming-golang/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

// Greet implements GreetServiceServer
func (s server) Greet(stream greetpb.GreetService_GreetServer) error {
	fmt.Println("Greet function was invoked with a streaming request")

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			// finished reading the client stream
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		result := fmt.Sprintf("Hello %s\n", request.GetGreeting().GetFirstName())

		err = stream.Send(&greetpb.GreetResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client %v", err)
		}
	}

	return nil
}

func main() {
	fmt.Println("Starting server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
