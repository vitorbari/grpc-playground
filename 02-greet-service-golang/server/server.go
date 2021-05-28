package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vitorbari/grpc-playground/02-greet-service-golang/greetpb"
	"google.golang.org/grpc"
)

type server struct {}

// Greet implements GreetServiceServer
func (s server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", request)

	g := request.GetGreeting()

	fn := g.GetFirstName()
	ln := g.GetLastName()

	r := greetpb.GreetResponse{
		Result: fmt.Sprintf("Hello %s %s", fn, ln),
	}

	return &r, nil
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
