package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vitorbari/grpc-playground/10-error-handling-golang/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Starting client...")

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := calculator.NewCalculatorServiceClient(conn)

	ctx := context.Background()

	radicands := []int{16, 9, 1231, -4, 25}

	for _, v := range radicands {
		getSquareRoot(ctx, c, v)
	}
}

func getSquareRoot(ctx context.Context, c calculator.CalculatorServiceClient, rad int) {
	res, err := c.SquareRoot(ctx, &calculator.SquareRootRequest{
		Radicand: int32(rad),
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		log.Printf("Error message from server: %v\n", errStatus.Message())
		log.Printf("Error code from server: %v\n", errStatus.Code())

		if codes.InvalidArgument == errStatus.Code() {
			log.Println("Sent a invalid argument")
			return
		}

		log.Fatalf("some other error while calling Calculator RPC: %v", err)
	}

	log.Printf("Sqrt of %d is: %f", rad, res.Result)
}
