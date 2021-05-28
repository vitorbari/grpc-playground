package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vitorbari/grpc-playground/02-greet-service-golang/greetpb"
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

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet RPC: %v", res.Result)
}
