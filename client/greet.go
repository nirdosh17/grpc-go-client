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
		log.Fatalln("err from RPC call:", err)
	}
	log.Println("Greet rpc response: ", resp.Result)
}
