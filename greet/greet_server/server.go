package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/onkarsutar/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.GreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName

	res := greetpb.GreetResponse{
		Result: result,
	}

	return &res, nil
}

func (*server) GreetMany(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyServer) error {
	log.Printf("GreetMany invoked with %v", req)

	firstName := req.Greeting.GetFirstName()

	for i := 0; i < 10; i++ {
		greet := fmt.Sprintf("Hello %s : %d", firstName, i)
		res := &greetpb.GreetManyTimesResponse{
			Result: greet,
		}
		log.Printf("Sending %s", greet)
		stream.Send(res)
		time.Sleep(time.Second)
	}

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet invoked with")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("Client stream Ends")
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		firstName := req.Greeting.GetFirstName()
		result += "Hello " + firstName + "\n"
	}

}

func main() {
	log.Println("Hello From Server.")

	lis, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
