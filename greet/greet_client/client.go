package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/onkarsutar/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello From Client")

	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "onkar",
			LastName:  "sutar",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call Greet RPS %v\n", err)
	}

	log.Printf("Response from Greet %v\n", res.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "onkar",
			LastName:  "sutar",
		},
	}

	res, err := c.GreetMany(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call GreetMany %v", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		log.Printf("Response from server %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	log.Println("Starting Client Streaming")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet %v", err)
	}

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "onkar",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "kishan",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "viraj",
			},
		},
	}

	for _, req := range requests {
		log.Println("Sending ", req.Greeting.GetFirstName())
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response %v", err)
	}

	log.Println(res.GetResult())
}
