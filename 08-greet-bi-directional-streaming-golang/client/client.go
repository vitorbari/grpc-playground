package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/vitorbari/grpc-playground/08-greet-bi-directional-streaming-golang/greetpb"
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

	requests := []*greetpb.GreetRequest{
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

	wait := make(chan struct{})

	// send a bunch of messages
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending req: %v\n", req)

			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Failed to send request %v", err)
			}

			time.Sleep(1000 * time.Millisecond)
		}

		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Failed to close and send stream %v", err)
		}
	}()

	// receive a bunch of messages
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(wait)
	}()

	<-wait
}
