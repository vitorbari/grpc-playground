package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vitorbari/grpc-playground/03-calculator-service-golang-php/calculator"
	"google.golang.org/grpc"
)

type server struct{}

// Add implements CalculatorService
func (s server) Add(ctx context.Context, request *calculator.AdditionRequest) (*calculator.AdditionResponse, error) {
	fmt.Printf("Add function was invoked with %v\n", request)

	add := request.GetAddition()

	r := calculator.AdditionResponse{
		Result: add.GetAugend() + add.GetAddend(),
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
	calculator.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
