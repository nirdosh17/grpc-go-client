package main

import (
	"log"
	"time"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// make sure to start service from this repo: https://github.com/nirdosh17/grpc-go-service
var rpcServerEndpoint string = "0.0.0.0:5051"

func main() {
	tls := false
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalln("error while loading CA trust cert", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(rpcServerEndpoint, opts...)
	if err != nil {
		log.Fatalf("failed to connect %v: %v\n", rpcServerEndpoint, err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	greet(client)
	greetManyTimes(client)
	longGreet(client)
	greetEveryOne(client)
	// must wait more than 3 seconds
	greetWithDeadline(client, 4*time.Second)
}
