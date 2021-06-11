package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/vitorbari/grpc-playground/04-greet-server-streaming-golang/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

// Greet implements GreetServiceServer
func (s server) Greet(request *greetpb.GreetRequest, stream greetpb.GreetService_GreetServer) error {
	fmt.Printf("Greet function was invoked with %v\n", request)

	g := request.GetGreeting()

	fn := g.GetFirstName()
	ln := g.GetLastName()

	for i := 0; i < 10; i++ {
		r := &greetpb.GreetManyTimesResponse{
			Result: fmt.Sprintf("%d - Hello %s %s", i, fn, ln),
		}
		stream.Send(r)
		time.Sleep(1000 * time.Millisecond)
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
