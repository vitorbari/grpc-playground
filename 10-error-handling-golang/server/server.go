package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/vitorbari/grpc-playground/10-error-handling-golang/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {}

// SquareRoot implements CalculatorServiceServer
func (s server) SquareRoot(c context.Context, request *calculator.SquareRootRequest) (*calculator.SquareRootResponse, error) {
	fmt.Printf("SquareRoot function was invoked with %v\n", request)

	rad := request.GetRadicand()

	if rad < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Radicand must be a positive number, got: %d", rad)
	}

	return &calculator.SquareRootResponse{
		Result: math.Sqrt(float64(rad)),
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
