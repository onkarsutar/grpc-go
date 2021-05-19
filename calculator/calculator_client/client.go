package main

import (
	"context"
	"io"
	"log"

	"github.com/onkarsutar/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	// doUnary(c)

	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		Num1: 100,
		Num2: 15,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call Sum RPC %v\n", err)
	}

	log.Printf("Response from Sum %v\n", res.Sum)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.FibRequest{
		Number: 20,
	}
	res, err := c.FibSeries(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call FibSeries")
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		log.Printf("Message from server %v ", msg.GetFib())
	}
}
