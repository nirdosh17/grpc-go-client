package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func greetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("calling rpc endpoint GreetWithDeadline...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "Nirdosh",
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("deadline exceeded")
				return
			} else {
				log.Fatalln("unexpected grpc error", err)
			}

		} else {
			log.Fatalln("non grpc error", err)
		}
	}

	log.Println("GreetWithDeadline rpc response: ", resp.Result)
}
