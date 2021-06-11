package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/vitorbari/grpc-playground/04-greet-server-streaming-golang/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client...")

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName:  "Days",
		},
	}

	resStream, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while readins stream: %v", err)
		}

		log.Printf("Stream Response: %v", msg.Result)
	}
}
