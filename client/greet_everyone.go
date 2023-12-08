package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
)

func greetEveryOne(c pb.GreetServiceClient) {
	log.Println("calling rpc endpoint GreetEveryone...")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln("error while creating stream", err)
	}

	greets := []*pb.GreetRequest{
		{FirstName: "Nirdosh1"},
		{FirstName: "Nirdosh2"},
		{FirstName: "Nirdosh3"},
	}
	waitc := make(chan struct{})

	go func() {
		for _, req := range greets {
			log.Println("sending", req.FirstName)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		// client saying that it finished sending all requests
		stream.CloseSend()
	}()

	// go routine to receive response from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("error while receiving", err)
				break
			}
			log.Println("GreetEveryone rpc response: ", res.Result)
		}
		close(waitc)
	}()

	<-waitc

}
