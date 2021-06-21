package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/vitorbari/grpc-playground/07-calculator-average-client-streaming-golang-php/calculator"
	"google.golang.org/grpc"
)

type server struct{}

// Average implements CalculatorServiceServer
func (s *server) Average(stream calculator.CalculatorService_AverageServer) error {
	fmt.Println("Average function was invoked with a streaming request")

	numbers := []int32{}

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			// finished reading the client stream
			return stream.SendAndClose(&calculator.AverageResponse{
				Result: getResult(numbers),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		numbers = append(numbers, request.GetNumber())
	}
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

func getResult(n []int32) float64 {
	s := int32(0)
	for _, v := range n {
		s += v
	}
	return float64(s) / float64(len(n))
}
