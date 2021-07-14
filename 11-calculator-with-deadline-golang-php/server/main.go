package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/vitorbari/grpc-playground/11-calculator-with-deadline-golang-php/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

// Factorial implements CalculatorService
func (s server) Factorial(ctx context.Context, request *calculator.FactorialRequest) (*calculator.FactorialResponse, error) {
	fmt.Printf("Factorial function was invoked with %v\n", request)

	n := uint64(request.GetNumber())

	// Implement a slow factorial algorithm
	var r uint64 = 1
	for i := uint64(1); i <= n; i++ {
		if ctx.Err() == context.Canceled {
			return nil, status.Error(codes.Canceled, "Client cancelled, abandoning.")
		}
		// Need to make it even slower
		time.Sleep(time.Millisecond * 300)
		r *= i
	}

	return &calculator.FactorialResponse{
		Result: r,
	}, nil
}

func main() {
	fmt.Println("Starting server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
