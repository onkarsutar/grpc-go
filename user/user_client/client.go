package main

import (
	"context"
	"io"
	"log"

	"github.com/onkarsutar/grpc-go/user/userpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello from Server...")

	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}
	defer conn.Close()

	c := userpb.NewUserServiceClient(conn)
	// getUser(c)
	getAllUsers(c)
}

func getUser(c userpb.UserServiceClient) {
	req := userpb.GetUserByIDRequest{
		ID: 10,
	}

	res, err := c.GetUserByID(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed ti call GetUserByID %v\n", err)
	}

	log.Println("Response : ", res)
}

func getAllUsers(c userpb.UserServiceClient) {

	req := userpb.GetAllUsersRequest{}

	res, err := c.GetAllUsers(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed ti call GetUserByID %v\n", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read stream %v\n", err)
		}

		log.Println(msg.GetUser())
	}
}
