package main

import (
	"context"
	"log"
	"net"
	"time"

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

func (*server) GetAllUsers(req *userpb.GetAllUsersRequest, stream userpb.UserService_GetAllUsersServer) error {
	log.Printf("Invoked GetAllUsers with %v\n", req)
	users, err := helper.GetAllUsers()
	if err != nil {
		log.Fatalf("Failed to get user %v", err)
	}

	for i := 0; i < len(users); i++ {
		res := &userpb.GetAllUsersResponse{
			User: &users[i],
		}
		log.Printf("Sending %s", res.User.GetFName())
		stream.Send(res)
		time.Sleep(time.Second)
	}

	return nil

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
