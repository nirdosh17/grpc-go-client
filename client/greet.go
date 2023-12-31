package main

import (
	"context"
	"log"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
)

func greet(c pb.GreetServiceClient) {
	log.Println("calling rpc endpoint Greet...")

	resp, err := c.Greet(context.Background(),
		&pb.GreetRequest{FirstName: "Nirdosh"},
	)

	if err != nil {
		log.Fatalln("Greet RPC failed:", err)
	}
	log.Println("Greet RPC response: ", resp.Result)
}
