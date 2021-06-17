package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/vitorbari/grpc-playground/06-greet-client-streaming-golang/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

// Greet implements GreetServiceServer
func (s server) Greet(stream greetpb.GreetService_GreetServer) error {
	fmt.Println("Greet function was invoked with a streaming request")

	result := ""

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			// finished reading the client stream
			return stream.SendAndClose(&greetpb.GreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client streamL %v", err)
		}

		fn := request.GetGreeting().GetFirstName()

		result += fmt.Sprintf("Hello %s\n", fn)
	}
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
