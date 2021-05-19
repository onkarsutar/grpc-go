package main

import (
	"context"
	"log"
	"net"

	"github.com/onkarsutar/grpc-go/user/user_server/helper"
	"github.com/onkarsutar/grpc-go/user/userpb"
	"google.golang.org/grpc"
)

type server struct {
	userpb.UserServiceServer
}

func (*server) GetUserByID(c context.Context, req *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	log.Printf("Invoked GetUserByID with %v\n", req)
	id := req.GetID()

	user, err := helper.FindUser(id)
	if err != nil {
		log.Fatalf("Failed to get user %v", err)
	}

	res := userpb.GetUserByIDResponse{
		User: &user,
	}

	return &res, nil
}

func main() {
	log.Println("Hello from Server...")
	lis, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err)
	}

	s := grpc.NewServer()

	userpb.RegisterUserServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
