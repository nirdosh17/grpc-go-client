package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
)

func longGreet(c pb.GreetServiceClient) {
	log.Println("calling rpc endpoint LongGreet...")

	reqs := []*pb.GreetRequest{
		{FirstName: "Nirdosh1"},
		{FirstName: "Nirdosh2"},
		{FirstName: "Nirdosh3"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalln("Error calling LongGreet:", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("error receiving response from LongGreet:", err)
	}

	log.Println("LongGreet rpc response:", res.Result)
}
