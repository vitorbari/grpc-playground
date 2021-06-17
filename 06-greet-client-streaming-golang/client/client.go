package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vitorbari/grpc-playground/06-greet-client-streaming-golang/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client streaming...")

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	stream, err := c.Greet(context.Background())

	if err != nil {
		log.Fatalf("error while calling Greet: %v", err)
	}

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
				LastName:  "Days",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Bob",
				LastName:  "Bar",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Alice",
				LastName:  "Ger",
			},
		},
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)

		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Failed to send request %v", err)
		}

		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response %v", err)
	}

	fmt.Printf("Response: %s\n", res.Result)
}
