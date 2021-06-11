package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vitorbari/grpc-playground/05-prime-number-decomposition-server-streaming-golang-php/calculator"
	"google.golang.org/grpc"
)

type server struct{}

// GetPrimeNumberDecomposition implements CalculatorService
func (s *server) GetPrimeNumberDecomposition(request *calculator.PrimeNumberDecompositionRequest, stream calculator.CalculatorService_GetPrimeNumberDecompositionServer) error {
	fmt.Printf("GetPrimeNumberDecomposition function was invoked with %v\n", request)

	n := request.GetNumber()
	k := int32(2)

	for n > 1 {
		if n%k == 0 {
			r := &calculator.PrimeNumberDecompositionResponse{
				Factor: k,
			}
			stream.Send(r)
			n /= k
			continue
		}

		k++
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
	calculator.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
