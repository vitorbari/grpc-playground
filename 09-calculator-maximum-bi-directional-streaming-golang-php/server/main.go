package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/vitorbari/grpc-playground/09-calculator-maximum-bi-directional-streaming-golang-php/calculator"
	"google.golang.org/grpc"
)

type server struct{}

// Average implements CalculatorServiceServer
func (s *server) Maximum(stream calculator.CalculatorService_MaximumServer) error {
	fmt.Println("Maximum function was invoked with a streaming request")

	max := int32(0)

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		if request.GetNumber() > max {
			max = request.GetNumber()

			err = stream.Send(&calculator.MaximumResponse{
				Maximum: max,
			})
			if err != nil {
				log.Fatalf("Error while sending data to client %v", err)
			}
		}
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
