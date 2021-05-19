package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/onkarsutar/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.CalculatorServiceServer
}

func fibGen() func() int64 {
	f1 := int64(0)
	f2 := int64(1)
	return func() int64 {
		f2, f1 = (f1 + f2), f2
		return f1
	}

}

func (*server) FibSeries(req *calculatorpb.FibRequest, stream calculatorpb.CalculatorService_FibSeriesServer) error {
	log.Printf("Received %v", req)
	iFunc := fibGen()

	number := req.GetNumber()
	for i := 0; i < int(number); i++ {
		fib := iFunc()
		res := &calculatorpb.FibResponse{
			Fib: fib,
		}
		log.Printf("Sending %s", res)
		stream.Send(res)
		time.Sleep(time.Second)
	}

	return nil
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Received %v", req)
	num1 := req.GetNum1()
	num2 := req.GetNum2()

	sum := num1 + num2

	res := &calculatorpb.SumResponse{
		Sum: sum,
	}

	return res, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("Unable to serve: %v", err)
	}
}
