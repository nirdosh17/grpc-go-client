package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
)

func greetManyTimes(c pb.GreetServiceClient) {
	// wait infinitely for stream of result unless we get an EOF
	log.Println("calling rpc endpoint GreetManyTimes...")

	req := &pb.GreetRequest{FirstName: "Nirdosh"}
	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error calling GreetManyTimes %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error reading stream: %v\n", err)
		}

		log.Printf("GreetManyTimes rpc response: %s\n", msg.Result)
	}
}
